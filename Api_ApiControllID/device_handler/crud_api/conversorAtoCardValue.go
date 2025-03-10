package crudapi

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// ConvertCard converte o número do cartão baseado no formato
func convertCard(cardNumber string) string {
	if strings.Contains(cardNumber, ",") {
		return convertWgToCardValue(cardNumber)
	} else if match, _ := regexp.MatchString("[a-fA-F]", cardNumber); match {
		// Se for HEX, converte para Decimal e depois para Wiegand
		decimalValue, _ := strconv.ParseInt(cardNumber, 16, 64)
		wg := convertDecimalToWg(decimalValue)
		return convertWgToCardValue(wg)
	} else {
		// Se for Decimal, converte para Wiegand e depois para CardValue
		decimalValue, _ := strconv.ParseInt(cardNumber, 10, 64)
		wg := convertDecimalToWg(decimalValue)
		return convertWgToCardValue(wg)
	}
}

// convertWgToCardValue converte o número Wiegand para Card Value
func convertWgToCardValue(cardNumber string) string {
	parts := strings.Split(cardNumber, ",")
	area, _ := strconv.ParseInt(parts[0], 10, 64)
	code, _ := strconv.ParseInt(parts[1], 10, 64)

	// Calcula o valor do cartão
	cardValue := float64(area)*math.Pow(2.0, 32.0) + float64(code)
	return fmt.Sprintf("%.0f", cardValue)
}

// convertDecimalToWg converte um valor decimal para o formato Wiegand
func convertDecimalToWg(decimalValue int64) string {
	area := decimalValue / 65536
	code := decimalValue % 65536
	return fmt.Sprintf("%d,%d", area, code)
}
