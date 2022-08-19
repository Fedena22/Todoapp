package data

// Tododata contains the struct for the todo tasks
type Tododata struct {
	ListenGroup string `json:"listengroup"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
