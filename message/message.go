package message

const (
	ActionStatus string = "STATUS"
	ActionStore         = "STORE"
	ActionGet           = "GET"
	ActionList          = "LIST"
	ActionDelete        = "DELETE"
	ActionSave          = "SAVE"
)

type Message struct {
	Action string
	Args   map[string]any
}

type Response struct {
	Ok    bool
	Value any
}
