package main

func main() {
	server := NewServer(":8000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()

}
