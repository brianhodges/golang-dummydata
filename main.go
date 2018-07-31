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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		posts := []post.Post{}

		for i := 1; i <= 20; i++ {
			p := post.Post{}
			p.ID = i
			p.Title = lorem.Sentence(1, 10)
			p.Text = lorem.Paragraph(4, 6)
			p.Source = lorem.Host()
			p.Author = randomdata.FirstName(randomdata.RandomGender)
			p.DateTimeCreated = randomdata.FullDate()
			posts = append(posts, p)
		}

		b, err := json.Marshal(posts)
		util.Check(err)
		fmt.Fprintf(w, string(b[:]))
		return
	}
}

//Route for favicon - Google Chrome fix - calls /index route twice
func handlerICon(w http.ResponseWriter, r *http.Request) {}

//Initialize Server with Routes
func main() {
	fmt.Println("Running local server @ http://localhost:" + os.Getenv("PORT"))
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/data.json", index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
