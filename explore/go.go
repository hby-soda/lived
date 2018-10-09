package explore

import (
	"math/rand"
	"time"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"

	"github.com/gin-gonic/gin"
)

type outputExplore struct {
	Type      string `json:"type"`
	Encounter int    `json:"encounter"`
	Count     int    `json:"count"`
}

func Go(c *gin.Context) {
	p, ok := getPlace(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	env.GetTime().Add(p.Nature().Time.H, p.Nature().Time.Mi)

	if t := p.Nature().SetLasttime(env.GetTimeInt().Time()); t != 0 {
		p.Nature().PublicIncrease(t)
	}

	base.Output(c, 0, struct {
		Product interface{} `json:"product"`
		Animal  interface{} `json:"animal"`
	}{Product: p.Nature().Product, Animal: p.Nature().Animal})
	return
}

func Explore(c *gin.Context) {
	p, ok := getPlace(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	env.GetTime().Add(0, timeRand())

	product, animal, count := p.Nature().Rand()
	if product == 0 && animal == 0 {
		base.Output(c, base.PlaceNothing, nil)
		return
	}

	ope := outputExplore{Encounter:product, Count:count}

	if product != 0 {
		ope.Type = "product"
	} else {
		ope.Type = "animal"
	}

	base.Output(c, 0, ope)
	return
}

func timeRand() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(30)
	if 0 <= i && i < 10 {
		return 10
	} else if 10 <= i && 20 > i {
		return 20
	} else {
		return 30
	}
}
