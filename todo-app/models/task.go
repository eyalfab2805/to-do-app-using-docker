package models

type Task struct {
    ID          int    `json:"id" db:"id"`
    Title       string `json:"title" db:"title"`
    Description string `json:"description,omitempty" db:"description"`
    Completed   bool   `json:"completed" db:"completed"`
}
