package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Article struct {
	Id      string `json:Id`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// using this article slice to simulate a DB
var Articles []Article

//if we want to send messages back to the client hitting our endpoints we need to make
// use of the http.ResponseWriter

func homePage(w http.ResponseWriter, r *http.Request) {
	println("Endpoint hit! -> Home Page")
	fmt.Fprintf(w, "Welcome to the Home Page!")

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	println("Endpoint hit! -> return all articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Welcome to Article %v", key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func indexPage(w http.ResponseWriter, r *http.Request) {
	println("Endpoint hit -> Index Page")
	fmt.Fprintf(w, "Welcome to the Index page Regulars")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/index", indexPage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	//importing of the ENV variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	WeatherKey := os.Getenv("WEATHER_KEY")
	fmt.Println(WeatherKey)

	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
