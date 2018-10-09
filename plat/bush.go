package plat

import (
	"github.com/liangran2018/lived/materiel"
)

func bushInit() {
	public[gmc] = &Nature{Time:Tc{H:1, Mi:0}, Lvl:1}

	public[gmc].Product = map[materiel.Product]*count {
		materiel.Wood:&count{maxNum:11, Increase:0.1},
		materiel.Pebble:&count{maxNum:12, Increase:0.33},
		materiel.Jujube:&count{maxNum:5, Increase:0.07},
	}

	public[gmc].Animal = map[materiel.Animal]*count {
		materiel.Snake:&count{maxNum:14, Increase:0.11},
	}
}

func bushNew() {
	public[gmc].Product[materiel.Wood].Num = 11
	public[gmc].Product[materiel.Pebble].Num = 8
	public[gmc].Product[materiel.Jujube].Num = 1
	public[gmc].Animal[materiel.Snake].Num = 12
}
