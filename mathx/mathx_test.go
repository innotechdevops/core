package mathx_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/innotechdevops/core/mathx"
)

func TestRoundToSignificantFigures(t *testing.T) {
	tests := []struct {
		input   float64
		sigFigs int
		//expected float64
		name string
	}{
		{123.456789, 4, "3"},
		{12.3456789, 4, "2"},
		{1.23456789, 4, "1"},
		{0.123456789, 4, "0"},
		{0.0123456789, 4, "0.0"},
		{0.00123456789, 4, "0.00"},
		{0.000123456789, 4, "0.000"},
		//{0.00000123456, 4, 0.000001235, "Very small number"},
		//{0.0123456, 4, 0.01235, "Small decimal"},
		//{123.456, 4, 123.5, "Regular number"},
		//{1234.56, 4, 1235, "Large number"},
		//{0.0, 4, 0.0, "Zero"},
		//{-0.00123456, 3, -0.00123, "Negative number"},
		//{1000000, 3, 1000000, "Round number"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mathx.RoundToSignificantFigures(test.input, test.sigFigs)
			fmt.Printf("input %f output %f\n", test.input, result)
			//if math.Abs(result-test.expected) > 1e-10 {
			//	t.Errorf("Expected %g, got %g", test.expected, result)
			//}
		})
	}
}

func TestRoundToSignificantFiguresPtr(t *testing.T) {
	tests := []struct {
		input    float64
		sigFigs  int
		expected float64
		name     string
	}{
		{0.00000123456, 4, 0.00000123456, "Very small number ptr"},
		{0.0123456, 4, 0.0123456, "Small decimal ptr"},
		{123.456, 4, 123.456, "Regular number ptr"},
		{1234.56, 4, 1234.56, "Large number ptr"},
		{0.0, 4, 0.0, "Zero ptr"},
		{-0.00123456, 3, -0.00123456, "Negative number ptr"},
		{1000000, 3, 1000000, "Round number ptr"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := test.input
			mathx.RoundToSignificantFiguresPtr(&val, test.sigFigs)
			if math.Abs(val-test.expected) > 1e-10 {
				t.Errorf("Expected %g, got %g", test.expected, val)
			}
		})
	}
}

func TestRoundToSignificantFiguresPtr_NilPointer(t *testing.T) {
	// ทดสอบกรณี nil pointer - ไม่ควรเกิด panic
	mathx.RoundToSignificantFiguresPtr(nil, 4)
	// ถ้าไม่ panic แสดงว่าผ่าน
}

func TestFormatSignificantFigures(t *testing.T) {
	tests := []struct {
		input    float64
		sigFigs  int
		expected string
		name     string
	}{
		{0.00000123456, 4, "0.000001235", "Very small formatted"},
		{0.0123456, 4, "0.01235", "Small decimal formatted"},
		{123.456, 4, "123.5", "Regular number formatted"},
		{1000.0, 3, "1000", "Round thousand"},
		{0.0, 4, "0", "Zero formatted"},
		{-123.456, 3, "-123", "Negative formatted"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mathx.FormatSignificantFigures(test.input, test.sigFigs)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestCountSignificantFigures(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		name     string
	}{
		{"123", 3, "Simple integer"},
		{"0.00123", 3, "Leading zeros"},
		{"1.230", 4, "Trailing zero after decimal"},
		{"1200", 4, "Ambiguous trailing zeros"},
		{"0.0", 0, "Single zero"},
		{"102.0", 4, "Zero between digits"},
		{"-123.45", 5, "Negative number"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mathx.CountSignificantFigures(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, got %d for input %s", test.expected, result, test.input)
			}
		})
	}
}

func BenchmarkRoundToSignificantFigures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathx.RoundToSignificantFigures(0.00000123456, 4)
	}
}

func BenchmarkRoundToSignificantFiguresPtr(b *testing.B) {
	val := 0.00000123456
	for i := 0; i < b.N; i++ {
		mathx.RoundToSignificantFiguresPtr(&val, 4)
		val = 0.00000123456 // reset for next iteration
	}
}

func BenchmarkFormatSignificantFigures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathx.FormatSignificantFigures(0.00000123456, 4)
	}
}
