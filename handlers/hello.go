package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("In Hello")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong :(", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Welcome, this is %s", d)

}
