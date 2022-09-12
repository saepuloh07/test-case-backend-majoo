package repository

import (
	"test-case-backend-majoo/database"
	"test-case-backend-majoo/models"
	"test-case-backend-majoo/utils"
)

func GetReportByMerchants(payload models.ReportMerchantsRequest, userId int) ([]models.ReportMerchantsDetail, error) {
	db, _ := database.Connect()

	query, args := BuildQueryReportByMerchants(payload, userId)
	rows, err := db.Query(query, args...)
	if err != nil {
		return []models.ReportMerchantsDetail{}, err
	}

	defer rows.Close()
	result := []models.ReportMerchantsDetail{}
	for rows.Next() {
		data := models.ReportMerchantsDetail{}
		err = rows.Scan(
			&data.MerchantName,
			&data.Omzet,
			&data.Tanggal,
		)
		result = append(result, data)
	}

	if len(result) != payload.Take {
		result = FillDataReportByMerchants(payload, result, userId)
	}
	return result, nil
}

func CheckDataInArray(elems []models.ReportMerchantsDetail, date string) (bool, int) {
	for i, s := range elems {
		if date == s.Tanggal {
			return true, i
		}
	}
	return false, 0
}

func FillDataReportByMerchants(req models.ReportMerchantsRequest, data []models.ReportMerchantsDetail, userId int) []models.ReportMerchantsDetail {
	var result []models.ReportMerchantsDetail

	merchant, _ := GetMerchantByUserId(userId)

	allDate := utils.GetAllDate(req.FromDate, req.ToDate, req.Skip, req.Take)
	for i, x := range allDate {
		result = append(result, models.ReportMerchantsDetail{
			Tanggal:      x,
			MerchantName: merchant.MerchantName,
			Omzet:        0,
		})
		check, num := CheckDataInArray(data, x)
		if check {
			result[i].Omzet = data[num].Omzet
		}
	}

	return result
}

func GetReportByOutlets(payload models.ReportOutletsRequest, userId int) ([]models.ReportOutletsDetail, error) {
	db, _ := database.Connect()

	query, args := BuildQueryReportByOutlets(payload, userId)
	rows, err := db.Query(query, args...)
	if err != nil {
		return []models.ReportOutletsDetail{}, err
	}

	defer rows.Close()
	result := []models.ReportOutletsDetail{}
	for rows.Next() {
		data := models.ReportOutletsDetail{}
		err = rows.Scan(
			&data.OutletName,
			&data.MerchantName,
			&data.Omzet,
			&data.Tanggal,
		)
		result = append(result, data)
	}

	if len(result) != payload.Take {
		result = FillDataReportByOutlets(payload, result, userId)
	}
	return result, nil
}

func CheckDataInArray1(elems []models.ReportOutletsDetail, data models.ReportOutletsDetail) (bool, int) {
	for i, s := range elems {
		if data.Tanggal == s.Tanggal && data.OutletName == s.OutletName && data.MerchantName == s.MerchantName {
			return true, i
		}
	}
	return false, 0
}

func FillDataReportByOutlets(req models.ReportOutletsRequest, data []models.ReportOutletsDetail, userId int) []models.ReportOutletsDetail {
	var result []models.ReportOutletsDetail

	outlets, _ := GetOutletByUserId(userId)

	allDate := utils.GetAllDate(req.FromDate, req.ToDate, req.Skip, req.Take)

	for _, x := range allDate {
		for _, y := range outlets {
			result = append(result, models.ReportOutletsDetail{
				Tanggal:      x,
				OutletName:   y.OutletName,
				MerchantName: y.MerchantName,
				Omzet:        0,
			})
		}
	}

	for i, z := range result {
		check, num := CheckDataInArray1(data, z)
		if check {
			result[i].Omzet = data[num].Omzet
		}
	}

	return result
}
