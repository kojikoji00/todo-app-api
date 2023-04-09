package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

func (t *Todo) IsNode() {}

func (t *Todo) GetID() string {
	return t.ID
}