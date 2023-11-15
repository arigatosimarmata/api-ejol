package handlers

import (
	"net/http"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type EjolHandler struct {
	ejolUsecase *usecases.EjolUseCase
}

type EjolHandlerInterface interface {
	GetEjlogNFS(c *fiber.Ctx) error
}

func NewEjolHandler(eu *usecases.EjolUseCase) *EjolHandler {
	return &EjolHandler{eu}
}

func (eh *EjolHandler) GetEjlogNFS(c *fiber.Ctx) error {
	fn := new(dto.EjolRequest)
	var out interface{}

	if err := c.BodyParser(fn); err != nil {
		return c.JSON(domain.WebResponse{
			Message: "",
			Code:    http.StatusOK,
			Data:    out,
		})
	}

	out = eh.ejolUsecase.GetEjlogNFS(fn.DateInsert.Format("20060102"))
	return c.JSON(domain.WebResponse{
		Message: "",
		Code:    http.StatusOK,
		Data:    out,
	})
}

func (eh *EjolHandler) GetEjlogDB(c *fiber.Ctx) error {

	response := eh.ejolUsecase.GetEjlogNFS("")
	return c.JSON(domain.WebResponse{
		Message: "",
		Code:    http.StatusOK,
		Data:    response,
	})
}
