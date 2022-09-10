package repository

import (
	"test-case-backend-majoo/database"
	"test-case-backend-majoo/models"
)

func GetMerchantByUserId(userId int) (models.Merchants, error) {
	db, _ := database.Connect()
	query := GetMerchantByUserIdQuery
	args := []interface{}{userId}

	rows, err := db.Query(query, args...)
	if err != nil {
		return models.Merchants{}, err
	}

	defer rows.Close()
	result := models.Merchants{}
	for rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.UserId,
			&result.MerchantName,
			&result.CreatedAt,
			&result.CreatedBy,
			&result.UpdatedAt,
			&result.UpdatedBy,
		)
	}

	return result, nil
}
