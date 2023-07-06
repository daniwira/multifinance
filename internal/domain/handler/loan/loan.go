package handlerloan

import (
	"net/http"
	"strconv"

	"github.com/daniwira/multifinance/internal/domain/loan"
	"github.com/daniwira/multifinance/internal/service"
	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	loanService service.LoanService
}

func NewLoanHandler(loanService service.LoanService) *LoanHandler {
	return &LoanHandler{
		loanService: loanService,
	}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var loan loan.LoanParams
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.loanService.CreateLoan(&loan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (h *LoanHandler) UpdateLoan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	var loan loan.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.loanService.UpdateLoan(uint(id), &loan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (h *LoanHandler) DeleteLoan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	if err := h.loanService.DeleteLoan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}

func (h *LoanHandler) GetLoanByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := h.loanService.GetLoanByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (h *LoanHandler) PaymentInstallment(c *gin.Context) {
	var params loan.PaymentInstallment
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.loanService.PaymentInstallment(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "successfully pay the installment"})
}
