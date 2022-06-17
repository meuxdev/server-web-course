package main

func main() {
	server := NewServer(":8000")
	server.Handle("/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Listen()

}
