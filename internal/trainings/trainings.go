package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о тренировке.
func (t *Training) Parse(datastring string) error {
	data := strings.Split(datastring, ",")

	if len(data) != 3 {
		return fmt.Errorf("неверный формат данных тренировки")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("некорректное количество шагов")
	}

	t.Steps = steps
	t.TrainingType = strings.TrimSpace(data[1])

	duration, err := time.ParseDuration(strings.TrimSpace(data[2]))
	if err != nil {
		return err
	}

	if duration <= 0 {
		return fmt.Errorf("некорректная продолжительность тренировки")
	}

	t.Duration = duration

	return nil
}

// ActionInfo формирует информацию о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)

	speed := spentenergy.MeanSpeed(
		t.Steps,
		t.Height,
		t.Duration,
	)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)

	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
