package plat

import (
	"github.com/liangran2018/lived/materiel"
)

func riprapInit() {
	public[lsd] = &Nature{Time:Tc{H:0, Mi:40}, Lvl:1}

	public[lsd].Product = map[materiel.Product]*count {
		materiel.Stone:&count{maxNum:10, Increase:0.17},
		materiel.Pebble:&count{maxNum:18, Increase:0.5},
	}

	public[lsd].Animal = map[materiel.Animal]*count {
		materiel.WeakWolf:&count{maxNum:3, Increase:0.07},
		materiel.Snake:&count{maxNum:9, Increase:0.09},
	}
}

func riprapNew() {
	public[lsd].Product[materiel.Stone].Num = 10
	public[lsd].Product[materiel.Pebble].Num = 15
	public[lsd].Animal[materiel.WeakWolf].Num = 2
	public[lsd].Animal[materiel.Snake].Num = 5
}
