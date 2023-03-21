package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TZID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC      float64 `json:"temp_c"`
		FeelsLikeC float64 `json:"feelslike_c"`
		Condition  struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		Humidity   int     `json:"humidity"`
		WindMph    float64 `json:"wind_mph"`
		PressureIn float64 `json:"pressure_in"`
		UV         float64 `json:"uv"`
	} `json:"current"`
}

// if we want to send messages back to the client hitting our endpoints we need to make
// use of the http.ResponseWriter
var client *http.Client

//ENDPOINTS

func weatherPage(w http.ResponseWriter, r *http.Request) {
	println("Port Elizabeth Weather -> End Point")
	fmt.Fprintf(w, "Port Elizabeth Weather -> End Point")
	GetWeatherInfo()

}

func homePage(w http.ResponseWriter, r *http.Request) {
	println("Endpoint hit! -> Home Page")
	fmt.Fprintf(w, "Welcome to the Home Page!")

}

func indexPage(w http.ResponseWriter, r *http.Request) {
	println("Endpoint hit -> Index Page")
	fmt.Fprintf(w, "Welcome to the Index page Regulars")
}

//END OF ENDPOINTS

func GetWeatherInfo() {

	url := "http://api.weatherapi.com/v1/current.json?key=6f811e8ca8a94ee1b6b100326231903&q=London&aqi=no"

	var weatherResponse WeatherResponse

	err := GetJson(url, &weatherResponse)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Current Weather in %v: \n\n %v", weatherResponse.Location.Name, weatherResponse.Current.TempC)
	}

}

// GENERIC METHOD USED FOR PARSING THE JSON DATA WE GET FROM OUR GET REQUESTS
func GetJson(url string, target interface{}) error {

	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// USED FOR THE HANDLING OF REQUESTS
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/index", indexPage)
	myRouter.HandleFunc("/weather", weatherPage)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}
	//importing of the ENV variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	WeatherKey := os.Getenv("WEATHER_KEY")
	fmt.Println(WeatherKey)

	fmt.Println("Rest API v2.0 - Mux Routers")

	handleRequests()
}
