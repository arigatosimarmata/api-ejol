package handlers

import (
	"bitbucket.bri.co.id/scm/ejol/api-ejol/helper"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/validation"
	"github.com/gofiber/fiber/v2"
)

func (eh *EjolHandler) PingEjshipment(c *fiber.Ctx) error {
	myValidator := &validation.XValidator{
		Validator: validation.Validate,
	}

	var request dto.EjshipmentPingRequest
	out := make([]string, 0)
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

	out_ping, err := eh.ejolUsecase.EjshipmentPing(request)
	if err != nil {
		response_error.Code = fiber.StatusNotFound
		response_error.Message = helper.Msg_StatusNotFound
		response_error.Error = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(response_error)
	}

	if out_ping.Code != fiber.StatusOK {
		response.Code = fiber.StatusInternalServerError
		response.Message = "failed"
		response.Data = out_ping.Error
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response.Data = out_ping.Message
	return c.Status(fiber.StatusOK).JSON(response)
}
