package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

func greet(greeting, name string) string {
    return fmt.Sprintf("%s %s!", greeting, name)
}

type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return greet("Ciao", name)
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
    return greet("Olá", name)
}
