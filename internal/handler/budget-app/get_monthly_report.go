package budget_app

import "github.com/gofiber/fiber/v2"

func (h handler) GetMonthlyReport(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	reportYearRow, err := c.ParamsInt("year", 2023)
	if err != nil {
		return err //return err in case invalid types of data for year
	}
	if reportYearRow < 2000 {
		return c.Status(fiber.StatusBadRequest).SendString("year is not valid")
	}
	reportMonthRow, err := c.ParamsInt("month", 0)
	if err != nil {
		return err //return err in case invalid types of data for month
	}
	if reportMonthRow <= 0 && reportMonthRow > 12 {
		return c.Status(fiber.StatusBadRequest).SendString("month is not valid")
	}

	fullReport, err := h.reportsService.GetMonthlyReport(c.Context(), username, reportYearRow, reportMonthRow)
	if err != nil {
		return err
	}
	return c.JSON(fullReport)
}
