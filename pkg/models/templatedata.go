package models

// TemplateData holds data sent from templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRPToken string
	Flash     string
	Warning   string
	Error     string
}
