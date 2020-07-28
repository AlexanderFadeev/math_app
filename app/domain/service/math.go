package service

import (
	"github.com/AlexanderFadeev/sentinel"
	"math_app/app/domain/model"
)

var ZeroDivisionError = sentinel.Error("Zero division")

type Math interface {
	Calculate(model.Task) (float64, error)
}

type math struct{}

func NewMath() Math {
	return &math{}
}

func (m math) Calculate(task model.Task) (float64, error) {
	switch task.Operator {
	case model.Add:
		return m.add(task)
	case model.Sub:
		return m.sub(task)
	case model.Mul:
		return m.mul(task)
	case model.Div:
		return m.div(task)
	default:
		panic("service.Math: invalid operand")
	}
}

func (m math) add(task model.Task) (float64, error) {
	return task.OperandA + task.OperandB, nil
}

func (m math) sub(task model.Task) (float64, error) {
	return task.OperandA - task.OperandB, nil
}

func (m math) mul(task model.Task) (float64, error) {
	return task.OperandA * task.OperandB, nil
}

func (m math) div(task model.Task) (float64, error) {
	if task.OperandB == 0 {
		return 0, ZeroDivisionError
	}

	return task.OperandA / task.OperandB, nil
}
