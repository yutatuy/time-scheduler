package entity

type Email struct {
	From      string
	Subject   string
	Body      string
	Receivers []string
}
