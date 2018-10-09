package frame

import (
	"time"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat/home"
	"github.com/liangran2018/lived/surplus"
	"github.com/liangran2018/lived/plat"

	"github.com/gin-gonic/gin"
)

const defaultName = "一一"

func NewGame(c *gin.Context) {
	name := c.Query("name")
	if base.Empty(name) {
		name = defaultName
	}

	//开始时间
	data.StartTime = time.Now().Format("2006-01-02 15:04:05")
	//上次游戏时间
	data.LastTime = time.Now().Format("2006-01-02 15:04:05")
	//玩家昵称
	data.Name = name
	//重置游戏内时间
	env.NewTime()
	//获取天气
	env.NewWeather()
	//获取气温
	env.NewTempToday()
	//新建人物
	human.NewHuman(name)
	//营地起始建筑，只有床
	home.NewOwnBuilding()
	//起始拥有物品
	materiel.NewOwnThings()
	//地图初始化
	plat.NewPublic()
	//记录
	log.GetLogger().Log(log.Info, "Newgame Start, NickName: " + name)

	surplus.MainWin(c)
}
