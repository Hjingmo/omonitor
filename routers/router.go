package routers

import (
	"omonitor/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/skin_config", &controllers.SkinConfigController{})
	beego.Router("/login", &controllers.LoginUserController{})
	beego.Router("/logout", &controllers.LogoutUserController{})
	beego.Router("/permission", &controllers.PermissionController{})
	// user
	beego.Router("/user/manage", &controllers.UserListController{})
	beego.Router("/user/adduser", &controllers.UserAddController{})
	beego.Router("/user/edituser", &controllers.UserEditController{})
	beego.Router("/user/deluser", &controllers.UserDelController{})
	beego.Router("/user/permission", &controllers.PermissionListController{})
	beego.Router("/user/addpermission", &controllers.PermissionAddController{})
	beego.Router("/user/editpermission", &controllers.PermissionEditController{})
	beego.Router("/user/delpermission", &controllers.PermissionDelController{})
	// alarm
	beego.Router("/alarm/alarmgroup", &controllers.AlarmGroups{})
	beego.Router("/alarm/addgroup", &controllers.AlarmGroupAdd{})
	beego.Router("/alarm/editgroup", &controllers.AlarmGroupEdit{})
	beego.Router("/alarm/delgroup", &controllers.AlarmGroupDel{})
	beego.Router("/alarm/groupusers", &controllers.AlarmGroupUsers{})
	beego.Router("/alarm/kafkaset", &controllers.AlarmConsumerGroups{})
	beego.Router("/alarm/consumergroupedit", &controllers.AlarmConsumerGroupEdit{})
	beego.Router("/alarm/consumergroupdel", &controllers.AlarmConsumerGroupDel{})
	beego.Router("/alarm/consumertopics", &controllers.AlarmConsumerGroupTopics{})
	beego.Router("/alarm/consumertopicedit", &controllers.ConsumerGroupTopicEdit{})
	beego.Router("/alarm/consumertopicdel", &controllers.ConsumerGroupTopicDel{})
	// kafka
	beego.Router("/kafka/consumers", &controllers.ConsumerGroups{})
	beego.Router("/kafka/consumertopics", &controllers.ConsumerGroupTopics{})
	beego.Router("/kafka/topics", &controllers.TopicList{})
	beego.Router("/kafka/topicpartitions", &controllers.TopicPartitions{})
	beego.Router("/kafka/servers", &controllers.KafkaServers{})
}
