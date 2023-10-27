package repository

import (
	"apod/database"
	"apod/models"
	"log"
	"time"
)

type ApodRepo struct {
}

func (ar *ApodRepo) SaveApod(apodModel models.ApodModel) {

	db := database.GetConnection()

	query := "INSERT INTO PICS VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (date) DO NOTHING"
	stmt, err := db.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	res, err := stmt.Exec(query, apodModel.Copyright, apodModel.Date, apodModel.Explanation, apodModel.HdUrl, apodModel.MediaType,
		apodModel.ServiceVersion, apodModel.Title, apodModel.Url, apodModel.Content)

	log.Println(res)

	if err != nil {
		log.Println(err)
	}
}

func (ar *ApodRepo) GetApodByDate(date time.Time) models.ApodModel {

	db := database.GetConnection()

	apodModel := models.ApodModel{}
	selectQuery := "SELECT * FROM PICS where DATE = $1"

	err := db.QueryRow(selectQuery, date).Scan(&apodModel)
	log.Println("Meow")
	if err != nil {
		log.Println(err)
	}

	return apodModel
}

func (ar *ApodRepo) GetAllPics() []models.ApodModel {
	db := database.GetConnection()

	apods := []models.ApodModel{}

	selectQuery := "SELECT * FROM PICS"

	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var apod models.ApodModel
		err := rows.Scan(&apod.Copyright, &apod.Date, &apod.Explanation, &apod.HdUrl, &apod.MediaType,
			&apod.ServiceVersion, &apod.Title, &apod.Url, &apod.Content)
		if err != nil {
			log.Println(err)
		}
		apods = append(apods, apod)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return apods
}
