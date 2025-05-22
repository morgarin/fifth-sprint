package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// парсинг строки типа "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	strings := strings.Split(datastring, ",")
	if len(strings) != 3 {
		return errors.New("parse: len(strings) != 3")
	}
	t.Steps, err = strconv.Atoi(strings[0])
	if err != nil {
		return err
	}
	if t.Steps < 1 {
		return errors.New("parse: steps <= 0")
	}
	t.Duration, err = time.ParseDuration(strings[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return errors.New("parse: steps <= 0")
	}
	t.TrainingType = strings[1]
	return err
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, t.Height)
	aveSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, aveSpeed, calories), err
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, aveSpeed, calories), err
	default:
		return "", errors.New("ActionInfo: неизвестный тип тренировки")
	}
}
