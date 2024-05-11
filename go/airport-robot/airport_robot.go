package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type German struct {
}

func (g German) LanguageName() string {
	return "German"
}
func (g German) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Hallo %s!", g.LanguageName(), name)
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", i.LanguageName(), name)
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", p.LanguageName(), name)
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}
