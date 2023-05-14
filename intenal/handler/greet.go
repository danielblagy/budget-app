package handler

import (
	"fmt"
	"net/http"
	"time"
)

func (h handler) Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
