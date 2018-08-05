package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	lorem "github.com/drhodes/golorem"
	"golang-dummydata/pkg/post"
	"golang-dummydata/pkg/util"
	"log"
	"net/http"
	"os"
)

//GET /index
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		posts := []post.Post{}

		for i := 1; i <= 20; i++ {
			p := post.Post{
				ID:              i,
				Title:           lorem.Sentence(1, 10),
				Text:            lorem.Paragraph(4, 6),
				Author:          randomdata.FirstName(randomdata.RandomGender),
				DateTimeCreated: randomdata.FullDate(),
			}
			posts = append(posts, p)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		b, err := json.Marshal(posts)
		util.Check(err)
		fmt.Fprintf(w, string(b[:]))
		return
	}
}

//Initialize Server with Routes
func main() {
	fmt.Println("Running local server @ http://localhost:" + os.Getenv("PORT"))
	http.HandleFunc("/data.json", index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
