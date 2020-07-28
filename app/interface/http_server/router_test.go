package http_server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math_app/app/domain/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	path     string
	status   int
	response response
}

func TestRouter(t *testing.T) {
	testCases := []testCase{
		{
			path:   "/add?a=3.25&b=4.25",
			status: http.StatusOK,
			response: response{
				Success: true,
				ErrCode: "",
				Value:   7.5,
			},
		},
		{
			path:   "/add?a=3.25",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?b=3.25",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?a=3.25&b=4.25&c=5.3",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?a=3.25&b=4.25&a=5.3",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?a=3.25&b=4.25&b=5.3",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?c=3.25&d=4.25",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/add?a=foo&b=bar",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: InvalidArgsError.Error(),
			},
		},
		{
			path:   "/sub?a=3.25&b=4.25",
			status: http.StatusOK,
			response: response{
				Success: true,
				ErrCode: "",
				Value:   -1.0,
			},
		},
		{
			path:   "/mul?a=-3.5&b=1.5",
			status: http.StatusOK,
			response: response{
				Success: true,
				ErrCode: "",
				Value:   -5.25,
			},
		},
		{
			path:   "/div?a=12.5&b=2.5",
			status: http.StatusOK,
			response: response{
				Success: true,
				ErrCode: "",
				Value:   5.0,
			},
		},
		{
			path:   "/div?a=3.41&b=0",
			status: http.StatusBadRequest,
			response: response{
				Success: false,
				ErrCode: service.ZeroDivisionError.Error(),
			},
		},
	}

	router := newRouter(func(err error) {
		t.Error(err.Error())
	})

	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api"+testCase.path, nil)

		router.ServeHTTP(rec, req)

		body, err := ioutil.ReadAll(rec.Body)
		assert.Nil(t, err)

		var resp response
		err = json.Unmarshal(body, &resp)
		assert.Nilf(t, err, "%#v", testCase)

		assert.Equal(t, testCase.status, rec.Code, "%#v", testCase)
		assert.Equal(t, testCase.response, resp, "%#v", testCase)
	}
}
