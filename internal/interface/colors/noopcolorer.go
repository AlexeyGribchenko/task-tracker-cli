package colors

type NoopColorer struct{}

func (n *NoopColorer) Bold(msg string) string {
	return msg
}

func (n *NoopColorer) Red(msg string) string {
	return msg
}

func (n *NoopColorer) Green(msg string) string {
	return msg
}

func (n *NoopColorer) Yellow(msg string) string {
	return msg
}

func (n *NoopColorer) Blue(msg string) string {
	return msg
}
