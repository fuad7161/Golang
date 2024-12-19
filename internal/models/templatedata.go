package models

import "github.com/fuad7161/Golang/tree/Project/Bookings/internal/forms"

// TemplateData holds data sent from templates
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
