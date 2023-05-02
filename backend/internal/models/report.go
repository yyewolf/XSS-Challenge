package models

import (
	"backend/internal/database"
	"errors"
)

type Report struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Issue string `json:"issue"`
}

func NewReport(name string, email string, issue string) error {
	r := Report{
		Name:  name,
		Email: email,
		Issue: issue,
	}

	return r.Create()
}

func (r *Report) Validate() bool {
	if r.Name == "" || r.Email == "" || r.Issue == "" {
		return false
	}
	return true
}

func (r *Report) Create() error {
	if !r.Validate() {
		return errors.New("invalid report")
	}
	r.ID = database.Add("reports", r)
	return nil
}

func GetAllReports() ([]*Report, error) {
	val, found := database.GetAll("reports")
	if !found {
		return []*Report{}, nil
	}

	reports := make([]*Report, len(val))
	for i, v := range val {
		reports[i] = v.(*Report)
		reports[i].ID = i
	}

	return reports, nil
}

func GetReport(id int) (*Report, error) {
	val, found := database.GetOne("reports", id)
	if !found {
		return nil, errors.New("report not found")
	}

	return val.(*Report), nil
}

func DelReport(id int) error {
	database.Delete("reports", id)
	return nil
}
