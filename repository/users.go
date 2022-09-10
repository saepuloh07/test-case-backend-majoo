package repository

import (
	"test-case-backend-majoo/database"
	"test-case-backend-majoo/models"
)

func GetUserAll() ([]models.Users, error) {
	db, _ := database.Connect()
	query := GetUserAllQuery

	rows, err := db.Query(query)
	if err != nil {
		return []models.Users{}, err
	}

	defer rows.Close()
	result := []models.Users{}
	for rows.Next() {
		user := models.Users{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
		)
		result = append(result, user)
	}

	return result, nil
}

func GetUserByUsername(username string, password string) (models.Users, error) {
	db, _ := database.Connect()
	query := GetUserByUsernameQuery
	args := []interface{}{username, password}

	rows, err := db.Query(query, args...)
	if err != nil {
		return models.Users{}, err
	}

	defer rows.Close()
	result := models.Users{}
	for rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Username,
			&result.Password,
			&result.CreatedAt,
			&result.CreatedBy,
			&result.UpdatedAt,
			&result.UpdatedBy,
		)
	}

	return result, nil
}
