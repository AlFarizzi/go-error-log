package errjson

import (
	"errors"
	"testing"
)

func TestErrJson(t *testing.T) {
	err := errors.New("Ini Error Pertama")

	WriteError("err.json", err.Error())
}
