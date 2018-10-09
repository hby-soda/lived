package plat

import (
	"github.com/liangran2018/lived/materiel"
)

type Tc struct {
	H  int `json:"h"`
	Mi int `json:"mi"`
}

type animalNature struct {
	meat0    int
	meat1    int
	meat2    int
	blood    int
	attack   int
	defend   int
	dodge    int
	critical int
}

type Nature struct {
	Time     Tc                          `json:"time"`
	LastTime int                         `json:"lastTime"`
	Product  map[materiel.Product]*count `json:"product"`
	Animal   map[materiel.Animal]*count  `json:"animal"`
	Lvl      int                         `json:"lvl"`
}

type count struct {
	Num      int		`json:"num"`
	maxNum   int
	Increase float64	`json:"increase"`
}

type Place int

const (
	lsd Place = iota
	gmc

	undefined
)

var placeName []string
var public map[Place]*Nature

func init() {
	publicInit()
}

func NewPublic() {
	riprapNew()
	bushNew()
}

func LoadPublic(m map[int]int) {
	for k, v := range m {
		public[Place(k)].LastTime = v
	}
}

func LoadProduct(m map[int]map[int]int) {
	for k, v := range m {
		for a, b := range v {
			public[Place(k)].Product[materiel.Product(a)].Num = b
		}
	}
}

func LoadAnimal(m map[int]map[int]int) {
	for k, v := range m {
		for a, b := range v {
			public[Place(k)].Animal[materiel.Animal(a)].Num = b
		}
	}
}

func GetPublic() map[Place]*Nature {
	return public
}

func (this *Nature) PublicIncrease(i float64) {
	for _, v := range this.Product {
		if v.Num >= v.maxNum {
			continue
		}

		v.Num += int(i * v.Increase)
		if v.Num > v.maxNum {
			v.Num = v.maxNum
		}
	}

	for _, v := range this.Animal {
		if v.Num >= v.maxNum {
			continue
		}

		v.Num += int(i * v.Increase)
		if v.Num > v.maxNum {
			v.Num = v.maxNum
		}
	}
}

func (this Place) Name() string {
	return placeName[int(this)]
}

func (this Place) Nature() *Nature {
	return public[this]
}

func publicInit() {
	placeName = make([]string, undefined)
	placeName = []string{"乱石堆", "灌木丛"}

	public = make(map[Place]*Nature, undefined)
	riprapInit()
	bushInit()
}
