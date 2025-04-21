package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type German struct {
}

func (german German) LanguageName() string {
	return "German"
}

func (german German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo %v!", visitorName)
}

type Italian struct{}

func (italian Italian) LanguageName() string {
	return "Italian"
}

func (italian Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %v!", visitorName)
}

type Portuguese struct{}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (portuguese Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %v!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeter.Greet(visitorName))
}
