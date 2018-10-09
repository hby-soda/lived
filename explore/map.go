package explore

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"

	"github.com/gin-gonic/gin"
)

type outputDetail struct {
	NeedTime plat.Tc            `json:"needTime"`
	Product  []materiel.Product `json:"product"`
	Animal   []materiel.Animal  `json:"animal"`
	Lvl      int                `json:"lvl"`
}

func PlatShow(c *gin.Context) {
	place := plat.GetPublic()

	p := make([]string, len(place))
	i := 0
	for k, _ := range place {
		p[i] = k.Name()
		i++
	}

	base.Output(c, 0, p)
	return
}

func Detail(c *gin.Context) {
	p, ok := getPlace(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	public := plat.GetPublic()

	opd := outputDetail{NeedTime:public[p].Time, Lvl:public[p].Lvl}
	opd.Product = make([]materiel.Product, len(public[p].Product))
	opd.Animal = make([]materiel.Animal, len(public[p].Animal))

	i := 0
	for k, _ := range public[p].Product {
		opd.Product[i] = k
		i++
	}

	i = 0
	for j, _ := range public[p].Animal {
		opd.Animal[i] = j
		i++
	}

	base.Output(c, 0, opd)
	return
}

func getPlace(c *gin.Context) (plat.Place, bool) {
	p := c.Query("place")
	if base.Empty(p) {
		return 0, false
	}

	i, err := base.IntVal(p)
	if err != nil || i < 0 || i >= len(plat.GetPublic()) {
		return 0, false
	}

	return plat.Place(i), true
}
