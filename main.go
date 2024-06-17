package main

type config struct {
	books []Book
}

func main() {
	cfg := config{
		books: make([]Book, 0),
	}
	startCLI(&cfg)
}