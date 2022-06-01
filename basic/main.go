package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "os"

    "github.com/tidwall/gjson"
)

const city string = "London"

const locationLat string = "51.5002"
const locationLon string = "-0.1262"

var codes = map[string]string{
    "0": "Clear sky",
    "1": "Mainly clear",
    "2": "Partly cloudy",
    "3": "Overcast",
    "45": "Fog and depositing rime fog",
    "48": "Depositing rime fog",
    "51": "Drizzle: Light intensity",
    "53": "Drizzle: Moderate intensity",
    "55": "Drizzle: Dense intensity",
    "56": "Freezing Drizzle: Light intensity",
    "57": "Freezing Drizzle: Dense intensity",
    "61": "Rain: Slight intensity",
    "63": "Rain: Moderate intensity",
    "65": "Rain: Heavy intensity",
    "66": "Freezing Rain: Light intensity",
    "67": "Freezing Rain: Heavy intensity",
    "71": "Snow fall: Slight intensity",
    "73": "Snow fall: Moderate intensity",
    "75": "Snow fall: Heavy intensity",
    "77": "Snow grains",
    "80": "Rain showers: Slight",
    "81": "Rain showers: Moderate",
    "82": "Rain showers: Violent",
    "85": "Snow showers: Slight",
    "86": "Snow showers: Heavy",
    "95": "Thunderstorm: Slight or moderate",
    "96": "Thunderstorm with slight hail",
    "99": "Thunderstorm with heavy hail",
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func weatherOrNot(city string, latitude string, longitude string) {
    request := []string{
                "https://api.open-meteo.com/v1/forecast?",
                "latitude=", latitude, 
                "&longitude=", longitude, 
                "&daily=weathercode&timezone=Europe%2FLondon"}

    resp, err := http.Get(strings.Join(request[:], ""))
    check(err)
    json, err := ioutil.ReadAll(resp.Body)
    check(err)
    
    currentWeatherCode := gjson.Get(string(json), "daily.weathercode.0")
    value := codes[currentWeatherCode.String()]
    fmt.Printf("The current weather condition in %s is %s.\n", city, value)
}

func main() {
    weatherOrNot(city, locationLat, locationLon)
}
