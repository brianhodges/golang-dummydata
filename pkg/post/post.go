package post

type Post struct {
	ID              int    `json:"id"`
	Title			string `json:"title"`
	Text            string `json:"text"`
	Author          string `json:"username"`
	DateTimeCreated string `json:"datetimecreated"`
}
