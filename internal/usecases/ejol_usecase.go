package usecases

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/config"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/dto"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/repository"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/validation"
)

type EjolUseCase struct {
	repository *repository.EjolRepository
}

func NewEjolUseCase(er *repository.EjolRepository) *EjolUseCase {
	return &EjolUseCase{er}
}

func (eu *EjolUseCase) GetEjlogNFS(request dto.EjolNFSRequest) (domain.Ejol, error) {
	var out domain.Ejol
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

	req := os.Getenv("EJLOG_DIRECTORY") + strings.ReplaceAll(request.DateRequest, "-", "") + "/ej_" + strings.ReplaceAll(get_mesin.IpAddress, ".", "_")
	content, err := os.ReadFile(req)
	if err != nil {
		resp := searchEjolNFS(request)
		out = resp
		return out, nil
	}

	out = domain.Ejol{
		Ejlog: string(content),
	}
	return out, nil
}

func searchEjolNFS(request dto.EjolNFSRequest) domain.Ejol {
	var out domain.Ejol
	date_nfs, err := strconv.Atoi(config.NFS_DAYS)
	if err != nil {
		return out
	}

	for i := 0; i < date_nfs; i++ {
		date_now := time.Now().Local().AddDate(0, 0, -i)
		req := os.Getenv("EJLOG_DIRECTORY") + date_now.Format(config.YYYYMMDD) + "/ej_" + strings.ReplaceAll(request.IpAddress, ".", "_")
		_, err := os.Stat(req)
		if err != nil {
			continue
		}

		out.LastFileEjlog = date_now.Format(config.YYYYMMDD_dash)
		break
	}

	return out
}

func (eu *EjolUseCase) GetEjlogDB(request dto.EjolDBRequest) (domain.EjolDBResponse, error) {
	var out domain.EjolDBResponse
	var get_mesin domain.Mesin
	var start_date, end_date time.Time
	var errors error
	if request.Tid != "" {
		get_mesin, errors = eu.repository.GetMesinByTID(request.Tid)
		errors = fmt.Errorf("%s", validation.MsgForTag(errors.Error(), request.Tid))
	}

	if request.IpAddress != "" {
		get_mesin, errors = eu.repository.GetMesinByIp(request.IpAddress)
		errors = fmt.Errorf("%s", validation.MsgForTag(errors.Error(), request.IpAddress))
	}

	if errors != nil {
		return out, fmt.Errorf("%v", errors)
	}

	if request.EndDate != "" {
		end_date, errors = time.Parse(config.YYYYMMDD_HHMMSS_dash, request.EndDate)
		if errors != nil {
			return out, fmt.Errorf("error parse end date %v", errors)
		}
	}

	if request.StartDate != "" {
		start_date, errors = time.Parse(config.YYYYMMDD_HHMMSS_dash, request.StartDate)
		if errors != nil {
			return out, fmt.Errorf("error parse start date %v", errors)
		}
	}

	if request.Page == 0 {
		request.Page = 1
	}

	repo_req := domain.EjolDBRequest{
		IdxEjol:   request.IdxEjlog,
		Tid:       get_mesin.Tid,
		IpAddress: get_mesin.IpAddress,
		StartDate: start_date,
		EndDate:   end_date,
		Limit:     request.Limit,
		Kanwil:    get_mesin.Kanwil,
		Page:      request.Page,
		DbName:    fmt.Sprintf("ejlog_%s_%s", get_mesin.Kanwil, start_date.Format(config.YYYYMMDD)),
		TableName: fmt.Sprintf("ej_%s", strings.ReplaceAll(request.IpAddress, ".", "_")),
	}

	resp, err := eu.repository.GetEjlogDB(repo_req)
	if err != nil {
		return out, nil
	}

	out = resp
	return out, nil
}
