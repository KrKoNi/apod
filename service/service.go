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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(imageResponse.Body)

	apod.Content, err = io.ReadAll(imageResponse.Body)
	as.apodRepo.SaveApod(apod)

	if err != nil {
		log.Println(err)
	}
}

func (as *ApodService) GetApodByDate(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	requestDate := params.Get("date")

	date, err := time.Parse("2006-01-02", requestDate)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid date format. Please use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	apod := as.apodRepo.GetApodByDate(date)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get APOD from database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(apod)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode APOD to JSON", http.StatusInternalServerError)
		return
	}
}

func (as *ApodService) GetAllApods(w http.ResponseWriter, r *http.Request) {
	apods := as.apodRepo.GetAllPics()

	err := json.NewEncoder(w).Encode(apods)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode APODs to JSON", http.StatusInternalServerError)
		return
	}
}
