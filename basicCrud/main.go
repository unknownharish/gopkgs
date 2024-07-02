package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

type movie struct {
	Name   string `json:"name"`
	Mtype  string `json:"movieType"`
	Direct director
}

type director struct {
	Name string `json:"directorName"`
	Age  int64  `json:"directorAge"`
}

var MovieList = []movie{
	{Name: "kabil", Mtype: "thrillor", Direct: director{Name: "harish chand", Age: 30}},
	{Name: "heropanti", Mtype: "romance", Direct: director{Name: "roman ", Age: 40}},
	{Name: "raone", Mtype: "good", Direct: director{Name: "rahul", Age: 31}},
	{Name: "adipurush", Mtype: "thrillor", Direct: director{Name: "aakash", Age: 35}},
}

func main() {

	// MovieList = append(MovieList, movie{Name: "kabil", Mtype: "thrillor", Direct: director{Name: "harish chand", Age: 30}})
	// MovieList = append(MovieList, movie{Name: "heropanti", Mtype: "romance", Direct: director{Name: "roman ", Age: 40}})
	// MovieList = append(MovieList, movie{Name: "raone", Mtype: "good", Direct: director{Name: "rahul", Age: 31}})
	// MovieList = append(MovieList, movie{Name: "adipurush", Mtype: "thrillor", Direct: director{Name: "aakash", Age: 35}})

	router := mux.NewRouter()
	router.HandleFunc("/", getAllHandler).Methods("GET")
	router.HandleFunc("/{id}", getAllHandlerGet).Methods("GET")
	router.HandleFunc("/set", addMovieHandler).Methods("POST")

	fmt.Println("server started at port 4000")
	no, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(no)
	(http.ListenAndServe(":4000", router))

}

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	query.Get("Name")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(query.Get("Name"))

}
func getAllHandlerGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(params["id"])

}

func addMovieHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var mve movie
	err := json.NewDecoder(r.Body).Decode(&mve)
	if err != nil {
		fmt.Println(err)
	}

	MovieList = append(MovieList, mve)

	type respo struct {
		Data   []movie `json:"data"`
		Status string  `json:"status"`
	}
	json.NewEncoder(w).Encode(respo{Status: "inserted", Data: MovieList})

}
