package utils

var (
	FormatDate = "2006-01-02"
)

func GetAllDate(startDate string, endDate string, start int, limit int) []string {
	var (
		dateArr        []string
		currentDate, _ = ParseStringToTime(startDate, FormatDate, "start date")
		endDateNew, _  = ParseStringToTime(endDate, FormatDate, "end date")
	)

	difference := endDateNew.Sub(currentDate)
	dayTotal := int(difference.Hours() / 24)
	if dayTotal != (limit - 1) {
		dayTotal = limit - 1
	}

	for i := start; i <= dayTotal; i++ {
		dateArr = append(dateArr, string(currentDate.Format(FormatDate)))

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dateArr
}
