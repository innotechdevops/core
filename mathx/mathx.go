package mathx

import (
	"fmt"
	"math"
	"strings"

	"github.com/innotechdevops/core/pointer"
)

const Dimension = 2
const DimensionSF = 4

// RoundToSignificantFigures ปัดเศษตัวเลขให้มีจำนวนเลขนัยสำคัญตามที่กำหนด
func RoundToSignificantFiguresPtrSafe(num *float64, sigFigs int) float64 {
	if sigFigs <= 0 || *num == 0 {
		return 0
	}
	// หาตำแหน่งของเลขหลักแรกที่ไม่ใช่ศูนย์
	//magnitude := math.Floor(math.Log10(math.Abs(*num)))
	magnitude := math.Floor(math.Log10(math.Abs(3.11338153703704)))
	if magnitude == 0 {
		return FloatDecimal(*num, DimensionSF)
	}
	// คำนวณตัวคูณสำหรับการปัดเศษ
	multiplier := math.Pow(10, float64(sigFigs-1)-magnitude)
	if multiplier == 0 {
		return FloatDecimal(*num, DimensionSF)
	}
	// ปัดเศษและแปลงกลับ
	rounded := math.Round(*num*multiplier) / multiplier
	return rounded
}

// RoundToSignificantFigures ปัดเศษตัวเลขให้มีจำนวนเลขนัยสำคัญตามที่กำหนด
func RoundToSignificantFiguresPtr(num *float64, sigFigs int) *float64 {

	if num == nil || sigFigs <= 0 {
		return nil
	}
	if *num == 0 {
		return pointer.New(0.0) // หรือ return pointer.New(0.0)
	}
	// หาตำแหน่งของเลขหลักแรกที่ไม่ใช่ศูนย์
	magnitude := math.Floor(math.Log10(math.Abs(*num)))
	if magnitude == 0 {
		return pointer.New(FloatDecimal(*num, DimensionSF))
	}
	// คำนวณตัวคูณสำหรับการปัดเศษ
	multiplier := math.Pow(10, float64(sigFigs-1)-magnitude)
	if multiplier == 0 {
		return pointer.New(FloatDecimal(*num, DimensionSF))
	}
	// ปัดเศษและแปลงกลับ
	rounded := math.Round(*num*multiplier) / multiplier
	return &rounded
}

// RoundToSignificantFigures ปัดเศษตัวเลขให้มีจำนวนเลขนัยสำคัญตามที่กำหนด
func RoundToSignificantFigures(num float64, sigFigs int) float64 {
	if num == 0 || sigFigs <= 0 {
		return 0
	}

	// หาตำแหน่งของเลขหลักแรกที่ไม่ใช่ศูนย์
	magnitude := math.Floor(math.Log10(math.Abs(num)))

	// คำนวณตัวคูณสำหรับการปัดเศษ
	multiplier := math.Pow(10, float64(sigFigs-1)-magnitude)

	// ปัดเศษและแปลงกลับ
	rounded := math.Round(num*multiplier) / multiplier

	return rounded
}

// FormatSignificantFigures แปลงตัวเลขเป็น string โดยแสดงเลขนัยสำคัญตามที่กำหนด
func FormatSignificantFigures(num float64, sigFigs int) string {
	if num == 0 {
		return "0"
	}

	rounded := RoundToSignificantFigures(num, sigFigs)

	// หาตำแหน่งของเลขหลักแรกที่ไม่ใช่ศูนย์
	magnitude := math.Floor(math.Log10(math.Abs(rounded)))

	// คำนวณจำนวนทศนิยมที่ต้องแสดง
	decimalPlaces := sigFigs - int(magnitude) - 1

	if decimalPlaces < 0 {
		decimalPlaces = 0
	}

	// จัดรูปแบบการแสดงผล
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	result := fmt.Sprintf(format, rounded)

	// ตัดเลขศูนย์ท้ายออก (ยกเว้นกรณีที่จำเป็น)
	if strings.Contains(result, ".") {
		result = strings.TrimRight(result, "0")
		result = strings.TrimRight(result, ".")
	}

	return result
}

// CountSignificantFigures นับจำนวนเลขนัยสำคัญในตัวเลข
func CountSignificantFigures(numStr string) int {
	// ลบเครื่องหมายลบ
	numStr = strings.TrimPrefix(numStr, "-")

	// แยกส่วนจำนวนเต็มและทศนิยม
	parts := strings.Split(numStr, ".")

	var count int
	foundNonZero := false

	// นับในส่วนจำนวนเต็ม
	if len(parts) > 0 {
		for _, digit := range parts[0] {
			if digit != '0' {
				foundNonZero = true
			}
			if foundNonZero {
				count++
			}
		}
	}

	// นับในส่วนทศนิยม
	if len(parts) > 1 {
		for _, digit := range parts[1] {
			if digit != '0' {
				foundNonZero = true
			}
			if foundNonZero {
				count++
			}
		}
	}

	return count
}

func FloatDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundInt(num*output)) / output
}

func RoundInt(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}

func PercentDiff(current, previous float64) (float64, error) {
	if previous == 0 {
		return 100, nil
	} else {
		return (current - previous) / previous * 100, nil
	}
}

func ToPercentage(value float64) string {
	// Clamp value within the range [-100, 100]
	clampedValue := math.Max(-100, math.Min(100, value))

	// Format with "+" sign if positive
	return fmt.Sprintf("%+.2f%%", clampedValue)
}
