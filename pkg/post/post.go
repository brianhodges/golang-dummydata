package post

type Post struct {
	ID              int    `json:"id"`
	Title			string `json:"title"`
	Text            string `json:"text"`
	Source          string `json:"source"`
	Author          string `json:"username"`
	DateTimeCreated string `json:"datetimecreated"`
}
