package main

type ZeroEvenOdd struct {
	n          int
	zeroToEven chan int
	zeroToOdd  chan int
	toZero     chan bool
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:          n,
		zeroToEven: make(chan int),
		zeroToOdd:  make(chan int),
		toZero:     make(chan bool),
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		printNumber(0)
		if i%2 == 1 {
			z.zeroToOdd <- i
		} else {
			z.zeroToEven <- i
		}
		<-z.toZero
	}
	close(z.zeroToOdd)
	close(z.zeroToEven)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for {
		v, ok := <-z.zeroToEven
		if !ok {
			return
		}

		printNumber(v)
		z.toZero <- true
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for {
		v, ok := <-z.zeroToOdd
		if !ok {
			return
		}

		printNumber(v)
		z.toZero <- true
	}
}
