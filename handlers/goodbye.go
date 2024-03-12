package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("In Goodbye")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong :(", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Goodbye, %s", d)

}
