package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/IvanProg00/otus-hw/hw09_struct_validator/validators"
)

var (
	ErrExpectedStructure      = errors.New("expected a structure")
	ErrTypeNotSupports        = errors.New("type not supports")
	ErrIncorrectTagAssertions = errors.New("incorrect tag assertions")
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
	val := reflect.ValueOf(v)
	valType := val.Type()

	if valType.Kind() != reflect.Struct {
		return ErrExpectedStructure
	}

	for i := 0; i < val.NumField(); i++ {
		fieldType := valType.Field(i)
		fieldVal := val.Field(i)

		if fieldVal.CanSet() {
			continue
		}

		if tag, ok := fieldType.Tag.Lookup("validate"); ok {
			params, err := GetValidateParams(tag)
			if err != nil {
				errs = append(errs, ValidationError{
					Field: fieldType.Name,
					Err:   err,
				})
				continue
			}

			err = ValidateField(fieldVal, params)
			if err != nil {
				errs = append(errs, ValidationError{
					Field: fieldType.Name,
					Err:   err,
				})
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

func GetValidateParams(tag string) (map[string]string, error) {
	res := make(map[string]string)

	for _, v := range strings.Split(tag, "|") {
		params := strings.Split(v, ":")
		if len(params) != 2 {
			return nil, ErrIncorrectTagAssertions
		}
		res[params[0]] = params[1]
	}

	return res, nil
}

func ValidateField(fieldVal reflect.Value, params map[string]string) error {
	switch v := fieldVal.Interface().(type) {
	case string:
		return validators.ValidateIfString(v, params)
	case []string:
		return validators.ValidateIfSliceString(v, params)
	case int:
		return validators.ValidateIfInt(v, params)
	case []int:
		return validators.ValidateIfSliceInt(v, params)
	default:
		return ErrTypeNotSupports
	}
}
