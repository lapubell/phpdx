package main

type mascot struct {
	name string
}

func NewMascot(n string) mascot {
	return mascot{
		name: n,
	}
}

func (m mascot) sayHello() string {
	return m.name + " says 'Hello!'"
}
