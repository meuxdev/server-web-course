package main

func main() {
	server := NewServer(":8000")
	server.Handle("/", server.AddMiddleware(HandleRoot, PrintRequest()))
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), PrintRequest()))
	server.Listen()

}
