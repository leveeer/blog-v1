package crons

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"github.com/robfig/cron"
	"time"
)

func SaveAndClearIpSet() {
	dayViewCount, err := common.GetRedisUtil().SCard(common.IpSet)
	if err != nil {
		logging.Logger.Debug(err)
	}
	err = model.AddUniqueView(&model.UniqueView{
		CreateTime: time.Now().Unix(),
		ViewsCount: int(dayViewCount),
	})
	if err != nil {
		logging.Logger.Debug(err)
	}
	members, err := common.GetRedisUtil().SMembers(common.IpSet)
	if err != nil {
		logging.Logger.Debug(err)
	}
	var slice []interface{}
	for _, member := range members {
		slice = append(slice, member)
	}
	common.GetRedisUtil().SRems(common.IpSet, slice...)
}

func RegisterCron() {
	go func() {
		c := cron.New()
		logging.Logger.Debug("cron register success...")
		spec := "0 0 0 1/1 * ?"
		_ = c.AddFunc(spec, SaveAndClearIpSet)
		c.Start()
		t := time.NewTimer(time.Second * 10)
		for {
			select {
			case <-t.C:
				t.Reset(time.Second * 10)
			}
		}
	}()
}
