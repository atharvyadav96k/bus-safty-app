package test

import (
	"reflect"
	"testing"

	"github.com/atharvyadav96k/gcp/app"
	testtable "github.com/atharvyadav96k/gcp/domain/test_table"
)

func TestUniqueValues(t *testing.T) {
	table := testtable.UniqueValueTestTable()

	for _, tc := range table {
		got := app.GetUniqueFields(tc.Input)
		if !reflect.DeepEqual(got, tc.Expected) {
			t.Errorf("\nFAILED: %s\nEXPECTED: %+v\nGOT:      %+v", tc.Name, tc.Expected, got)
		}
	}
}
