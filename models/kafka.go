package models

import (
	"encoding/json"
	"fmt"
	"omonitor/utils"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/samuel/go-zookeeper/zk"
)

func ZookeeperClient(hosts []string) (conn *zk.Conn, err error) {
	zkConn, _, err := zk.Connect(hosts, 3*time.Second)
	if err != nil {
		return nil, err
	}
	return zkConn, nil
}

func GetChildren(hosts []string, zkPath string) ([]string, error) {
	zkClient, err := ZookeeperClient(hosts)
	defer zkClient.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	childrens, _, err := zkClient.Children(zkPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return childrens, nil
}

func GetKafkaServers(zkClient *zk.Conn) ([]string, error) {
	childrens, _, err := zkClient.Children("/brokers/ids")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	servers := make([]string, len(childrens))
	for i, v := range childrens {
		data, _, err := zkClient.Get("/brokers/ids/" + v)
		if err != nil {
			continue
		}
		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			fmt.Println("json error:", err)
		}
		servers[i] = strings.Join([]string{m["host"].(string), strconv.FormatFloat(m["port"].(float64), 'g', -1, 32)}, ":")
	}
	return servers, nil
}

func GetConsumerGroups(hosts []string, keyword string) ([]string, error) {
	consumerGroups, err := GetChildren(hosts, "/consumers")
	if err != nil {
		return nil, err
	}
	if keyword != "" {
		var newConsumerGroups []string
		for _, groupName := range consumerGroups {
			if strings.Contains(groupName, keyword) {
				newConsumerGroups = append(newConsumerGroups, groupName)
			}
		}
		return newConsumerGroups, nil
	}
	return consumerGroups, nil
}

func GetTopicList(hosts []string, keyword string) ([]string, error) {
	topicList, err := GetChildren(hosts, "/brokers/topics")
	if err != nil {
		return nil, err
	}
	if keyword != "" {
		var newTopicList []string
		for _, topic := range topicList {
			if strings.Contains(topic, keyword) {
				newTopicList = append(newTopicList, topic)
			}
		}
		return newTopicList, nil
	}
	return topicList, nil
}

type KafkaConsumerInfo struct {
	Env          string
	Topic        string
	GroupName    string
	PartitionNum int
	Offset       int64
	LogSize      int64
	Lag          int64
}

func GetConsumerGroupTopics(hosts []string, env, consumerGroup string) ([]KafkaConsumerInfo, error) {
	zkClient, err := ZookeeperClient(hosts)
	defer zkClient.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var info []KafkaConsumerInfo
	var data KafkaConsumerInfo
	topics, _, err := zkClient.Children("/consumers/" + consumerGroup + "/offsets")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, topic := range topics {
		data.Env = env
		data.Topic = topic
		data.GroupName = consumerGroup
		partitions, offsetCount, err := GetConsumerTopicPartitions(zkClient, consumerGroup, topic)
		if err != nil {
			fmt.Println(err)
			continue
		}
		logSize, err := GetKafkaOffset(zkClient, topic, partitions)
		if err != nil {
			fmt.Println("kafka error:", err)
			continue
		}
		data.PartitionNum = len(partitions)
		data.Offset = offsetCount
		data.LogSize = logSize
		data.Lag = logSize - offsetCount
		info = append(info, data)
	}
	return info, nil
}

func GetConsumerTopicPartitions(zkClient *zk.Conn, consumerGroup, topic string) ([]string, int64, error) {
	partitions, _, err := zkClient.Children("/consumers/" + consumerGroup + "/offsets/" + topic)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	var offsetCount int64
	for _, partition := range partitions {
		offsetStr, _, err := zkClient.Get("/consumers/" + consumerGroup + "/offsets/" + topic + "/" + partition)
		if err != nil {
			fmt.Println(err)
			continue
		}
		offset, _ := strconv.ParseInt(string(offsetStr), 10, 64)
		offsetCount += offset
	}
	return partitions, offsetCount, nil
}

func GetKafkaOffset(zkClient *zk.Conn, topic string, partitions []string) (int64, error) {
	servers, err := GetKafkaServers(zkClient)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	client, err := sarama.NewClient(servers, nil)
	defer client.Close()
	if err != nil {
		fmt.Println("can't connect to broker")
		return 0, err
	}
	var kafkaOffsetCount int64
	for _, partition := range partitions {
		partitionInt, _ := strconv.Atoi(partition)
		kafkaOffset, err := client.GetOffset(topic, int32(partitionInt), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("kafka offset error:", err)
		}
		kafkaOffsetCount += kafkaOffset
	}
	return kafkaOffsetCount, nil
}

type TopicPartitionsInfo struct {
	Partition string
	Offset    int64
	LogSize   int64
	Lag       int64
	Owner     string
}

func GetTopicPartitions(hosts []string, consumerGroup, topic string) ([]TopicPartitionsInfo, error) {
	zkClient, err := ZookeeperClient(hosts)
	defer zkClient.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	partitions, _, err := zkClient.Children("/consumers/" + consumerGroup + "/offsets/" + topic)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var info []TopicPartitionsInfo
	var data TopicPartitionsInfo
	servers, err := GetKafkaServers(zkClient)
	if err != nil {
		fmt.Println("GetTopicPartitions kafka error:", err)
		return nil, err
	}
	client, err := sarama.NewClient(servers, nil)
	defer client.Close()
	if err != nil {
		fmt.Println("can't connect to broker")
		return nil, err
	}
	for _, partition := range partitions {
		data.Partition = partition
		offsetStr, _, err := zkClient.Get("/consumers/" + consumerGroup + "/offsets/" + topic + "/" + partition)
		if err != nil {
			fmt.Println(err)
			continue
		}
		owner, _, err := zkClient.Get("/consumers/" + consumerGroup + "/owners/" + topic + "/" + partition)
		if err != nil {
			fmt.Println(err)
			continue
		}
		partitionInt, _ := strconv.Atoi(partition)
		kafkaOffset, err := client.GetOffset(topic, int32(partitionInt), sarama.OffsetNewest)
		offset, _ := strconv.ParseInt(string(offsetStr), 10, 64)
		data.Offset = offset
		data.LogSize = kafkaOffset
		data.Lag = kafkaOffset - offset
		data.Owner = string(owner)
		info = append(info, data)
	}
	return info, nil
}

// kafka monitor

type Kafka_consumer_groups struct {
	Id         int
	Groupname  string
	Monitoring int
	Alarmval   int
	Alarmgroup int
	Alerts     int
	Comment    string
}

type Kafka_consumer_group_topics struct {
	Id         int
	Topicname  string
	Groupid    int
	Monitoring int
	Alarmval   int
	Already    int
	Comment    string
}

func init() {
	orm.RegisterModel(new(Kafka_consumer_groups))
	orm.RegisterModel(new(Kafka_consumer_group_topics))
}

func AddKafkaConsumerGroup(groupName string) {
	o := orm.NewOrm()
	group := new(Kafka_consumer_groups)
	group.Groupname = groupName
	group.Alarmgroup = 1
	group.Alerts = 1
	_, err := o.Insert(group)
	if err != nil {
		fmt.Println(err)
	}
}

func GetKafkaConsumerGroup(groupName string) (result Kafka_consumer_groups, err error) {
	var group Kafka_consumer_groups
	err = orm.NewOrm().QueryTable("kafka_consumer_groups").Filter("Groupname", groupName).One(&group)
	if err != nil {
		return group, err
	}
	return group, nil
}

func GetKafkaConsumerGroupById(consumerId int) (result Kafka_consumer_groups, err error) {
	var group Kafka_consumer_groups
	err = orm.NewOrm().QueryTable("kafka_consumer_groups").Filter("Id", consumerId).One(&group)
	if err != nil {
		return group, err
	}
	return group, nil
}

func CountAlarmConsumerGroup() int64 {
	var groups []Kafka_consumer_groups
	num, _ := orm.NewOrm().QueryTable("kafka_consumer_groups").All(&groups)
	return num
}

type NewConsumerGroups struct {
	Kafka_consumer_groups
	AlarmGroupStr string
}

func GetAlarmConsumerGroupList(page int, offset int, keyword string) (group []NewConsumerGroups, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("kafka_consumer_groups")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("Groupname__icontains", keyword)).Or("Monitoring__icontains", keyword)
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()
	var groups []Kafka_consumer_groups
	_, err = qs.Limit(offset, start).All(&groups)
	var newGroups []NewConsumerGroups
	var newag NewConsumerGroups
	for _, a := range groups {
		alarm_group, _ := GetAlarmGroupById(a.Alarmgroup)
		newag.Id = a.Id
		newag.Groupname = a.Groupname
		newag.Monitoring = a.Monitoring
		newag.Alarmval = a.Alarmval
		newag.Alarmgroup = a.Alarmgroup
		newag.Alerts = a.Alerts
		newag.AlarmGroupStr = alarm_group.Group
		newag.Comment = a.Comment
		newGroups = append(newGroups, newag)
	}
	return newGroups, err
}

func UpdateAlarmConsumerGroup(consumerId int, cgroup Kafka_consumer_groups) error {
	var consumer Kafka_consumer_groups
	o := orm.NewOrm()
	consumer = Kafka_consumer_groups{Id: consumerId}
	consumer.Groupname = cgroup.Groupname
	consumer.Monitoring = cgroup.Monitoring
	consumer.Alarmval = cgroup.Alarmval
	consumer.Alarmgroup = cgroup.Alarmgroup
	consumer.Alerts = cgroup.Alerts
	consumer.Comment = cgroup.Comment
	_, err := o.Update(&consumer)
	return err
}

func DelAlarmConsumerGroup(consumerId string) error {
	consumerIdInt, _ := strconv.Atoi(consumerId)
	o := orm.NewOrm()
	_, err := o.Delete(&Kafka_consumer_groups{Id: consumerIdInt})
	return err
}

func AddKafkaConsumerGroupTopic(topicName string, groupId int) {
	o := orm.NewOrm()
	topic := new(Kafka_consumer_group_topics)
	topic.Topicname = topicName
	topic.Groupid = groupId
	_, err := o.Insert(topic)
	if err != nil {
		fmt.Println(err)
	}
}

func GetKafkaConsumerGroupTopic(topicName string, groupId int) (result Kafka_consumer_group_topics, err error) {
	var topic Kafka_consumer_group_topics
	err = orm.NewOrm().QueryTable("kafka_consumer_group_topics").Filter("Topicname", topicName).Filter("Groupid", groupId).One(&topic)
	if err != nil {
		return topic, err
	}
	return topic, nil
}

func GetAlarmConsumerTopicById(topicId int) (result Kafka_consumer_group_topics, err error) {
	var topic Kafka_consumer_group_topics
	err = orm.NewOrm().QueryTable("kafka_consumer_group_topics").Filter("Id", topicId).One(&topic)
	return topic, err
}

func CountAlarmConsumerTopics(groupId int) int64 {
	var topics []Kafka_consumer_group_topics
	num, _ := orm.NewOrm().QueryTable("kafka_consumer_group_topics").Filter("groupid", groupId).All(&topics)
	return num
}

func GetAlarmConsumerTopicList(groupId int, page int, offset int, keyword string) (topic []Kafka_consumer_group_topics, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("kafka_consumer_group_topics")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("topicname__icontains", keyword))
	}
	cond = cond.And("groupid", groupId)
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()
	var topics []Kafka_consumer_group_topics
	_, err = qs.Limit(offset, start).All(&topics)
	return topics, err
}

func UpdateConsumerTopic(topicId int, newtopic Kafka_consumer_group_topics) error {
	var topic Kafka_consumer_group_topics
	o := orm.NewOrm()
	topic = Kafka_consumer_group_topics{Id: topicId}
	topic.Topicname = newtopic.Topicname
	topic.Monitoring = newtopic.Monitoring
	topic.Alarmval = newtopic.Alarmval
	topic.Already = newtopic.Already
	topic.Comment = newtopic.Comment
	_, err := o.Update(&topic, "Topicname", "Monitoring", "Alarmval", "Already", "Comment")
	return err
}

func UpdateConsumerTopicAlready(topicId, already int) {
	var topic Kafka_consumer_group_topics
	o := orm.NewOrm()
	topic = Kafka_consumer_group_topics{Id: topicId}
	topic.Already = already
	_, err := o.Update(&topic, "Already")
	if err != nil {
		fmt.Println(err)
	}
}

func DelConsumerTopic(topicId string) error {
	topicIdInt, _ := strconv.Atoi(topicId)
	o := orm.NewOrm()
	_, err := o.Delete(&Kafka_consumer_group_topics{Id: topicIdInt})
	return err
}

func MonitorConsumerGroups() {
	zkHosts := beego.AppConfig.String("kafka_zk_1")
	zkHostList := strings.Split(zkHosts, ",")
	zkClient, _, err := zk.Connect(zkHostList, 3*time.Second)
	defer zkClient.Close()
	if err != nil {
		fmt.Println(err)
	}
	consumerGroups, _, err := zkClient.Children("/consumers")
	if err != nil {
		fmt.Println(err)
	}
	if len(consumerGroups) > 0 {
		for _, groupName := range consumerGroups {
			group, err := GetKafkaConsumerGroup(groupName)
			if err != nil {
				go AddKafkaConsumerGroup(groupName)
			} else {
				topics, _, err := zkClient.Children("/consumers/" + groupName + "/offsets")
				if err != nil {
					fmt.Println(err)
					continue
				}
				for _, topic := range topics {
					isTopic, err := GetKafkaConsumerGroupTopic(topic, group.Id)
					if err != nil {
						go AddKafkaConsumerGroupTopic(topic, group.Id)
					}
					if group.Monitoring > 0 {
						if isTopic.Monitoring > 0 {
							partitions, offsetCount, err := GetConsumerTopicPartitions(zkClient, groupName, topic)
							if err != nil {
								continue
							}
							logSize, err := GetKafkaOffset(zkClient, topic, partitions)
							if err != nil {
								continue
							}
							Lag := logSize - offsetCount
							go KafkaAlarm(group, isTopic, Lag)
						}
					}
				}
			}
		}
	}
}

func KafkaAlarm(consumerGroup Kafka_consumer_groups, topic Kafka_consumer_group_topics, lag int64) {
	alarmVal := int64(topic.Alarmval)
	if topic.Alarmval == 0 {
		alarmVal = int64(consumerGroup.Alarmval)
	}
	if alarmVal > 0 {
		groupAlarm, err := GetAlarmGroupById(consumerGroup.Alarmgroup)
		if err != nil {
			fmt.Println(err)
		}
		alarmUserIds := GetAlarmGroupIds(groupAlarm.Id)
		alarmUserList, err := GetUserListByIds(alarmUserIds)
		if err != nil {
			fmt.Println(err)
		} else {
			var mobileList []int64
			var emailList []string
			for _, user := range alarmUserList {
				mobileList = append(mobileList, user.Mobile)
				emailList = append(emailList, user.Email)
			}
			if lag > alarmVal {
				if topic.Already < consumerGroup.Alerts {
					if groupAlarm.Sms > 0 {
						nowTime := strings.Split(time.Now().String(), ".")[0]
						msg := fmt.Sprintf("\n消费组：%s\nTopic名称：%s\n当前未消费：%d\n告警阀值：%d\n时间：%s", consumerGroup.Groupname, topic.Topicname, lag, alarmVal, nowTime)
						utils.SendSms(mobileList, msg)
					}
					if groupAlarm.Email > 0 {
						msg := fmt.Sprintf(`<html><body>
						<h3>消费组： %s</h3>
						<h3>Topic名称： %s</h3>
						<h3>告警阀值： %d</h3>
						<h3>当前值： %d</h3>
						<h3>时间： %s</h3>
						<h3>当前状态： 故障</h3>
						</body></html>`, consumerGroup.Groupname, topic.Topicname, alarmVal, lag, time.Now().String())
						utils.SendEmail(emailList, msg)
					}
					UpdateConsumerTopicAlready(topic.Id, topic.Already+1)
				}
			} else {
				if topic.Already > 0 {
					if groupAlarm.Sms > 0 {
						nowTime := strings.Split(time.Now().String(), ".")[0]
						msg := fmt.Sprintf("\n消费组：%s\nTopic名称：%s\n当前未消费：%d\n状态：已恢复\n时间：%s", consumerGroup.Groupname, topic.Topicname, lag, nowTime)
						utils.SendSms(mobileList, msg)
					}
					if groupAlarm.Email > 0 {
						msg := fmt.Sprintf(`<html><body>
						<h3>消费组： %s</h3>
						<h3>Topic名称： %s</h3>
						<h3>告警阀值： %d</h3>
						<h3>当前值： %d</h3>
						<h3>时间： %s</h3>
						<h3>当前状态： 已恢复</h3>
						</body></html>`, consumerGroup.Groupname, topic.Topicname, alarmVal, lag, time.Now().String())
						utils.SendEmail(emailList, msg)
					}
					UpdateConsumerTopicAlready(topic.Id, 0)
				}
			}
		}
	}
}
