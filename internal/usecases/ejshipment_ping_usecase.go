package usecases

import (
	"encoding/json"
	"fmt"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (eu *EjolUseCase) EjshipmentPing(request dto.EjshipmentPingRequest) (domain.PingHostResponse, error) {
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

	req := dto.PingHostRequest{
		Url: fmt.Sprintf("http://%s/api/ping", get_mesin.IpAddress),
	}

	resp := hitPingService(req)
	if resp.Error != nil {
		return out, resp.Error
	}

	out.Code = resp.Code
	out.Message = resp.Message
	return out, nil
}

func hitPingService(req dto.PingHostRequest) dto.PingHostResponse {
	out := dto.PingHostResponse{}
	agent := fiber.Get(req.Url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		fmt.Println("error Get", statusCode, body, errs)
		out.Code = statusCode
		out.Error = fmt.Errorf("%v", errs)
		return out
	}

	var something fiber.Map
	err := json.Unmarshal(body, &something)
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
