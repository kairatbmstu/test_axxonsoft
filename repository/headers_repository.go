package repository

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/domain"
	"github.com/gofrs/uuid"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

type HeaderRepository struct {
}

func (h HeaderRepository) Create(tx *sql.Tx, header *domain.Header) error {
	uid, err := uuid.NewV7()
	if err != nil {
		log.Println("an error occurred while generating uuid : ", err.Error())
		return err
	}
	header.Id = uid

	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()
	sb.InsertInto("headers").Cols("id", "request_headers_task_id", "header_name", "header_value").
		Values(header.Id, *header.RequestTaskId, header.Name, header.Value)

	query, args := sb.Build()
	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return err
	}
	return nil
}

func (h HeaderRepository) Update(tx *sql.Tx, header *domain.Header) error {

	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update("headers")
	sb.Set(sb.Equal("request_headers_task_id", header.RequestTaskId))
	sb.Set(sb.Equal("response_headers_task_id", header.ResponseTaskId))
	sb.Set(sb.Equal("header_name", header.Name))
	sb.Set(sb.Equal("header_value", header.Value))
	sb.Where(sb.Equal("id", header.Id))
	query, args := sb.Build()

	_, err := tx.Exec(query, args...)
	if err != nil {
		log.Println("an error occurred while executing Update statement : ", err.Error())
		return err
	}
	return nil
}

func (h HeaderRepository) GetRequestHeaders(tx *sql.Tx, taskId uuid.UUID) (*[]domain.Header, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "request_headers_task_id", "response_headers_task_id", "header_name", "header_value").
		From("headers").Where(sb.Equal("request_headers_task_id", taskId))
	query, args := sb.Build()
	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Println("an error occurred while executing select statement : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	var result = make([]domain.Header, 0)

	for rows.Next() {
		header := new(domain.Header)
		err = rows.Scan(&header.Id, &header.RequestTaskId, &header.ResponseTaskId,
			&header.Name, &header.Value)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		result = append(result, *header)
	}

	return &result, nil
}

func (h HeaderRepository) GetResponseHeaders(tx *sql.Tx, taskId uuid.UUID) (*[]domain.Header, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "request_headers_task_id", "response_headers_task_id", "header_name", "header_value").
		From("headers").Where(sb.Equal("response_headers_task_id", taskId))
	query, args := sb.Build()
	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Println("an error occurred while executing insert statement : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	var result = make([]domain.Header, 0)

	for rows.Next() {
		header := new(domain.Header)
		err = rows.Scan(&header.Id, &header.RequestTaskId, &header.ResponseTaskId,
			&header.Name, &header.Value)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		result = append(result, *header)
	}

	return &result, nil
}
