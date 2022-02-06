package models

//easyjson:json
type Task struct {
	Id          int64  `json:"id,omitempty"`
	Creator     string `json:"creator,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Executor    string `json:"executor,omitempty"`
	Created     string `json:"created,omitempty"`
	Completed   bool   `json:"completed"`
}
