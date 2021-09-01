package structures

type TestMessage struct {
	Message string `json:"message" validate:"lte=100"`
}
