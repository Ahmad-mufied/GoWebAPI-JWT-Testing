package main

import (
	"net/url"
	"strings"
)

type erros map[string][]string

func (e erros) Get(field string) string {
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}
	return  errorSlice[0]
}

func (e erros) Add(field, message string) {
	e[field] = append(e[field], message)
}


type Form struct {
	Data url.Values
	Errors erros
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data: data,
		Errors: map[string][]string{},
	}
}

func (f *Form) Has(field string) bool{
	x := f.Data.Get(field)
	if x == "" {
		return false
	}
	return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Check(ok bool, ke, message string) {
	if !ok {
		f.Errors.Add(ke, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}



