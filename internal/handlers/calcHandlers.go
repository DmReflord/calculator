package handlers

import (
	"calculator/internal/service"
	"net/http"

	"github.com/labstack/echo"
)

type CalculationHandler struct {
	service service.CalculationService
}

func NewCalculationHandler(s service.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calulations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch calculations"})
	}
	return c.JSON(http.StatusOK, calulations)
}

func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req service.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")
	var req service.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updCalc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, updCalc)
}

func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteCalculation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}
