package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Coordinate struct {
	Latitude  string
	Longitude string
}

func (c *Coordinate) GetCoordinate(city string) {
	
	geocodingURL := constructCoordinateURL(city)

	res, err := http.Get(geocodingURL)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println("coordinate body: ", string(body))

	var result []map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return
	}

	data := result[0]

	lat := fmt.Sprintf("%v", data["lat"])
	lon := fmt.Sprintf("%v", data["lon"])

	c.Latitude = lat
	c.Longitude = lon
}

func (c *Coordinate) IsEmpty() bool {
	return *c == Coordinate{}
}

func (c *Coordinate) ToString() string {
	return "Coordinate: " + c.Latitude + " ," + c.Longitude
}

func constructCoordinateURL(city string) string {
	limit := 1
	geocodingURL := openweathermapURL

	geocodingURL += "/geo/1.0/direct"
	geocodingURL += "?q=" + city
	geocodingURL += "&limit=" + strconv.Itoa(limit)
	geocodingURL += "&appid=" + os.Getenv(openweathermapApikey)

	return geocodingURL
}