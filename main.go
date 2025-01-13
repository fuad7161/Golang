package main

func main() {

	addr := ":8080"
	s := NewAPIServer(addr)
	s.Run()
}
