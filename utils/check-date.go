package utils

var (
	FormatDate = "2006-01-02"
)

func GetAllDate(startDate string, endDate string) []string {
	var (
		dateArr        []string
		currentDate, _ = ParseStringToTime(startDate, FormatDate, "start date")
		endDateNew, _  = ParseStringToTime(endDate, FormatDate, "end date")
	)

	difference := endDateNew.Sub(currentDate)
	dayTotal := int(difference.Hours() / 24)

	for i := 0; i <= dayTotal; i++ {
		dateArr = append(dateArr, string(currentDate.Format(FormatDate)))

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dateArr
}
