package handlers

import (
	"net/http"
	"strconv"
	"test-case-backend-majoo/middlewares"
	"test-case-backend-majoo/models"
	"test-case-backend-majoo/repository"

	"github.com/labstack/echo/v4"
)

func GetReportByMerchants(c echo.Context) error {
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	take, _ := strconv.Atoi(c.QueryParam("take"))
	if take == 0 {
		take = 10
	}
	payload := models.ReportMerchantsRequest{
		FromDate: c.QueryParam("fromDate"),
		ToDate:   c.QueryParam("toDate"),
		Skip:     skip,
		Take:     take,
	}
	userFromToken, _ := middlewares.GetDataFromToken(c)

	resData, err := repository.GetReportByMerchants(payload, userFromToken.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	resp := models.ReportMerchantsResponse{
		Data: resData,
		Skip: skip,
		Take: take,
	}

	return c.JSON(http.StatusOK, resp)
}

func GetReportByOutlets(c echo.Context) error {
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	take, _ := strconv.Atoi(c.QueryParam("take"))
	if take == 0 {
		take = 10
	}
	payload := models.ReportOutletsRequest{
		FromDate: c.QueryParam("fromDate"),
		ToDate:   c.QueryParam("toDate"),
		Skip:     skip,
		Take:     take,
	}

	if err := c.Validate(payload); err != nil {
		return err
	}

	userFromToken, _ := middlewares.GetDataFromToken(c)
	resData, err := repository.GetReportByOutlets(payload, userFromToken.ID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	resp := models.ReportOutletsResponse{
		Data: resData,
		Skip: skip,
		Take: take,
	}

	return c.JSON(http.StatusOK, resp)
}
