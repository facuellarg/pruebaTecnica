package api

import (
	"encoding/json"
	"fmt"
	"ipcom1/parte1/model"

	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

//PurchasesByDay comment
func PurchasesByDay(c echo.Context) error {
	var (
		days          = 1
		date          = c.Param("date")
		daysParameter = c.QueryParam("dias")
		err           error
		response      = model.NewSummaryPurchase()
	)
	if daysParameter != "" {
		days, err = strconv.Atoi(daysParameter)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "dias query parameter should be a number")
		}
	}

	clients := make([]model.Client, 0)
	validDate := regexp.MustCompile(DATE_PATTERN)
	if !validDate.MatchString(date) {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect date format YYYY-MM-DD")
	}
	currentDate, err := DateStringToDate(date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for i := 0; i < days; i++ {
		dateToFetch := TimeToDate(currentDate.AddDate(0, 0, i))
		resp, err := http.Get(fmt.Sprintf(BASE_GET_URL, dateToFetch))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = json.NewDecoder(resp.Body).Decode(&clients); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for _, client := range clients {
			if !client.Compro {
				response.NoCompraron++
				continue
			}
			if client.Monto > response.CompraMasAlta {
				response.CompraMasAlta = client.Monto
			}
			response.ComprasPorTDC[client.Tdc] += client.Monto
			response.Total += client.Monto

		}
	}

	return c.JSON(http.StatusOK, response)
}
