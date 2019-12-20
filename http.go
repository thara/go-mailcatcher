package main

import (
	"encoding/json"
	"net/http"
)

type httpHandlers struct {
	inbox *Inbox
}

func (h *httpHandlers) listMessages(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(h.inbox)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
