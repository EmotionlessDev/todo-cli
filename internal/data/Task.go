package data

type Task struct {
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
}
