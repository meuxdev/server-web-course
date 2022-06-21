package main

func main() {
	server := NewServer(":8000")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Listen()

}
