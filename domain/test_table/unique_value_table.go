package testtable

import (
	"github.com/atharvyadav96k/gcp/app"
)

// TestCase defines the structure for your table-driven test
type TestCase struct {
	Name     string
	Input    interface{}
	Expected []app.UniqueField
}

func UniqueValueTestTable() []TestCase {
	return []TestCase{
		{
			Name: "Single unique field",
			Input: struct {
				Email string `unique:"true" firestore:"email"`
				Name  string
			}{Email: "test@test.com", Name: "John"},
			Expected: []app.UniqueField{
				{Name: "email", Value: "test@test.com"},
			},
		},
		{
			Name: "Multiple unique fields",
			Input: struct {
				Code string `unique:"true" firestore:"code"`
				ID   string `unique:"true" firestore:"id"`
			}{Code: "ABC", ID: "123"},
			Expected: []app.UniqueField{
				{Name: "code", Value: "ABC"},
				{Name: "id", Value: "123"},
			},
		},
		{
			Name: "No unique fields",
			Input: struct {
				Name string
				Age  int
			}{Name: "John", Age: 30},
			Expected: nil,
		},
	}
}
