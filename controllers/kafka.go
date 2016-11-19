package controllers

import (
	"fmt"
	"omonitor/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type ConsumerGroups struct {
	BaseController
}

func (c *ConsumerGroups) Get() {
	keyword := c.GetString("keyword")
	env, err := c.GetInt("env")
	if err != nil {
		c.Abort("401")
	}
	c.Layout = "inc/base.tpl"
	if err != nil {
		c.Abort("401")
	}
	c.TplName = "kafka/consumerGroups.tpl"
	envStr := strconv.Itoa(env)
	zkHosts := beego.AppConfig.String("kafka_zk_" + envStr)
	zkHostList := strings.Split(zkHosts, ",")
	if len(zkHostList) == 0 {
		c.Abort("401")
	}
	consumerGroups, err := models.GetConsumerGroups(zkHostList, keyword)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	var data []map[string]string
	for _, group := range consumerGroups {
		m := make(map[string]string)
		m["group"] = group
		m["env"] = envStr
		data = append(data, m)
	}
	c.Data["consumerGroups"] = data
	c.Data["env"] = envStr
}

type TopicList struct {
	BaseController
}

func (c *TopicList) Get() {
	keyword := c.GetString("keyword")
	env, err := c.GetInt("env")
	if err != nil {
		c.Abort("401")
	}
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/topicList.tpl"
	envStr := strconv.Itoa(env)
	zkHosts := beego.AppConfig.String("kafka_zk_" + envStr)
	zkHostList := strings.Split(zkHosts, ",")
	if len(zkHostList) == 0 {
		c.Abort("401")
	}
	topicList, err := models.GetTopicList(zkHostList, keyword)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["topicList"] = topicList
	c.Data["env"] = envStr
}

type KafkaServers struct {
	BaseController
}

func (c *KafkaServers) Get() {
	env, err := c.GetInt("env")
	if err != nil {
		c.Abort("401")
	}
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/kafkaServers.tpl"
	envStr := strconv.Itoa(env)
	zkHosts := beego.AppConfig.String("kafka_zk_" + envStr)
	zkHostList := strings.Split(zkHosts, ",")
	if len(zkHostList) == 0 {
		c.Abort("401")
	}
	zkClient, err := models.ZookeeperClient(zkHostList)
	defer zkClient.Close()
	if err != nil {
		fmt.Println(err)
	}
	servers, err := models.GetKafkaServers(zkClient)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["servers"] = servers
	c.Data["env"] = envStr
}

type ConsumerGroupTopics struct {
	BaseController
}

func (c *ConsumerGroupTopics) Get() {
	group := c.GetString("group")
	env, err := c.GetInt("env")
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	envStr := strconv.Itoa(env)
	zkHosts := beego.AppConfig.String("kafka_zk_" + envStr)
	zkHostList := strings.Split(zkHosts, ",")
	if len(zkHostList) == 0 {
		fmt.Println("Error: zkHostList is Null")
		c.Abort("401")
	}
	topicOffsets, err := models.GetConsumerGroupTopics(zkHostList, envStr, group)
	if err != nil {
		c.Abort("401")
	}
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/consumerGroupTopics.tpl"
	c.Data["env"] = envStr
	c.Data["group"] = group
	c.Data["topicOffsets"] = topicOffsets
}

type TopicPartitions struct {
	BaseController
}

func (c *TopicPartitions) Get() {
	group := c.GetString("group")
	topic := c.GetString("topic")
	env, err := c.GetInt("env")
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	envStr := strconv.Itoa(env)
	zkHosts := beego.AppConfig.String("kafka_zk_" + envStr)
	zkHostList := strings.Split(zkHosts, ",")
	if len(zkHostList) == 0 {
		fmt.Println("Error: zkHostList is Null")
		c.Abort("401")
	}
	topicPartition, err := models.GetTopicPartitions(zkHostList, group, topic)
	if err != nil {
		c.Abort("401")
	}
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/topicPartitions.tpl"
	c.Data["env"] = envStr
	c.Data["topic"] = topic
	c.Data["partitions"] = topicPartition
}

// Monitor

type AlarmConsumerGroups struct {
	BaseController
}

func (c *AlarmConsumerGroups) Get() {
	keyword := c.GetString("keyword")
	page, err := c.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	countGroups := models.CountAlarmConsumerGroup()
	pagination.SetPaginator(c.Ctx, offset, countGroups)
	groups, err := models.GetAlarmConsumerGroupList(page, offset, keyword)
	if err != nil {
		c.Abort("401")
	}
	c.Data["groups"] = groups
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/alarmConsumerGroups.tpl"
}

type AlarmConsumerGroupEdit struct {
	BaseController
}

func (c *AlarmConsumerGroupEdit) Get() {
	consumerId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	consumerGroup, err := models.GetKafkaConsumerGroupById(consumerId)
	if err != nil {
		c.Abort("401")
	}
	alarmGroups, err := models.GetAllAlarmGroups()
	if err != nil {
		c.Abort("401")
	}
	c.Data["alarmGroups"] = alarmGroups
	c.Data["consumer"] = consumerGroup
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/alarmConsumerGroupEdit.tpl"
}

func (c *AlarmConsumerGroupEdit) Post() {
	consumerId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	consumerGroup := c.GetString("consumername")
	monitor, err := c.GetInt("monitoring")
	if err != nil {
		c.Abort("401")
	}
	alarmval, err := c.GetInt("alarmval")
	if err != nil {
		c.Abort("401")
	}
	alarmgroup, err := c.GetInt("alarmgroup")
	if err != nil {
		alarmgroup = 1
	}
	alerts, err := c.GetInt("alerts")
	if err != nil {
		alerts = 1
	}
	comment := c.GetString("comment")
	var group models.Kafka_consumer_groups
	group.Groupname = consumerGroup
	group.Monitoring = monitor
	group.Alarmval = alarmval
	group.Alarmgroup = alarmgroup
	group.Alerts = alerts
	group.Comment = comment
	err = models.UpdateAlarmConsumerGroup(consumerId, group)
	if err != nil {
		c.Abort("401")
	}
	c.Ctx.Redirect(302, "/alarm/kafkaset")
}

type AlarmConsumerGroupDel struct {
	BaseController
}

func (c *AlarmConsumerGroupDel) Get() {
	consumerId := c.GetString("id")
	consumerIdList := strings.Split(consumerId, ",")
	for _, cid := range consumerIdList {
		err := models.DelAlarmConsumerGroup(cid)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	c.Ctx.Redirect(302, "/alarm/kafkaset")
}

type AlarmConsumerGroupTopics struct {
	BaseController
}

func (c *AlarmConsumerGroupTopics) Get() {
	keyword := c.GetString("keyword")
	page, err := c.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	groupId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	countTopic := models.CountAlarmConsumerTopics(groupId)
	pagination.SetPaginator(c.Ctx, offset, countTopic)
	topics, err := models.GetAlarmConsumerTopicList(groupId, page, offset, keyword)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["topics"] = topics
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/alarmConsumerGroupTopics.tpl"
}

type ConsumerGroupTopicEdit struct {
	BaseController
}

func (c *ConsumerGroupTopicEdit) Get() {
	topicId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	consumerTopic, err := models.GetAlarmConsumerTopicById(topicId)
	if err != nil {
		c.Abort("401")
	}
	c.Data["topic"] = consumerTopic
	c.Layout = "inc/base.tpl"
	c.TplName = "kafka/consumerGroupTopicEdit.tpl"
}

func (c *ConsumerGroupTopicEdit) Post() {
	topicId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	topicName := c.GetString("topicname")
	monitor, err := c.GetInt("monitoring")
	if err != nil {
		c.Abort("401")
	}
	alarmval, err := c.GetInt("alarmval")
	if err != nil {
		c.Abort("401")
	}
	already, err := c.GetInt("already")
	if err != nil {
		already = 0
	}
	comment := c.GetString("comment")
	groupId, err := c.GetInt("groupid")
	if err != nil {
		groupId = 1
	}
	var topic models.Kafka_consumer_group_topics
	topic.Topicname = topicName
	topic.Monitoring = monitor
	topic.Alarmval = alarmval
	topic.Already = already
	topic.Comment = comment
	err = models.UpdateConsumerTopic(topicId, topic)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Ctx.Redirect(302, fmt.Sprintf("/alarm/consumertopics?id=%d", groupId))
}

type ConsumerGroupTopicDel struct {
	BaseController
}

func (c *ConsumerGroupTopicDel) Get() {
	groupId, err := c.GetInt("groupid")
	if err != nil {
		groupId = 1
	}
	topicId := c.GetString("id")
	topicIdList := strings.Split(topicId, ",")
	for _, tid := range topicIdList {
		err := models.DelConsumerTopic(tid)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	c.Ctx.Redirect(302, fmt.Sprintf("/alarm/consumertopics?id=%d", groupId))
}
