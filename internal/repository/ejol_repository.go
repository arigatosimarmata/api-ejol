package repository

import (
	"database/sql"
	"time"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
)

type EjolRepository struct {
	db *sql.DB
}

func NewEjolRepository(db *sql.DB) *EjolRepository {
	return &EjolRepository{db: db}
}

func (er *EjolRepository) GetEjlogDB(string) ([]domain.Ejol, error) {
	var out []domain.Ejol
	var kanwil_id string = "20"

	rows, err := er.db.Query("select ejlog from ejlog_" + kanwil_id + "_" + time.Now().Format("20060102") + "ej_10_102_13_103")
	if err != nil {
		return nil, err
	}
	er.db.Close()

	for rows.Next() {
		var data_ej string
		err := rows.Scan(&data_ej)
		if err != nil {
			continue
		}
		out = append(out, domain.Ejol{Ejlog: data_ej})
	}

	return out, nil
}

func (er *EjolRepository) GetKanwilMesin(tid string) (domain.Ejol, error) {

	return domain.Ejol{}, nil
}
