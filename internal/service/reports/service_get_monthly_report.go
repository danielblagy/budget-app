package reports

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) GetMonthlyReport(ctx context.Context, username string, year int, month int) (*model.MonthlyReport, error) {
	report, err := s.reportsQuery.GetMonthlyReport(ctx, username, year, month)
	if err != nil {
		return nil, errors.Wrap(err, "can't get monthly report")
	}

	var expenses []*model.MonthlyReportInputItem
	var incomes []*model.MonthlyReportInputItem

	for i := 0; i < len(report); i++ {
		var inputItem model.MonthlyReportInputItem
		inputItem.CategoryId, inputItem.CategoryName, inputItem.Amount = report[i].CategoryId, report[i].CategoryName, report[i].Amount
		if report[i].Type == model.CategoryTypeExpense {
			expenses = append(expenses, &inputItem)
		} else {
			incomes = append(incomes, &inputItem)
		}
	}

	var fullReport model.MonthlyReport
	fullReport.User, fullReport.Year, fullReport.Month = username, year, month
	fullReport.Expenses, fullReport.Incomes = expenses, incomes
	return &fullReport, nil
}
