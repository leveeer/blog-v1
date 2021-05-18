package jobs

import (
	"blog-go-gin/common"
	"github.com/robfig/cron"
	"log"
	"time"
)

func ClearIpSet() {
	members := common.RedisUtil.SMembers(common.IpSet)
	var slice []interface{}
	for _, member := range members {
		slice = append(slice, member)
	}
	common.RedisUtil.SRems(common.IpSet, slice...)
}

func RegisterCron() {
	go func() {
		c := cron.New()
		log.Println("cron register success ...")
		spec := "0 0 0 1/1 * ?"
		_ = c.AddFunc(spec, ClearIpSet)
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
