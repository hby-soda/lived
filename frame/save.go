package frame

import (
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"
	"github.com/liangran2018/lived/plat/home"

	"github.com/gin-gonic/gin"
)

func Save(c *gin.Context) {
	h := human.GetHuman()
	data.User = base.User{Hurt: h.Hurt, Ill: h.Ill, Hungry: h.Hungry, Thirst: h.Thirst, Blood: h.Blood,
		Mood: h.Mood, Wake: h.Wake, Lvl: h.Lvl, Exp: h.Exp, Hero: int(h.Hero)}
	data.GameTime = base.Time{Time: env.GetTimeInt().Time(), Overday: env.GetTimeInt().Overday()}
	data.Weather = int(env.GetWeather())
	data.Temprature = env.GetBaseTemp()

	for k, v := range home.GetOwnBuilding().Own {
		if v.Lvl == 0 {
			continue
		}
		data.OwnBuild[k] = base.OB{Lvl: v.Lvl, Dur: v.Dur}
	}

	for k, v := range materiel.GetOwnThings().OwnProduct() {
		data.OwnProduct[int(k)] = v
	}

	pl := plat.GetPublic()
	for p, q := range pl {
		data.PlatLastTime[int(p)] = q.LastTime
		data.PlatProduct[int(p)] = make(map[int]int, len(q.Product))
		for a, b := range q.Product {
			data.PlatProduct[int(p)][int(a)] = b.Num
		}

		data.PlatAnimal[int(p)] = make(map[int]int, len(q.Animal))
		for c, b := range q.Animal {
			data.PlatAnimal[int(p)][int(c)] = b.Num
		}
	}

	err := base.DeleteFile(fileName)
	if err != nil {
		log.GetLogger().Log(log.Wrong, "Save DeleteFile err: " + err.Error(), fileName)
		base.Output(c, base.DelFileErr, err.Error())
		return
	}

	f := base.NewFile()
	d, _ := json.Marshal(data)
	_, err = f.Write(d)
	if err != nil {
		log.GetLogger().Log(log.Wrong, "Save Write err: " + err.Error(), f)
		base.Output(c, base.SaveFileErr, err.Error())
		return
	}

	log.GetLogger().Log(log.Info, "Saved", data)
	base.Output(c, 0, nil)
	return
}
