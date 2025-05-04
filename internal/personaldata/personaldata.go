package personaldata

import "fmt"

type Personal struct {
	Name           string
	Weight, Height float64
}

func (p Personal) Print() {
	str := fmt.Sprintf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
	fmt.Print(str)
}
