package plat

import (
	"time"
	"math/rand"
)

func (this *Nature) Rand() (product, animal, count int) {
	total := 0
	for _, v := range this.Product {
		total += v.Num
	}

	for _, v := range this.Animal {
		total += v.Num
	}

	if total == 0 {
		return 0, 0, 0
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(total)

	for k, v := range this.Product {
		if i <= v.Num {
			return int(k), 0, rand.Intn(v.Num)
		}

		i -= v.Num
	}

	for j, v := range this.Animal {
		if i <= v.Num {
			return 0, int(j), 1
		}

		i -= v.Num
	}

	return 0, 0, 0
}
