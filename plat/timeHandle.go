package plat

import (
	"github.com/liangran2018/lived/env"
)

func (this *Nature) SetLasttime(i int) float64 {
	if this.LastTime == 0 {
		this.LastTime = i
		return 0
	}

	newTime := env.Int2Time(i)
	oldTime := env.Int2Time(this.LastTime)
	this.LastTime = i

	var plus float64 = 0
	if newTime.Year > oldTime.Year {
		plus += 365
		if newTime.Month < oldTime.Month {
			for j:=newTime.Month; j<oldTime.Month; j++ {
				plus -= float64(env.MonthDay[j])
			}
		}
	}

	if newTime.Month > oldTime.Month {
		for j:=oldTime.Month; j<newTime.Month; j++ {
			plus += float64(env.MonthDay[j])
		}
	}

	if newTime.Day != oldTime.Day {
		if newTime.Day < oldTime.Day {
			plus -= float64(oldTime.Day - newTime.Day)
		} else {
			plus += float64(newTime.Day - oldTime.Day)
		}
	}

	plus = plus * 24
	if newTime.Hour != oldTime.Hour {
		if newTime.Hour < oldTime.Hour {
			plus -= float64(oldTime.Hour - newTime.Hour)
		} else {
			plus += float64(newTime.Hour - oldTime.Hour)
		}
	}


	if newTime.Minute != oldTime.Minute {
		if newTime.Minute < oldTime.Minute {
			plus -= float64(oldTime.Minute - newTime.Minute)/60
		} else {
			plus += float64(newTime.Minute - oldTime.Minute)/60
		}
	}
	return plus
}
