package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("Get Request")
	case "DELETE":
		println("Delete Request")
	case "POST":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("Unhandled method")
	}
}
