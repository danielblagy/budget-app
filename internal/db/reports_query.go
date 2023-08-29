package db

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type ReportsQuery interface {
	GetMonthlyReport(ctx context.Context, username string, year int, month int) ([]*model.MonthlyReportItem, error)
}

type reportsQuery struct {
	db pgxscan.Querier
}

func newReportsQuery(db pgxscan.Querier) ReportsQuery {
	return &reportsQuery{
		db: db,
	}
}

func (q reportsQuery) GetMonthlyReport(ctx context.Context, username string, year int, month int) ([]*model.MonthlyReportItem, error) {
	var getMonthlyReportTemplate = "select c.id AS category_id,c.name AS category_name, SUM(e.amount) AS amount,e.type FROM entries e INNER JOIN categories c ON e.category_id=c.id WHERE EXTRACT(year from e.date)='%d' and EXTRACT(month from e.date)='%d' and e.user_id = '%s' GROUP BY c.id, e.type ORDER BY e.type"
	var report []*model.MonthlyReportItem

	err := pgxscan.Select(ctx, q.db, &report, fmt.Sprintf(getMonthlyReportTemplate, year, month, username))
	if err != nil {
		return nil, err
	}

	return report, nil
}
