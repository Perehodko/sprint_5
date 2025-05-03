package daysteps

import ( 
	"time"
	"strings"
	"fmt"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)


type DaySteps struct {
	Steps int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid input")
	}

	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("invalid steps")
	}

	ds.Steps=steps

	durationStr := parts[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration")
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration")
	}

	ds.Duration=duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps < 0 {
		return "", fmt.Errorf("invalid steps")
	}

	if ds.Duration <= 0 {
		return "", fmt.Errorf("invalid duration")
	}

	if ds.Personal.Weight <= 0 {
		return "", fmt.Errorf("invalid weight")
	}

	if ds.Personal.Height <= 0 {
		return "", fmt.Errorf("invalid height")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	spentKkal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentKkal)

	return str, nil
}
