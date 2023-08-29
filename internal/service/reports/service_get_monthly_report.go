package reports

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) GetMonthlyReport(ctx context.Context, username string, year int, month int) ([]*model.MonthlyReportItem, error) {
	report, err := s.reportsQuery.GetMonthlyReport(ctx, username, year, month)
	if err != nil {
		return nil, errors.Wrap(err, "can't get monthly report")
	}
	return report, nil
}
