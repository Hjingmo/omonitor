package tasks

import (
	"omonitor/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	TaskInt()
}

func TaskInt() {
	task, err := beego.AppConfig.Bool("task")
	if err != nil {
		task = false
	}
	tk1 := toolbox.NewTask("taska", "18 19 */1 * * *", func() error { models.MonitorConsumerGroups(); return nil })
	//	tk1.Run()
	if task {
		toolbox.AddTask("taska", tk1)
		toolbox.StartTask()
	}
}
