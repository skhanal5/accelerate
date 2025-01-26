package handler

import (
	"io"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Healthy status :)")
}