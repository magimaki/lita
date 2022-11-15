package lita

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	resp, _ := json.Marshal(NewError(ErrCmdInput, "api error test", nil))
	_, _ = w.Write(resp)
	return
}

func TestStartSimpleServer(t *testing.T) {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
