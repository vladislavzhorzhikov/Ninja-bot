package database

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

const NotMatching = "not matching"

type Randoms struct {
	repo  *sql.DB
	table string
}

type Random struct {
	UserName  string `json:"user_name"`
	RandCount int    `json:"rand_count"`
}

func NewRandoms(s *sql.DB, table string) *Randoms {
	return &Randoms{s, table}
}

func (r *Randoms) IsPrevious(userName string) (int, error) {
	row := r.repo.QueryRow(r.getPrevious())
	var random Random
	err := row.Scan(&random.UserName, &random.RandCount)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	if userName != random.UserName {
		return 0, errors.New(NotMatching)
	}
	return random.RandCount, nil
}

func (r *Randoms) getPrevious() string {
	return fmt.Sprintf(`
		SELECT user_name, rand_count
		FROM %s`,
		r.table)
}

func (r *Randoms) AddUserName(userName string) error {
	_, err := r.repo.Query(r.addUserName(), userName)
	if err != nil {
		errors.WithStack(err)
		return err
	}
	return nil
}

func (r *Randoms) addUserName() string {
	return fmt.Sprintf(`
		INSERT INTO %s VALUES ($1, 0)`,
		r.table)
}

func (r *Randoms) UpCount(userName string, count int) error {
	_, err := r.repo.Query(r.upCount(), userName, count)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *Randoms) upCount() string {
	return fmt.Sprintf(`
		UPDATE %s SET user_name = $1, rand_count = $2
		WHERE id = 0`,
		r.table)
}
