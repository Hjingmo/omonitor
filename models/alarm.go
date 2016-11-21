package models

import (
	"fmt"
	"strconv"

	"github.com/yangji168/omonitor/utils"

	"github.com/astaxie/beego/orm"
)

type Alarm_groups struct {
	Id      int
	Group   string `orm:"size(64)"`
	Sms     int
	Email   int
	Comment string `orm:"size(128)"`
}

type Alarm_group_users struct {
	Id      int
	Groupid int
	Userid  int64
}

func init() {
	orm.RegisterModel(new(Alarm_groups))
	orm.RegisterModel(new(Alarm_group_users))
}

func CountAlarmGroups() int64 {
	var groups []Alarm_groups
	num, _ := orm.NewOrm().QueryTable("alarm_groups").All(&groups)
	return num
}

func GetAllAlarmGroups() (group []Alarm_groups, err error) {
	var groups []Alarm_groups
	_, err = orm.NewOrm().QueryTable("alarm_groups").All(&groups)
	return groups, err
}

func GetAlarmGroups(page int, offset int, keyword string) (group []Alarm_groups, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("alarm_groups")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("group__icontains", keyword))
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()
	var groups []Alarm_groups
	_, err = qs.Limit(offset, start).All(&groups)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return groups, nil
}

func GetAlarmGroupByName(groupName string) bool {
	result := false
	var group Alarm_groups
	err := orm.NewOrm().QueryTable("alarm_groups").Filter("group", groupName).One(&group)
	if err != nil {
		result = true
	}
	return result
}

func GetAlarmGroupById(groupId int) (g Alarm_groups, err error) {
	var group Alarm_groups
	err = orm.NewOrm().QueryTable("alarm_groups").Filter("Id", groupId).One(&group)
	return group, err
}

func AddAlarmGroup(groupName, comment string, sms, email int) error {
	o := orm.NewOrm()
	group := new(Alarm_groups)
	group.Group = groupName
	group.Sms = sms
	group.Email = email
	group.Comment = comment
	_, err := o.Insert(group)
	return err
}

func UpdateAlarmGroup(groupId, sms, email int, groupName, comment string, selectIds []string) error {
	group_users := GetAlarmGroupIds(groupId)
	selectIdInts := utils.StringToInt64(selectIds)
	adds, dels := utils.UpdateCompareInt64(group_users, selectIdInts)
	if len(adds) > 0 {
		for _, add := range adds {
			aerr := AddAlarmGroupUser(groupId, add)
			if aerr != nil {
				fmt.Println(aerr)
			}
		}
	}
	if len(dels) > 0 {
		for _, del := range dels {
			derr := DelAlarmGroupUser(groupId, del)
			if derr != nil {
				fmt.Println(derr)
			}
		}
	}
	var alarmGroup Alarm_groups
	o := orm.NewOrm()
	alarmGroup = Alarm_groups{Id: groupId}
	alarmGroup.Group = groupName
	alarmGroup.Sms = sms
	alarmGroup.Email = email
	alarmGroup.Comment = comment
	_, err := o.Update(&alarmGroup)
	return err
}

func DelAlarmGroup(groupId string) error {
	idInt, _ := strconv.Atoi(groupId)
	o := orm.NewOrm()
	_, err := o.Delete(&Alarm_groups{Id: idInt})
	return err
}

func AddAlarmGroupUser(groupId int, userId int64) error {
	o := orm.NewOrm()
	groupUser := new(Alarm_group_users)
	groupUser.Groupid = groupId
	groupUser.Userid = userId
	_, err := o.Insert(groupUser)
	return err
}

func DelAlarmGroupUser(groupId int, userId int64) error {
	var groupUser Alarm_group_users
	err := orm.NewOrm().QueryTable("alarm_group_users").Filter("groupId", groupId).Filter("Userid", userId).One(&groupUser)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&Alarm_group_users{Id: groupUser.Id})
	return err
}

func GetAlarmGroupIds(groupId int) []int64 {
	var ids []int64
	o := orm.NewOrm()
	qs := o.QueryTable("alarm_group_users")
	var groupUserIds []Alarm_group_users
	_, err := qs.Filter("Groupid", groupId).All(&groupUserIds)
	if err != nil {
		fmt.Println(err)
		return ids
	}
	for _, uid := range groupUserIds {
		ids = append(ids, uid.Userid)
	}
	return ids
}

func GetALarmUserNoSelect(userIds []int64) (users []Users) {
	groupUserIds, err := GetUserListByIds(userIds)
	if err != nil {
		fmt.Println(err)
	}
	user_ids := &groupUserIds
	allUserList, err := GetAllUsers()
	var result []Users
	var user Users
	for _, u := range allUserList {
		chk := CheckUserInUsers(*user_ids, u.Id)
		if chk {
			user.Id = u.Id
			user.Username = u.Username
			user.Firstname = u.Firstname
			user.Lastname = u.Lastname
			user.Mobile = u.Mobile
			user.Email = u.Email
			result = append(result, user)
		}
	}
	return result
}

func CheckUserInUsers(userlist []Users, userId int64) bool {
	result := true
	for _, u := range userlist {
		if userId == u.Id {
			result = false
		}
	}
	return result
}

func GetAlarmGroupNoSelect(groupId int) (u []Users) {
	group_users := GetAlarmGroupIds(groupId)
	userList := GetALarmUserNoSelect(group_users)
	return userList
}

func GetAlarmGroupSelect(groupId int) (u []Users) {
	group_users := GetAlarmGroupIds(groupId)
	groupUserIds, err := GetUserListByIds(group_users)
	if err != nil {
		fmt.Println(err)
	}
	return groupUserIds
}
