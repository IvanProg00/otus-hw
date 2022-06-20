package internalhttp

import (
	"fmt"
	"net/http"
)

func (s *Server) helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world")
	req.Body.Close()
}
