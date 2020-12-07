package main

func main() {
	b := NewBoard()
	for true {
		b.Print()
		b.Input()
	}
}
