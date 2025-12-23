package entity

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	IsDone      bool   `json:"done"`
}
