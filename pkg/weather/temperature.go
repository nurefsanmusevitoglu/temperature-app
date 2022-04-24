package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	openweathermapApikey = "apikey"
	openweathermapURL    = "http://api.openweathermap.org"
)

type Temperature struct {
	Temperature string `json:"temperature"`
}

func (t *Temperature) GetTemperature(city string) {
	var coordinate Coordinate

	coordinate.GetCoordinate(city)
	log.Println("COORDINATE: " + coordinate.ToString())
	if coordinate.IsEmpty() {
		return
	}

	weatherURL := constructWeatherURL(coordinate)

	res, err := http.Get(weatherURL)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println("weather body: ", string(body))

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return
	}

	main := result["main"].(map[string]interface{})
	temp := main["temp"].(float64)
	tempCelcius := kelvin2Celsius(temp)

	t.Temperature = fmt.Sprintf("%.2f", tempCelcius)
}

func (t *Temperature) IsEmpty() bool {
	return *t == Temperature{}
}

func kelvin2Celsius(k float64) float64 {
	return (k - 273.15)
}

func constructWeatherURL(coordinate Coordinate) string {
	weatherURL := openweathermapURL

	weatherURL += "/data/2.5/weather"
	weatherURL += "?lat=" + coordinate.Latitude
	weatherURL += "&lon=" + coordinate.Longitude
	weatherURL += "&appid=" + os.Getenv(openweathermapApikey)

	return weatherURL
}
