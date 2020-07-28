package http_server

import (
	"bytes"
	"encoding/json"
	"github.com/AlexanderFadeev/sentinel"
	"io"
	"math_app/app/domain/model"
	"math_app/app/domain/service"
	"net/http"
	"net/url"
	"strconv"
)

var InvalidArgsError = sentinel.Error("Invalid args")

type handler interface {
	Add(http.ResponseWriter, *http.Request)
	Sub(http.ResponseWriter, *http.Request)
	Mul(http.ResponseWriter, *http.Request)
	Div(http.ResponseWriter, *http.Request)
}

type response struct {
	Success bool    `json:"Success"`
	ErrCode string  `json:"ErrCode"`
	Value   float64 `json:"Value"`
}

type handlerImpl struct {
	svc          service.Math
	errorHandler ErrorHandler
}

func newHandler(errorHandler ErrorHandler) handler {
	return &handlerImpl{
		svc:          service.NewMath(),
		errorHandler: errorHandler,
	}
}

func (h *handlerImpl) Add(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req, model.Add)
}

func (h *handlerImpl) Sub(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req, model.Sub)
}

func (h *handlerImpl) Mul(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req, model.Mul)
}

func (h *handlerImpl) Div(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req, model.Div)
}

func (h *handlerImpl) handle(w http.ResponseWriter, req *http.Request, operator model.Operator) {
	result, err := h.handleImpl(req, operator)
	var errCode string
	if err != nil {
		errCode = err.Error()
	}

	resp := response{
		Success: err == nil,
		ErrCode: errCode,
		Value:   result,
	}

	var buf = bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encodeErr := encoder.Encode(&resp)
	if encodeErr != nil {
		w.WriteHeader(translateError(encodeErr))
		return
	}

	w.WriteHeader(translateError(err))
	_, err = io.Copy(w, &buf)
	if err != nil {
		h.errorHandler(err)
	}
}

func (h *handlerImpl) handleImpl(req *http.Request, operator model.Operator) (float64, error) {
	opA, opB, err := h.parseOperands(req)
	if err != nil {
		return 0, err
	}

	task := model.Task{
		Operator: operator,
		OperandA: opA,
		OperandB: opB,
	}
	return h.svc.Calculate(task)
}

func (h *handlerImpl) parseOperands(req *http.Request) (float64, float64, error) {
	query := req.URL.Query()
	if len(query) != 2 {
		return 0, 0, InvalidArgsError
	}

	opA, err := h.getFloatValueFromQuery(query, "a")
	if err != nil {
		return 0, 0, err
	}

	opB, err := h.getFloatValueFromQuery(query, "b")
	if err != nil {
		return 0, 0, err
	}

	return opA, opB, nil
}

func (h *handlerImpl) getFloatValueFromQuery(q url.Values, key string) (float64, error) {
	strVal, err := h.getValueFromQuery(q, key)
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return 0, InvalidArgsError
	}

	return val, nil
}

func (h *handlerImpl) getValueFromQuery(q url.Values, key string) (string, error) {
	vals, ok := q[key]
	if !ok {
		return "", InvalidArgsError
	}

	if len(vals) != 1 {
		return "", InvalidArgsError
	}

	return vals[0], nil
}
