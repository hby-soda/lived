package explore

import (
	"github.com/liangran2018/lived/materiel"
)

type equipChoose struct {
	name   materiel.Product
	detail []int
}

func init() {
	NewEquip()
	NewBag()
}
