package rfc7807

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RFC7807 is a struct that represents the RFC7807 standard.
type RFC7807 struct {
	Type       string         `json:"type,omitempty"`
	Title      string         `json:"title"`
	Status     int            `json:"status,omitempty"`
	Detail     string         `json:"detail,omitempty"`
	Instance   string         `json:"instance,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
}

// SimpleResponse writes a simple response with the RFC7807 struct to the provided writer.
func SimpleResponse(w http.ResponseWriter, status int, title string) error {
	e := NewError(fmt.Sprintf("%s: %s", statusCodes[status], title)).SetStatus(status)
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(e.Status)
	return json.NewEncoder(w).Encode(e)
}

// NewError returns a new RFC7807 struct.
func NewError(title string) *RFC7807 {
	return &RFC7807{
		Title: title,
	}
}

// SetType sets the type field of the RFC7807 struct.
func (e *RFC7807) SetType(t string) *RFC7807 {
	e.Type = t
	return e
}

// SetStatus sets the status field of the RFC7807 struct.
func (e *RFC7807) SetStatus(s int) *RFC7807 {
	e.Status = s
	return e
}

// SetDetail sets the detail field of the RFC7807 struct.
func (e *RFC7807) SetDetail(d string) *RFC7807 {
	e.Detail = d
	return e
}

// SetInstance sets the instance field of the RFC7807 struct.
func (e *RFC7807) SetInstance(i string) *RFC7807 {
	e.Instance = i
	return e
}

// SetExtensions sets the extensions field of the RFC7807 struct.
func (e *RFC7807) SetExtensions(ext map[string]any) *RFC7807 {
	e.Extensions = ext
	return e
}
