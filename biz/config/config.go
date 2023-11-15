package config

import (
	"runtime"
	"scnu_acm_rank/biz/model"
)

var Conf model.Config
var Update chan struct{}
var updateList []updateConf

func init() {
	Conf = model.Config{}
	Update = make(chan struct{}, 0)
	updateList = make([]updateConf, 0, 20)
	runtime.KeepAlive(Conf)
	runtime.KeepAlive(Update)
	runtime.KeepAlive(updateList)
	go update()
}

func Add(a updateConf) {
	updateList = append(updateList, a)
}

func update() {
	for {
		<-Update
		model.DB.Model(&model.Config{}).First(&Conf)
		for _, v := range updateList {
			v.Update()
		}
	}
}
