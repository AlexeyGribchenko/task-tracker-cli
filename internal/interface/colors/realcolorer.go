package colors

import "fmt"

type RealColorer struct{}

func (r *RealColorer) Bold(msg string) string {
	return fmt.Sprintf("%s%s%s", bold, msg, reset)
}

func (r *RealColorer) Red(msg string) string {
	return fmt.Sprintf("%s%s%s", red, msg, reset)
}

func (r *RealColorer) Green(msg string) string {
	return fmt.Sprintf("%s%s%s", green, msg, reset)
}

func (r *RealColorer) Yellow(msg string) string {
	return fmt.Sprintf("%s%s%s", yellow, msg, reset)
}

func (r *RealColorer) Blue(msg string) string {
	return fmt.Sprintf("%s%s%s", blue, msg, reset)
}
