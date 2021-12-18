package hw09structvalidator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User1 struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	User2 struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Nums1  []int    `validate:"min:1"`
		Nums2  []int    `validate:"max:120"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in interface{}
	}{
		{
			in: App{
				Version: "4.8.5",
			},
		},
		{
			in: Response{
				Code: 200,
			},
		},
		{
			in: Response{
				Code: 404,
			},
		},
		{
			Response{
				Code: 500,
			},
		},
		{
			User2{
				ID:     "69websrwamsjh68rg7fjcp8qefnczenus84x",
				Name:   "hello",
				Age:    27,
				Email:  "example@gmail.com",
				Nums1:  []int{1, 43, 994, 92},
				Nums2:  []int{120, 10, -40, -120, 41},
				Phones: []string{"98438493843", "49592909478"},
			},
		},
		{
			User2{
				ID:     "69websrwamsjh68rg7fjcp8qefnczenus84x",
				Name:   "hello",
				Age:    27,
				Email:  "example@gmail.com",
				Phones: []string{"98438493843", "49592909478"},
			},
		},
		{
			Token{},
		},
		{
			Token{
				Header:    []byte{'a', 'c', 'd'},
				Payload:   []byte{'p', 'a', 'y', 'l', 'o', 'd'},
				Signature: []byte{'s'},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.NoError(t, err)
		})
	}
}

func TestValidate_incorrectValues(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "Hello",
			expectedErr: ErrExpectedStructure,
		},
		{
			in: App{
				Version: "1.0.",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   errors.New("length not equals 5"),
				},
			},
		},
		{
			in: App{
				Version: "1.0.84",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   errors.New("length not equals 5"),
				},
			},
		},
		{
			in: Response{
				Code: 20,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   errors.New("number not in 200, 404, 500"),
				},
			},
		},
		{
			in: Response{
				Code: 405,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   errors.New("number not in 200, 404, 500"),
				},
			},
		},
		{
			in: User1{
				ID:     "1",
				Name:   "hello",
				Age:    9,
				Email:  "someemail@mail.com",
				Role:   "some role",
				Phones: []string{"493849", "48394"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   errors.New("length not equals 36"),
				},
				ValidationError{
					Field: "Age",
					Err:   errors.New("minimum value is 18"),
				},
				ValidationError{
					Field: "Role",
					Err:   errors.New("type not supports"),
				},
				ValidationError{
					Field: "Phones",
					Err:   errors.New("length not equals 11"),
				},
			},
		},
		{
			in: User2{
				ID:    "t3o5qh8y9j324z8orheg5rauhxx3pyghxce9",
				Name:  "Name",
				Age:   19,
				Email: "incorrect email",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Email",
					Err:   errors.New("value not match the pattern \"^\\w+@\\w+\\.\\w+$\""),
				},
			},
		},
		{
			in: User2{
				ID:    "t3o5qh8y9j324z8orheg5rauhxx3pyghxce9",
				Name:  "Name",
				Age:   19,
				Email: "emailnovalid",
				Nums1: []int{-5, -4, 10, 107},
				Nums2: []int{120, 121},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Email",
					Err:   errors.New("value not match the pattern \"^\\w+@\\w+\\.\\w+$\""),
				},
				ValidationError{
					Field: "Nums1",
					Err:   errors.New("minimum value is 1"),
				},
				ValidationError{
					Field: "Nums2",
					Err:   errors.New("maximum value is 120"),
				},
			},
		},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.in)
			require.Error(t, err)
			require.EqualError(t, err, tt.expectedErr.Error())

			var resErrors ValidationErrors
			ok1 := errors.As(err, &resErrors)
			var expectedErrors ValidationErrors
			ok2 := errors.As(tt.expectedErr, &expectedErrors)

			if ok1 && ok2 {
				for i, e1 := range expectedErrors {
					require.Equal(t, e1.Field, resErrors[i].Field)
					require.Equal(t, e1.Err, resErrors[i].Err)
				}
			}
		})
	}
}
