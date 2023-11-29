package repository

import (
	"database/sql"
	"fmt"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/config"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
)

type EjolRepository struct {
	DB *sql.DB
}

func NewEjolRepository(db *sql.DB) *EjolRepository {
	return &EjolRepository{DB: db}
}

func (er *EjolRepository) GetEjlogDB(req domain.EjolDBRequest) (domain.EjolDBResponse, error) {
	var out domain.EjolDBResponse

	dbs, err := config.DBConn(req.DbName)
	if err != nil {
		return out, err
	}
	er.DB = dbs
	defer er.DB.Close()

	query_content := `SELECT id, ejlog, created_at from ` + req.TableName

	if req.Page > 1 {
		data_page_request, err := getIdFromPageRequest(req, er.DB)
		if err != nil {
			fmt.Println("err getIdFromPageRequest", err)
			return out, err
		}

		query_content += fmt.Sprintf(" WHERE id>=%d", data_page_request.IdxPage)
	} else {
		query_content += fmt.Sprintf(" WHERE id>=%d", req.IdxEjol)
	}

	data_pagination, err := getPaginationData(req, er.DB)
	if err != nil {
		fmt.Println("err getPaginationData", err)
		return out, err
	}

	if !req.StartDate.IsZero() {
		query_content += fmt.Sprintf(` AND created_at>="%s"`, req.StartDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	if !req.EndDate.IsZero() {
		query_content += fmt.Sprintf(` AND created_at<="%s"`, req.EndDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	if req.Limit > 100000 {
		query_content += ` ORDER BY id LIMIT ` + fmt.Sprintf("%d", 100000)
	} else {
		query_content += ` ORDER BY id LIMIT ` + fmt.Sprintf("%d", req.Limit)
	}

	rows, err := er.DB.Query(query_content)
	if err != nil {
		return out, err
	}

	for rows.Next() {
		var data_ej string
		var idx_ej int
		var created_at string
		err := rows.Scan(&idx_ej, &data_ej, &created_at)
		if err != nil {
			continue
		}
		out.Data = append(out.Data, domain.EjolDB{
			IdxEjlog:  idx_ej,
			Ejlog:     data_ej,
			CreatedAt: created_at,
		})
	}

	out.Total = data_pagination.TotalCount
	out.CurrentPage = req.Page
	out.LastPage = data_pagination.TotalPage
	return out, nil
}

type GetIdFromPageRequest struct {
	TotalCount int
	IdxPage    int
}

func getIdFromPageRequest(req domain.EjolDBRequest, er *sql.DB) (GetIdFromPageRequest, error) {
	var out GetIdFromPageRequest
	target_offset := (req.Page - 1) * req.Limit
	query_total_data := fmt.Sprintf(`SELECT id FROM %s`, req.TableName)
	if !req.StartDate.IsZero() {
		query_total_data += fmt.Sprintf(` WHERE created_at>="%s"`, req.StartDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	if !req.EndDate.IsZero() {
		query_total_data += fmt.Sprintf(` AND created_at<="%s"`, req.StartDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	if req.Page > 1 {
		query_total_data += fmt.Sprintf(` LIMIT %d,1`, target_offset)
	}

	res_query_total_data := er.QueryRow(query_total_data).Scan(&out.IdxPage)
	if res_query_total_data != nil {
		return out, res_query_total_data
	}

	return out, nil
}

type GetPaginationData struct {
	TotalCount int
	TotalPage  int
}

func getPaginationData(req domain.EjolDBRequest, er *sql.DB) (GetPaginationData, error) {
	var out GetPaginationData
	query_total_page := `SELECT CEIL(COUNT(*) / ?) AS last_page, COUNT(*) as total_count FROM ` + req.TableName
	if !req.StartDate.IsZero() {
		query_total_page += fmt.Sprintf(` WHERE created_at>="%s"`, req.StartDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	if !req.EndDate.IsZero() {
		query_total_page += fmt.Sprintf(` AND created_at<="%s"`, req.StartDate.Format(config.YYYYMMDD_HHMMSS_dash))
	}

	res_query_total_page := er.QueryRow(query_total_page, req.Limit).Scan(&out.TotalPage, &out.TotalCount)
	if res_query_total_page != nil {
		return out, res_query_total_page
	}

	return out, nil
}

func (er *EjolRepository) GetMesinByTID(tid string) (domain.Mesin, error) {
	data, err := er.getDataMesin("tid", tid)
	if err != nil {
		return domain.Mesin{}, err
	}
	return data, nil
}

func (er *EjolRepository) GetMesinByIp(ip string) (domain.Mesin, error) {
	data, err := er.getDataMesin("ip_address", ip)
	if err != nil {
		return domain.Mesin{}, err
	}
	return data, nil
}

func (er *EjolRepository) getDataMesin(param, value string) (domain.Mesin, error) {
	var out domain.Mesin
	dbs, err := config.DBConn("")
	if err != nil {
		return out, err
	}
	er.DB = dbs

	query := `select tid,ip_address,kanwil2 from ` + domain.TableAtmMapping() + ` where ` + param + `=?`
	err_query := er.DB.QueryRow(query, value).Scan(&out.Tid, &out.IpAddress, &out.Kanwil)
	if err_query != nil {
		return out, err_query
	}
	defer er.DB.Close()
	return out, nil
}
