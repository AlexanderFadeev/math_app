package service

import (
	"github.com/stretchr/testify/suite"
	"math_app/app/domain/model"
	"testing"
)

type MathServiceTestSuite struct {
	suite.Suite

	math Math
}

func TestMathServiceTestSuite(t *testing.T) {
	suite.Run(t, &MathServiceTestSuite{})
}

func (s *MathServiceTestSuite) SetupTest() {
	s.math = NewMath()
}

func (s *MathServiceTestSuite) TestAdd() {
	var task = model.Task{
		Operator: model.Add,
		OperandA: 3.5,
		OperandB: 15.25,
	}

	var result, err = s.math.Calculate(task)

	s.Assert().Nil(err)
	s.Assert().Equal(18.75, result)
}

func (s *MathServiceTestSuite) TestSub() {
	var task = model.Task{
		Operator: model.Sub,
		OperandA: 18.75,
		OperandB: 15.25,
	}

	var result, err = s.math.Calculate(task)

	s.Assert().Nil(err)
	s.Assert().Equal(3.5, result)
}

func (s *MathServiceTestSuite) TestMul() {
	var task = model.Task{
		Operator: model.Mul,
		OperandA: 2.5,
		OperandB: 1.1,
	}

	var result, err = s.math.Calculate(task)

	s.Assert().Nil(err)
	s.Assert().Equal(2.75, result)
}

func (s *MathServiceTestSuite) TestDiv() {
	var task = model.Task{
		Operator: model.Div,
		OperandA: 2.75,
		OperandB: 1.1,
	}

	var result, err = s.math.Calculate(task)

	s.Assert().Nil(err)
	s.Assert().Equal(2.5, result)
}

func (s *MathServiceTestSuite) TestDivZero() {
	var task = model.Task{
		Operator: model.Div,
		OperandA: 2.75,
		OperandB: 0,
	}

	var _, err = s.math.Calculate(task)

	s.Assert().Equal(err, ZeroDivisionError)
}

func (s *MathServiceTestSuite) TestInvalidOperand() {
	var task = model.Task{
		Operator: 42,
	}

	s.Assert().Panics(func() {
		_, _ = s.math.Calculate(task)
	})
}
