package service

import (
	"apod/models"
	"apod/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	apiKey, _ = os.LookupEnv("NASA_API_KEY")
	apiUrl, _ = os.LookupEnv("NASA_APOD_LINK")
)

type ApodService struct {
	apodRepo repository.ApodRepo
}

func (as *ApodService) SaveApod() {

	date := time.Now().UTC()

	params := url.Values{}
	params.Set("api_key", apiKey)
	params.Set("date", date.Format("2006-01-02"))

	imageUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())

	resp, err := http.Get(imageUrl)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	var apod models.ApodModel
	err = json.NewDecoder(resp.Body).Decode(&apod)
	if err != nil {
		log.Println(err)
	}

	imageResponse, err := http.Get(apod.HdUrl)
	if err != nil {
		log.Println(err)
	}
	defer imageResponse.Body.Close()

	apod.Content, err = io.ReadAll(imageResponse.Body)
	as.apodRepo.SaveApod(apod)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("APOD for " + date.String() + " has been saved")
}

type GetApodRequest struct {
	date time.Time
}

func (as *ApodService) GetApodByDate(w http.ResponseWriter, r *http.Request) {

	apodRequest := GetApodRequest{}

	err := json.NewDecoder(r.Body).Decode(&apodRequest)
	if err != nil {
		log.Println(err)
	}

	date := apodRequest.date

	apod := as.apodRepo.GetApodByDate(date)

	err = json.NewEncoder(w).Encode(&apod)
	if err != nil {
		log.Println(err)
	}

}

func (as *ApodService) GetAllApods(w http.ResponseWriter, r *http.Request) {
	apods := as.apodRepo.GetAllPics()

	err := json.NewEncoder(w).Encode(&apods)
	if err != nil {
		log.Println(err)
	}
}
