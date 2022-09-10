package models

type ReportMerchantsRequest struct {
	FromDate string `json:"fromDate" form:"fromDate" query:"fromDate" validate:"required"`
	ToDate   string `json:"toDate" form:"toDate" query:"toDate" validate:"required"`
	Skip     int    `json:"skip" form:"skip" query:"skip"`
	Take     int    `json:"take" form:"take" query:"take"`
}

type ReportMerchantsResponse struct {
	Data []ReportMerchantsDetail `json:"data" form:"data" query:"data"`
	Skip int                     `json:"skip" form:"skip" query:"skip"`
	Take int                     `json:"take" form:"take" query:"take"`
}

type ReportMerchantsDetail struct {
	MerchantName string `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	Omzet        int    `json:"omzet" form:"omzet" query:"omzet"`
	Tanggal      string `json:"tanggal" form:"tanggal" query:"tanggal"`
}

type ReportOutletsRequest struct {
	FromDate string `json:"fromDate" form:"fromDate" query:"fromDate" validate:"required"`
	ToDate   string `json:"toDate" form:"toDate" query:"toDate" validate:"required"`
	Skip     int    `json:"skip" form:"skip" query:"skip"`
	Take     int    `json:"take" form:"take" query:"take"`
}

type ReportOutletsResponse struct {
	Data []ReportOutletsDetail `json:"data" form:"data" query:"data"`
	Skip int                   `json:"skip" form:"skip" query:"skip"`
	Take int                   `json:"take" form:"take" query:"take"`
}

type ReportOutletsDetail struct {
	OutletName   string `json:"outlet_name" form:"outlet_name" query:"outlet_name"`
	MerchantName string `json:"merchant_name" form:"merchant_name" query:"merchant_name"`
	Omzet        int    `json:"omzet" form:"omzet" query:"omzet"`
	Tanggal      string `json:"tanggal" form:"tanggal" query:"tanggal"`
}
