package usecases

import (
	"encoding/json"
	"fmt"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (eu *EjolUseCase) EjshipmentSearch(request dto.EjshipmentSearchRequest) (domain.PingHostResponse, error) {
	var out domain.PingHostResponse
	var get_mesin domain.Mesin
	var errors error

	if request.Tid != "" {
		get_mesin, errors = eu.repository.GetMesinByTID(request.Tid)
	}

	if request.IpAddress != "" {
		get_mesin, errors = eu.repository.GetMesinByIp(request.IpAddress)
	}

	if errors != nil {
		return out, fmt.Errorf("err:%v|tid / ip mesin not found", errors)
	}

	req := dto.SearchHostRequest{
		Url:     fmt.Sprintf("http://%s/api/ping", get_mesin.IpAddress),
		Request: request,
	}

	resp := getFileEjlogService(req)
	if resp.Error != nil {
		return out, resp.Error
	}

	out.Code = resp.Code
	out.Message = resp.Message
	return out, nil
}

func getFileEjlogService(req dto.SearchHostRequest) dto.PingHostResponse {
	out := dto.PingHostResponse{}
	agent := fiber.Post(req.Url)
	req_agent, err := json.Marshal(req)
	if err != nil {
		out.Code = fiber.StatusInternalServerError
		out.Error = err
		return out
	}

	agent.Body(req_agent)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		fmt.Println("error Get", statusCode, body, errs)
		out.Code = statusCode
		out.Error = fmt.Errorf("%v", errs)
		return out
	}

	var something fiber.Map
	err = json.Unmarshal(body, &something)
	if err != nil {
		fmt.Println("error unmarshal:", err)
		out.Code = statusCode
		out.Error = err
		return out
	}

	out.Code = statusCode
	out.Message = something
	return out
}
