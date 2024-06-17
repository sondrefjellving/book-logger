package main


func main() {
	cfg := config{
		books: make([]Book, 0),
	}
	startCLI(&cfg)
}