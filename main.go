package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Fooddetail struct {
	ID    string `json:"id"`
	Image struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"image"`
	Name      string `json:"name"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
	Type string `json:"type"`
}

type ColorThumbnail struct {
	ColorID string `json:"colorID"`
	Colors  struct {
		Category string `json:"category"`
		Code     struct {
			Hex  string `json:"hex"`
			Rgba string `json:"rgba"`
		} `json:"code"`
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"colors"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}

func getFood(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Fooddetail

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	FoodID := request.ID
	ImageHeight := request.Image.Height
	ImageURL := request.Image.URL
	ImageWidth := request.Image.Width
	Name := request.Name
	ThumbnailHeight := request.Thumbnail.Height
	ThumbnailURL := request.Thumbnail.URL
	ThumbnailWidth := request.Thumbnail.Width
	Type := request.Type

	stmt, err := db.Prepare("INSERT INTO fooddetail (FoodID,ImageHeight,ImageURL,ImageWidth,Name,ThumbnailHeight,ThumbnailURL,ThumbnailWidth,Type) VALUES(?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(FoodID, ImageHeight, ImageURL, ImageWidth, Name, ThumbnailHeight, ThumbnailURL, ThumbnailWidth, Type)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func getColorThumbnail(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request ColorThumbnail

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}
	ColorID := request.ColorID
	Color := request.Colors.Color
	Category := request.Colors.Category
	Type := request.Colors.Type
	Rgba := request.Colors.Code.Rgba
	Hex := request.Colors.Code.Hex
	ThumbnailURL := request.Thumbnail.URL
	ThumbnailHeight := request.Thumbnail.Height
	ThumbnailWidth := request.Thumbnail.Width

	stmt, err := db.Prepare("INSERT INTO colorthumbnail (ColorID,Color,Category,Type,Rgba,Hex,ThumbnailURL,ThumbnailHeight,ThumbnailWidth) VALUES(?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(ColorID, Color, Category, Type, Rgba, Hex, ThumbnailURL, ThumbnailHeight, ThumbnailWidth)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func searchByColorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var colorthumbnail []ColorThumbnail
	params := mux.Vars(r)

	sql := `SELECT
				ColorID,
				IFNULL(Category,'') Category
			FROM colorthumbnail WHERE ColorID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var colorthumb ColorThumbnail
	for result.Next() {

		err := result.Scan(&colorthumb.ColorID, &colorthumb.Colors.Category)

		if err != nil {
			panic(err.Error())
		}

		colorthumbnail = append(colorthumbnail, colorthumb)
	}

	json.NewEncoder(w).Encode(colorthumbnail)
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/task4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8080")

	// Route handles & endpoints
	// pilihan json 1
	r.HandleFunc("/getFood", getFood).Methods("POST")

	//pilihan json 2
	r.HandleFunc("/getColorThumbnail", getColorThumbnail).Methods("POST")
	r.HandleFunc("/searchByColorID/{id}", searchByColorID).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
