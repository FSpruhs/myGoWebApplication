package forms

import (
	"fmt"
	"net/url"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if value == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MinLength(field string, length int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must have min length %d", length))
	}
}

func (f *Form) IsEmail(field string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !govalidator.IsEmail(value) {
		f.Errors.Add(field, "Invalid email address.")
	}
}
