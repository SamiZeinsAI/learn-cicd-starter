package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "ApiKey samizein")
	got, err := auth.GetAPIKey(req.Header)
	if err != nil {
		t.Fatal(err)
	}
	expected := "samizein"
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}
