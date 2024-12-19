package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mH myHandler
	h := NoSurf(&mH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mH myHandler
	h := SessionLoad(&mH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}
