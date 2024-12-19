package main

import (
	"fmt"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
