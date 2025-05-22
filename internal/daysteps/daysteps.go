package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	strings := strings.Split(datastring, ",")
	if len(strings) != 2 {
		return errors.New("dsParse: len(strings) != 2")
	}
	ds.Steps, err = strconv.Atoi(strings[0])
	if err != nil {
		return err
	}
	if ds.Steps < 1 {
		return errors.New("dsParse: steps <= 0")
	}
	ds.Duration, err = time.ParseDuration(strings[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("dsParse: steps <= 0")
	}
	return err
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, spentenergy.Distance(ds.Steps, ds.Personal.Height), calories), err
}
