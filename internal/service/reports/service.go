package reports

import (
	"context"

	"github.com/danielblagy/budget-app/internal/db"
	"github.com/danielblagy/budget-app/internal/model"
)

type Service interface {
	GetMonthlyReport(ctx context.Context, username string, year int, month int) ([]*model.MonthlyReportItem, error)
}

type service struct {
	reportsQuery db.ReportsQuery
}

func NewService(reportsQuery db.ReportsQuery) Service {
	return &service{
		reportsQuery: reportsQuery,
	}
}
