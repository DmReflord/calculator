package service

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id string, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	err = s.repo.CreateCalculation(calc)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}

func (s *calcService) GetAllCalculation() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}

	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc.Expression = expression
	calc.Result = result

	err = s.repo.UpdateCalculation(calc)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
}
