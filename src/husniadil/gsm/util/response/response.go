package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"data":  data,
		"error": nil,
	}
	js, err := json.Marshal(resp)
	if err != nil {
		ResponseFailed(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func ResponseFailed(w http.ResponseWriter, status int, err error) {
	resp := map[string]interface{}{
		"data":  nil,
		"error": fmt.Sprintf("%s", err),
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
