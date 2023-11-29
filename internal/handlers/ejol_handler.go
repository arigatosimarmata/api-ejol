package handlers

import (
	"bitbucket.bri.co.id/scm/ejol/api-ejol/helper"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/usecases"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/validation"
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
	myValidator := &validation.XValidator{
		Validator: validation.Validate,
	}
	myValidator.Validator.RegisterValidation("DateOnlyWithDash", validation.DateOnlyWithDash)
	myValidator.Validator.RegisterValidation("DateLessThan", validation.DateLessThan)

	var request dto.EjolNFSRequest
	var out interface{}
	response_error := domain.ResponseError{
		Code:    fiber.StatusInternalServerError,
		Message: helper.Msg_StatusInternalServerError,
		Error:   out,
	}

	response := domain.WebResponse{
		Message: helper.Msg_StatusOK,
		Code:    fiber.StatusOK,
		Data:    out,
	}
	errMsgs := []domain.Validator{}

	if err := c.BodyParser(&request); err != nil {
		errMsgs = append(errMsgs, domain.Validator{
			Message: err.Error(),
		})

		response_error.Message = helper.Msg_StatusUnprocessableEntity
		response_error.Code = fiber.StatusUnprocessableEntity
		response_error.Error = errMsgs
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response_error)
	}

	errs := myValidator.ValidateStruct(request)
	if len(errs) > 0 && errs[0].Error {
		errMsgs = []domain.Validator{}
		for _, err := range errs {
			errMsgs = append(errMsgs, domain.Validator{
				Field:   err.FailedField,
				Message: err.Tag,
			})
		}

		if len(errMsgs) < 1 {
			return c.Status(fiber.StatusInternalServerError).JSON(&response_error)
		}

		response_error.Code = fiber.StatusBadRequest
		response_error.Message = helper.Msg_BadRequest
		response_error.Error = errMsgs
		return c.Status(fiber.StatusBadRequest).JSON(response_error)
	}

	out_nfs, err := eh.ejolUsecase.GetEjlogNFS(request)
	if err != nil {
		response_error.Code = fiber.StatusNotFound
		response_error.Message = helper.Msg_StatusNotFound
		response_error.Error = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(response_error)
	}

	if len(out_nfs.Ejlog) < 1 {
		response.Code = fiber.StatusNotFound
		response.Message = "failed"
		response.Data = out_nfs
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response.Data = out_nfs
	return c.Status(fiber.StatusOK).JSON(response)
}

func (eh *EjolHandler) GetEjlogDB(c *fiber.Ctx) error {
	myValidator := &validation.XValidator{
		Validator: validation.Validate,
	}
	myValidator.Validator.RegisterValidation("DateTimeWithDash", validation.DateTimeWithDash)

	var request dto.EjolDBRequest
	var out interface{}
	response_error := domain.ResponseError{
		Code:    fiber.StatusInternalServerError,
		Message: helper.Msg_StatusInternalServerError,
		Error:   out,
	}

	response := domain.WebResponse{
		Message: helper.Msg_StatusOK,
		Code:    fiber.StatusOK,
		Data:    out,
	}
	errMsgs := []domain.Validator{}

	if err := c.BodyParser(&request); err != nil {
		errMsgs = append(errMsgs, domain.Validator{
			Message: validation.MsgForTag(err.Error(), nil),
		})

		response_error.Message = helper.Msg_StatusUnprocessableEntity
		response_error.Code = fiber.StatusUnprocessableEntity
		response_error.Error = errMsgs
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response_error)
	}

	errs := myValidator.ValidateStruct(request)
	if len(errs) > 0 && errs[0].Error {
		errMsgs = []domain.Validator{}
		for _, err := range errs {
			errMsgs = append(errMsgs, domain.Validator{
				Field:   err.FailedField,
				Message: err.Tag,
			})
		}

		if len(errMsgs) < 1 {
			return c.Status(fiber.StatusInternalServerError).JSON(&response_error)
		}

		response_error.Code = fiber.StatusBadRequest
		response_error.Message = helper.Msg_BadRequest
		response_error.Error = errMsgs
		return c.Status(fiber.StatusBadRequest).JSON(response_error)
	}

	res, err := eh.ejolUsecase.GetEjlogDB(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response_error)
	}

	if len(res.Data) < 1 {
		response.Code = fiber.StatusNotFound
		response.Message = "failed"
		response.Data = out
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response.Data = res
	return c.Status(fiber.StatusOK).JSON(response)
}
