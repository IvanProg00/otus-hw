package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/IvanProg00/otus-hw/hw09_struct_validator/validators"
)

var (
	ErrExpectedStructure = errors.New("expected a structure")
	ErrTypeNotSupports   = errors.New("type not supports")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	res := strings.Builder{}

	for i, val := range v {
		if i > 0 {
			res.WriteByte('\n')
		}
		res.WriteString(fmt.Sprintf("incorrect parameter \"%s\": %s", val.Field, val.Err.Error()))
	}

	return res.String()
}

func Validate(v interface{}) error {
	errs := make(ValidationErrors, 0)

	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return ErrExpectedStructure
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag, ok := field.Tag.Lookup("validate"); ok {
			valParams := make(map[string]string)
			for _, v := range strings.Split(tag, "|") {
				params := strings.Split(v, ":")
				valParams[params[0]] = params[1]
			}

			val := reflect.ValueOf(v).Field(i)
			switch v := val.Interface().(type) {
			case string:
				if err := validators.ValidateIfString(v, valParams); err != nil {
					errs = append(errs, ValidationError{
						Field: field.Name,
						Err:   err,
					})
				}
			case []string:
				if err := validators.ValidateIfSliceString(v, valParams); err != nil {
					errs = append(errs, ValidationError{
						Field: field.Name,
						Err:   err,
					})
				}
			case int:
				if err := validators.ValidateIfInt(v, valParams); err != nil {
					errs = append(errs, ValidationError{
						Field: field.Name,
						Err:   err,
					})
				}
			case []int:
				if err := validators.ValidateIfSliceInt(v, valParams); err != nil {
					errs = append(errs, ValidationError{
						Field: field.Name,
						Err:   err,
					})
				}
			default:
				errs = append(errs, ValidationError{
					Field: field.Name,
					Err:   ErrTypeNotSupports,
				})
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
