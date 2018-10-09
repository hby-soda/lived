package explore

import (
	"fmt"
	"time"
	"math/rand"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

func Fight(c *gin.Context) {
	animal := c.Query("animal")
	if base.Empty(animal) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, err := base.IntVal(animal)
	if err != nil || i < 0 || i >= int(materiel.AnimalEnd) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	a := materiel.Animal(i)

	str := make([]string, 0)

	an := a.Hot()
	p := human.GetHuman()
	e := exploreEquip.e
	ab := an.Blood
	ah := p.Hurt
	arrow := exploreBag.product[materiel.Arrow]

	distance := 5

	for {
		str = append(str, fmt.Sprintf("你有%d血量，%s有%d血量, 距离%d", p.Blood, a.Name(), ab, distance))

		if e[0].EquipHot().Attack != 0 && distance > e[0].EquipHot().Distance {
			distance--
			continue
		}

		if e[0].EquipHot().Attack == 0 {
			if e[1].EquipHot().Attack != 0 {
				if distance > e[1].EquipHot().Distance {
					distance--
					continue
				}
			} else if distance > 1 {
				distance--
				continue
			}
		}

		str = append(str, "你发起进攻")

		if die := attack(true, distance, p.Attack, p.Critical, e[0].EquipHot(),
			&arrow, e[1].EquipHot(), &ab, &ah, an.Defend, an.Dodge, materiel.Fight{}, &str); die {
			str = append(str, a.Name()+"死亡，捕猎成功")
			break
		}

		if distance > 1 {
			distance--
			continue
		}

		str = append(str, fmt.Sprintf("你有%d血量，%s有%d血量", p.Blood, a.Name(), ab))
		str = append(str, a.Name()+"发起反击")

		if die := attack(false, distance, an.Attack, an.Critical, materiel.Fight{}, nil,
			materiel.Fight{}, &p.Blood, nil, p.Defend, p.Dodge, e[2].EquipHot(), &str); die {
			base.Output(c, 0, str)
			panic(base.DEAD{Reason: fmt.Sprintf("捕猎被%s打死", a.Name())})
		}
	}

	p.ExpAdd(a.Exp())

	log.GetLogger().Log(log.Info, "fight", str)

	base.Output(c, 0, struct {
		Log []string `json:"log"`
		Get map[materiel.Product]int `json:"get"`
	}{Log:str, Get:map[materiel.Product]int {materiel.Meat:a.Meat()}})
	return
}

// 距离，进攻者攻击、暴击，远程武器攻击、暴击、数量（箭），近战武器攻击、暴击，防守者血量、防御、闪避，护甲防御、闪避
// 日志记录（返回防守者剩余血量，是否死亡）
func attack(humanAttack bool, distance, attackerAtt, attackerCri int, remote materiel.Fight, remoteNum *int,
	melee materiel.Fight, defenderBld, defenderHurt *int, defenderDef, defenderDod int, armor materiel.Fight, str *[]string) bool {
	// 如果血量为0，直接死亡
	if *defenderBld <= 0 {
		return true
	}

	rand.Seed(time.Now().UnixNano())

	// 装备远程武器，并且有箭，开始远程攻击
	if humanAttack && (remote.Attack != 0 && *remoteNum > 0 && distance <= remote.Distance) {
		*remoteNum--
		*str = append(*str, "你射了一箭")
		// 防守者闪避不为0，并且闪避过去，直接到近战攻击阶段
		if defenderDod != 0 && rand.Intn(10) < defenderDod {
			*str = append(*str, "但是被躲避过去")
			goto melee
		}

		/* 人攻击时，动物没有护甲
		// 防守者无闪避或闪避失败，护甲有闪避并成功时，到近战攻击
		if armorDod != 0 && rand.Intn(10) < armorDod {
			*str = append(*str, "但是被躲避过去")
			goto melee
		}
		*/
		// 伤害百分比 1/（1+0.06*def），攻击和防御可线性叠加
		remoteHurt := damage(float64(attackerAtt+remote.Attack), float64(defenderDef))
		*str = append(*str, fmt.Sprintf("造成伤害%d", remoteHurt))
		*defenderBld -= remoteHurt
		if *defenderBld <= 0 {
			return true
		}

		// 进攻者暴击
		if attackerCri != 0 && rand.Intn(10) < attackerCri {
			*str = append(*str, fmt.Sprintf("暴击造成伤害%d", remoteHurt))
			*defenderBld -= remoteHurt
			if *defenderBld <= 0 {
				return true
			}

			goto melee
		}

		// 进攻者无暴击，判断武器暴击
		if remote.Critical != 0 && rand.Intn(10) < remote.Critical {
			*str = append(*str, fmt.Sprintf("暴击造成伤害%d", remoteHurt))
			*defenderBld -= remoteHurt
			if *defenderBld <= 0 {
				return true
			}
		}
	}
melee:
	if melee.Attack != 0 && distance > melee.Distance {
		return false
	}

	var noMelee bool

	// 距离不够
	if melee.Attack == 0 {
		if distance > 1 {
			return false
		} else {
			noMelee = true
		}
	}

	if humanAttack {
		if noMelee {
			*str = append(*str, "你挥了一下拳头")
		} else {
			*str = append(*str, "你挥了一下武器")
		}
	} else {
		*str = append(*str, "咬了你一口")
	}

	// 近战阶段
	if defenderDod != 0 && rand.Intn(10) < defenderDod {
		*str = append(*str, "但是被躲避过去")
		return false
	}

	if !humanAttack && (armor.Dodge != 0 && rand.Intn(10) < armor.Dodge) {
		*str = append(*str, "但是被躲避过去")
		return false
	}

	var meleeHurt int
	if humanAttack {
		meleeHurt = damage(float64(attackerAtt+melee.Attack), float64(defenderDef))
	} else {
		meleeHurt = damage(float64(attackerAtt), float64(defenderDef+armor.Defend))
		*defenderHurt -= 4
	}

	*str = append(*str, fmt.Sprintf("造成伤害%d", meleeHurt))
	*defenderBld -= meleeHurt

	if *defenderBld <= 0 {
		return true
	}

	if attackerCri != 0 && rand.Intn(10) < attackerCri {
		*str = append(*str, fmt.Sprintf("暴击造成伤害%d", meleeHurt))
		*defenderBld -= meleeHurt
		if *defenderBld <= 0 {
			return true
		} else {
			return false
		}
	}

	if humanAttack && (melee.Critical != 0 && rand.Intn(10) < melee.Critical) {
		*str = append(*str, fmt.Sprintf("暴击造成伤害%d", meleeHurt))
		*defenderBld -= meleeHurt
		if *defenderBld <= 0 {
			return true
		} else {
			return false
		}
	}

	return false
}

func damage(att, def float64) int {
	return int(att / (1 + 0.06*def))
}
