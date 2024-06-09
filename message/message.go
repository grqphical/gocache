package message

const (
	ActionStatus string = "STATUS"
)

type Message interface {
	Action() string
	Values() map[string]any
}

type Response struct {
	Ok    bool
	Value any
}

type StatusMessage struct{}

func (s *StatusMessage) Action() string {
	return ActionStatus
}

func (s *StatusMessage) Values() map[string]any {
	return nil
}
