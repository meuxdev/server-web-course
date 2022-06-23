package main

func main() {
	server := NewServer(":8000")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logger()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Handle("POST", "/api/user", server.AddMiddleware(UserPostRequest, Logger()))
	server.Listen()

}
