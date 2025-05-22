package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// потеря калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, errors.New("шагов меньше 1")
	}
	if weight < 1 {
		return 0, errors.New("вес меньше 1")
	}
	if height < 1 {
		return 0, errors.New("рост меньше 1")
	}
	if duration < 1 {
		return 0, errors.New("время тренировки меньше 1")
	}
	return (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

// потеря калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, errors.New("шагов меньше 1")
	}
	if weight < 1 {
		return 0, errors.New("вес меньше 1")
	}
	if height < 1 {
		return 0, errors.New("рост меньше 1")
	}
	if duration < 1 {
		return 0, errors.New("время тренировки меньше 1")
	}
	return (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH, nil
}

// средняя скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// дистанция в км
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}
