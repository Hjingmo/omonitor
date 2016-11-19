package controllers

import (
	"fmt"
	"omonitor/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type AlarmGroups struct {
	BaseController
}

func (c *AlarmGroups) Get() {
	keyword := c.GetString("keyword")
	page, err := c.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	countGroups := models.CountAlarmGroups()
	pagination.SetPaginator(c.Ctx, offset, countGroups)
	groups, err := models.GetAlarmGroups(page, offset, keyword)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["groups"] = groups
	c.Layout = "inc/base.tpl"
	c.TplName = "alarm/alarmGroups.tpl"
}

type AlarmGroupAdd struct {
	BaseController
}

func (c *AlarmGroupAdd) Get() {
	c.Layout = "inc/base.tpl"
	c.TplName = "alarm/alarmGroupAdd.tpl"
}

func (c *AlarmGroupAdd) Post() {
	c.Layout = "inc/base.tpl"
	c.TplName = "alarm/alarmGroupAdd.tpl"
	groupName := c.GetString("groupname")
	sms, err := c.GetInt("smsOptions")
	if err != nil {
		c.Abort("401")
	}
	email, err := c.GetInt("emailOptions")
	if err != nil {
		c.Abort("401")
	}
	comment := c.GetString("comment")
	groupTest := models.GetAlarmGroupByName(groupName)
	if groupTest {
		err := models.AddAlarmGroup(groupName, comment, sms, email)
		if err != nil {
			c.Data["emg"] = groupName + "添加失败"
		} else {
			c.Data["smg"] = groupName + "添加成功"
		}
	} else {
		c.Data["emg"] = groupName + "组已存在"
	}
}

type AlarmGroupEdit struct {
	BaseController
}

func (c *AlarmGroupEdit) Get() {
	c.Layout = "inc/base.tpl"
	c.TplName = "alarm/alarmGroupEdit.tpl"
	groupId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	groupNoSelect := models.GetAlarmGroupNoSelect(groupId)
	groupSelect := models.GetAlarmGroupSelect(groupId)
	group, err := models.GetAlarmGroupById(groupId)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["groupNoSelect"] = groupNoSelect
	c.Data["groupSelect"] = groupSelect
	c.Data["group"] = group
}

func (c *AlarmGroupEdit) Post() {
	groupId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	groupSelect := c.GetStrings("project_select")
	groupName := c.GetString("groupname")
	sms, err := c.GetInt("smsOptions")
	if err != nil {
		c.Abort("401")
	}
	email, err := c.GetInt("emailOptions")
	if err != nil {
		c.Abort("401")
	}
	comment := c.GetString("comment")
	err = models.UpdateAlarmGroup(groupId, sms, email, groupName, comment, groupSelect)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Redirect(302, "/alarm/alarmgroup")
	}
	c.Ctx.Redirect(302, "/alarm/alarmgroup")
}

type AlarmGroupDel struct {
	BaseController
}

func (c *AlarmGroupDel) Get() {
	groupId := c.GetString("id")
	groupIdList := strings.Split(groupId, ",")
	for _, gid := range groupIdList {
		err := models.DelAlarmGroup(gid)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	c.Ctx.Redirect(302, "/alarm/alarmgroup")
}

type AlarmGroupUsers struct {
	BaseController
}

func (c *AlarmGroupUsers) Get() {
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
	userIds := models.GetAlarmGroupIds(groupId)
	pagination.SetPaginator(c.Ctx, offset, int64(len(userIds)))
	userList, err := models.GetPageUserListByIds(userIds, page, offset, keyword)
	if err != nil {
		fmt.Println(err)
		c.Abort("401")
	}
	c.Data["users"] = userList
	c.Data["activeid"] = "alarmset"
	c.Data["class"] = "alarmgrp"
	c.Layout = "inc/base.tpl"
	c.TplName = "users/userList.tpl"
}
