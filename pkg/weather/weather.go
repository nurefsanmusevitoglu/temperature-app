package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

type Coordinate struct {
	Latitude  string
	Longitude string
}

// const apikey = "apikey"
const apikey = "apikey"

func Temperature(city string) (float64, error) {
	apikey := os.Getenv(apikey)
	coordinates, err := coordinates(city)
	if err != nil {
		return math.NaN(), err
	}

	weatherURL := "http://api.openweathermap.org/data/2.5/weather"
	weatherURL = weatherURL + "?lat=" + coordinates.Latitude
	weatherURL = weatherURL + "&lon=" + coordinates.Longitude
	weatherURL = weatherURL + "&appid=" + apikey

	res, err := http.Get(weatherURL)
	if err != nil {
		log.Println(err)
		return math.NaN(), err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return math.NaN(), err
	}
	log.Println("weather body: ", string(body))

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return math.NaN(), err
	}

	main := result["main"].(map[string]interface{})
	temp := main["temp"].(float64)

	return kelvin2Celsius(temp), nil
}

func coordinates(city string) (Coordinate, error) {
	var coordinate Coordinate

	apikey := os.Getenv(apikey)
	limit := 1
	geocodingURL := "http://api.openweathermap.org/geo/1.0/direct"
	geocodingURL = geocodingURL + "?q=" + city
	geocodingURL = geocodingURL + "&limit=" + strconv.Itoa(limit)
	geocodingURL = geocodingURL + "&appid=" + apikey

	res, err := http.Get(geocodingURL)
	if err != nil {
		log.Println(err)
		return coordinate, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return coordinate, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return coordinate, err
	}

	data := result[0]
	coordinate = Coordinate{
		Latitude:  fmt.Sprintf("%v", data["lat"]),
		Longitude: fmt.Sprintf("%v", data["lon"]),
	}

	return coordinate, nil
}

func kelvin2Celsius(k float64) float64 {
	return (k - 273.15)
}
