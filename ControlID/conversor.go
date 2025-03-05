package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func ConvertCard(cardNumber string) string {
	if strings.Contains(cardNumber, ",") {
		return convertWgToCardValue(cardNumber)
	} else if match, _ := regexp.MatchString("[a-fA-F]", cardNumber); match {
		decimalValue, _ := strconv.ParseInt(cardNumber, 16, 64)
		wg := convertDecimalToWg(decimalValue)
		return convertWgToCardValue(wg)
	} else {
		decimalValue, _ := strconv.ParseInt(cardNumber, 10, 64)
		wg := convertDecimalToWg(decimalValue)
		return convertWgToCardValue(wg)
	}
}

func convertWgToCardValue(cardNumber string) string {
	parts := strings.Split(cardNumber, ",")
	area, _ := strconv.ParseInt(parts[0], 10, 64)
	code, _ := strconv.ParseInt(parts[1], 10, 64)

	cardValue := float64(area)*math.Pow(2.0, 32.0) + float64(code)
	return fmt.Sprintf("%.0f", cardValue)
}

func convertDecimalToWg(decimalValue int64) string {
	area := decimalValue / 65536
	code := decimalValue % 65536
	return fmt.Sprintf("%d,%d", area, code)
}
