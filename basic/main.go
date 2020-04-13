package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"strings"
	"os"

	"github.com/tidwall/gjson"
)

const london string = "44418"

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func weatherOrNot(location string) {
    resp, err := http.Get("https://www.metaweather.com/api/location/" + location)
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
	check(err)
	
	title := gjson.Get(string(body), "title")
	value := gjson.Get(string(body), "consolidated_weather.0.weather_state_name")
	fmt.Printf("The current weather condition in %s is %s.\n", title, strings.ToLower(value.String()))
}

func main() {
	weatherOrNot(london)
}
