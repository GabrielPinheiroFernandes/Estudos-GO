package utils_test

import (
	"APIControlID/utils"
	"testing"
)

func TestConvertCard(t *testing.T) {
	cardValueWg:="010,54467"
	expected:="42949727427"
	result:=utils.ConvertCard(cardValueWg)
	if result != expected {
		t.Error("Valor de cartão WG invalido")
	}
}
