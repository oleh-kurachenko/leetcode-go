package main

type FooBar struct {
	n       int
	fooChan chan bool
	barChan chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n, fooChan: make(chan bool), barChan: make(chan bool)}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		printFoo()
		fb.fooChan <- true
		<-fb.barChan
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChan
		printBar()
		fb.barChan <- true
	}
}
