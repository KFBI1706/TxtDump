package main

import (
	"fmt"
	"net/http"
)

func routerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Im working maayn")
}
