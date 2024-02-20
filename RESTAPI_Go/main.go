package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Genre    string    `json:"genre"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	// Initialize the movies slice
	movies = []Movie{}

	// Define movies
	movie1 := Movie{
		ID:    1,
		Name:  "Dabaang",
		Genre: "Adventure",
		Director: &Director{
			FirstName: "Arbaaz",
			LastName:  "Khan",
		},
	}

	movie2 := Movie{
		ID:    2,
		Name:  "MS Dhoni: The untold story",
		Genre: "Sports",
		Director: &Director{
			FirstName: "Neeraj",
			LastName:  "Pandey",
		},
	}

	movie3 := Movie{
		ID:    3,
		Name:  "Deewar",
		Genre: "Adventure",
		Director: &Director{
			FirstName: "Yash",
			LastName:  "Chopra",
		},
	}

	// Add movies to the slice
	movies = append(movies, movie1, movie2, movie3)

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/movies", ListMovies).Methods("GET")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page"))
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(movies)
	jsonBytes, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("err")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
