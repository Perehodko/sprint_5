package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, fmt.Errorf("invalid input parameters: value <= 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	kkal := ((duration.Minutes() * weight * avgSpeed) / 60) * walkingCaloriesCoefficient
	return kkal, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration.Minutes() <= 0 {
		return 0.0, fmt.Errorf("invalid input parameters: value < 0")
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	kkal := (weight * avgSpeed * duration.Minutes()) / minInH
	return kkal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0.0
	}

	if duration <= 0 {
		return 0.0
	}

	dist := Distance(steps, height)
	avrSpeed := dist / duration.Hours()
	return avrSpeed
}

func Distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	return (float64(steps) * stepLen) / mInKm
}
