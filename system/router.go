package system

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/frame"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/explore"
	"github.com/liangran2018/lived/plat/home"
	"github.com/liangran2018/lived/surplus"

	"github.com/gin-gonic/gin"
)

func AddRoute(app *gin.Engine) {
	// 系统相关
	app.GET("/api/main", surplus.MainWin)
	app.GET("/api/newgame", frame.NewGame)
	app.GET("/api/getbackup", frame.GetBackup)
	app.GET("/api/loadbackup", frame.ChooseBackup)
	app.GET("/api/save", frame.Save)

	app.GET("/api/ownthingshow", materiel.Show)
	app.POST("/api/ownthingadd", materiel.Add)
	app.POST("/api/ownthingplus", materiel.Plus)

	app.GET("/api/equipnotice", explore.EquipNotice)
	app.POST("/api/equip", explore.Equip)
	app.GET("/api/equipshow", explore.Show)

	app.GET("/api/bagnotice", explore.BagNotice)
	app.POST("/api/bagchoose", explore.BagChoose)
	app.GET("/api/bagshow", explore.BagShow)

	app.GET("/api/buildshow", home.Show)
	app.GET("/api/buildupdatenotice", home.Notice)
	app.GET("/api/buildupdate", home.Update)
	app.GET("/api/actionnotice", home.ActionNotice)
	app.GET("/api/actionchoose", home.ActionChoose)
	app.GET("/api/actioncheck", home.ActionCheck)
	app.GET("/api/harvest", home.Harvest)

	app.GET("/api/platshow", explore.PlatShow)
	app.GET("/api/platdetail", explore.Detail)
	app.GET("/api/platgo", explore.Go)

	app.GET("/api/fight", explore.Fight)

	app.GET("/test", test)
	// 无路由或无方法
	app.NoRoute(error)
	app.NoMethod(error)
}

func test(c *gin.Context) {
	base.Output(c, 0, "test")
}

func error(c *gin.Context) {
	base.Output(c, 400, "access denied")
}
