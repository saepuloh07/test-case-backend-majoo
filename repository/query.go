package repository

import (
	"test-case-backend-majoo/models"
)

var (
	// Users
	GetUserAllQuery        = "SELECT * FROM Users"
	GetUserByUsernameQuery = "SELECT * FROM Users WHERE user_name = ? AND password = MD5(?)"

	// Merchants
	GetMerchantByUserIdQuery = "SELECT * FROM Merchants WHERE user_id = ?"

	// Outlets
	GetOutletByUserIdQuery = "SELECT a.outlet_name,b.merchant_name,c.user_name FROM Outlets a INNER JOIN Merchants b ON a.merchant_id=b.id INNER JOIN Users c ON b.user_id=c.id WHERE c.id=?"

	// Transactions
	GetReportByMerchantsQuery = "SELECT b.merchant_name,SUM(a.bill_total) AS omzet,DATE_FORMAT(a.created_at, '%Y-%m-%d') AS tanggal FROM Transactions a INNER JOIN Merchants b ON a.merchant_id=b.id INNER JOIN Users c ON b.user_id=c.id WHERE c.id=?"
	GetReportByOutletsQuery   = "SELECT b.outlet_name,c.merchant_name,SUM(a.bill_total) AS omzet,DATE_FORMAT(a.created_at, '%Y-%m-%d') AS tanggal FROM Transactions a INNER JOIN Outlets b ON a.outlet_id=b.id INNER JOIN Merchants c ON b.merchant_id=c.id INNER JOIN Users d ON c.user_id=d.id WHERE d.id=?"
)

func BuildQueryReportByMerchants(payload models.ReportMerchantsRequest, userId int) (string, []interface{}) {
	baseQuery := GetReportByMerchantsQuery
	args := make([]interface{}, 0)
	args = append(args, userId)

	if payload.FromDate != "" && payload.ToDate != "" {
		baseQuery += " AND a.created_at BETWEEN ? AND ?"

		args = append(args, payload.FromDate+" 00:00:00", payload.ToDate+" 23:59:59")
	}

	baseQuery += " GROUP BY tanggal, b.merchant_name ORDER BY tanggal ASC LIMIT ? OFFSET ?"
	args = append(args, payload.Take, payload.Skip)

	return baseQuery, args
}

func BuildQueryReportByOutlets(payload models.ReportOutletsRequest, userId int) (string, []interface{}) {
	baseQuery := GetReportByOutletsQuery
	args := make([]interface{}, 0)
	args = append(args, userId)

	if payload.FromDate != "" && payload.ToDate != "" {
		baseQuery += " AND a.created_at BETWEEN ? AND ?"

		args = append(args, payload.FromDate+" 00:00:00", payload.ToDate+" 23:59:59")
	}

	baseQuery += " GROUP BY tanggal, b.outlet_name, c.merchant_name ORDER BY tanggal ASC LIMIT ? OFFSET ?"
	args = append(args, payload.Take, payload.Skip)

	return baseQuery, args
}
