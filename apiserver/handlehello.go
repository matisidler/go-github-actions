package apiserver

import "net/http"

type Blabla struct {
	Skr string
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(("Hello world")))
	}
}
