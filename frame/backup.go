package frame

import (
	"os"
	"time"
	"strings"
	"io/ioutil"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

type backupInfo struct {
	FileName string `json:"filename"`
	Name string `json:"name"`
	OverDay int `json:"overday"`
	StartTime string `json:"starttime"`
	LastTime string `json:"lasttime"`
}

var backup []*backupInfo
var fileName string

func backupInit() {
	backup = make([]*backupInfo, 0)
}

func backupClear() {
	backup = backup[0:0]
}

func GetBackup(c *gin.Context) {
	backupInit()
	getBackup()
	if len(backup) == 0 {
		base.Output(c, base.NoBackupFile, "没有存档")
		return
	}

	base.Output(c, 0, backup)
	return
}

func ChooseBackup(c *gin.Context) {
	b := c.Query("backup")
	if base.Empty(b) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, err := base.IntVal(b)
	if err != nil {
		base.Output(c, base.TypeConvertErr, err.Error())
		return
	}

	if i < 0 || i >= len(backup) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	fileName = backup[i].FileName + ".json"
	d, err := loadFile(fileName)
	if err != nil {
		base.Output(c, base.OpenFileErr, err.Error())
		return
	}

	load(d)
	base.Output(c, 0, fillPara())
	backupClear()
	return
}

func getBackup() {
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".json") {
			continue
		} else {
			f := file.Name()
			b, err := getFileInfo(f)
			if err == nil {
				backup = append(backup, b)
			} else {
				log.GetLogger().Log(log.Wrong, "getBackup err", err.Error(), f)
			}
		}
	}
}

func getFileInfo(file string) (*backupInfo, error) {
	data, err := loadFile(file)
	if err != nil {
		return nil, err
	}

	return &backupInfo{FileName:file[:len(file)-5], Name:data.Name, OverDay:data.GameTime.Overday,
		StartTime:data.StartTime, LastTime:data.LastTime}, nil
}

func loadFile(file string) (*base.Data, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	var data base.Data
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func load(d *base.Data) {
	//加载游戏时间
	env.LoadTime(d.GameTime)
	//加载人物
	human.LoadHuman(d.User, d.Name)
	//加载拥有建筑
	//home.LoadOwnBuilding(d.OwnBuild)
	//加载拥有物品
	//materiel.LoadOwnThings(d.OwnProduct)
	//更新各地点上次到访时间
	//plat.LoadPublic(d.PlatLastTime)
	//更新各地点物品数量
	//plat.LoadProduct(d.PlatProduct)
	//更新各地点动物数量
	//plat.LoadAnimal(d.PlatAnimal)

	data.StartTime = d.StartTime
	//上次游戏时间更新
	data.LastTime = time.Now().Format("2006-01-02 15:04:05")
	data.Name = d.Name
}
