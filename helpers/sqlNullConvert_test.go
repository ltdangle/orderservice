package helpers_test

import (
	"database/sql"
	"orders/helpers"
	"testing"
)

func Test_SqlNullConvert(t *testing.T) {
	want := "a string"
	nlstr := sql.NullString{String: want, Valid: true}
	got := helpers.NullStringToString(nlstr)
	if want != got {
		t.Fatalf("%s != %s", want, got)
	}
}
