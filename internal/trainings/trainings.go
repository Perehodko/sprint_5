package trainings

import ( 
	"time"
	"strings"
	"fmt"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	Personal personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("invalid input")
	}

	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("invalid steps")
	}
	
	if steps <=0 {
		return fmt.Errorf("invalid steps")
	}
	
	t.Steps=steps
	t.TrainingType=parts[1]

	durationStr := parts[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration")
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration")
	}

	t.Duration=duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t == (Training{}) {
		return "", fmt.Errorf("invalid input: Training is empty")
	}

	if t.Steps <= 0 {
		return "", fmt.Errorf("invalid input: steps must be greater than 0")
	}

	if t.Personal.Height <= 0 {
		return "", fmt.Errorf("invalid input: height must be greater than 0")
	}

	if t.Personal.Weight <= 0 {
		return "", fmt.Errorf("invalid input: weight must be greater than 0")
	}

	if t.Duration <= 0 {
		return "", fmt.Errorf("invalid input: duration must be greater than 0")
	}


	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	spentKkal := 0.0
	var err error
	switch t.TrainingType {
	case "Ходьба":
		spentKkal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		spentKkal, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, spentKkal)
	return str, nil
}
