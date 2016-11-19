package controllers

import (
	"omonitor/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type BaseController struct {
	beego.Controller
	IsLogin       bool
	UserUserId    int64
	UserUsername  string
	User_Name     string
	UserAvatar    string
	UserSuperuser int
}

func init() {
	AccessRegister()
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userInfo")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userInfo")).(string), "||")
		userid, _ := strconv.Atoi(tmp[0])
		is_superuser, _ := strconv.Atoi(tmp[4])
		longid := int64(userid)
		this.Data["UserId"] = longid
		this.Data["UserName"] = tmp[1]
		this.Data["user_name"] = tmp[2]
		this.Data["UserAvatar"] = tmp[3]

		this.UserUserId = longid
		this.UserUsername = tmp[1]
		this.User_Name = tmp[2]
		this.UserAvatar = tmp[3]
		this.UserSuperuser = is_superuser
	}
	this.Data["IsLogin"] = this.IsLogin
}

func AccessRegister() {
	var Check = func(ctx *context.Context) {
		userinfoSession := ctx.Input.Session("userInfo")
		if userinfoSession != nil {
			params := strings.Split(strings.ToLower(ctx.Request.RequestURI), "/")
			if len(params) > 2 {
				if params[1] == "skin_config" {
					return
				}
				permission := params[1] + "_" + strings.Split(params[2], "?")[0]
				userId, _ := strconv.Atoi(strings.Split(userinfoSession.(string), "||")[0])
				if userId == 1 {
					return
				}
				ret, err := models.CheckPermission(userId, permission)
				if err != nil {
					ctx.Redirect(302, "/permission")
				}
				if !ret {
					ctx.Redirect(302, "/permission")
				}
			}
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Check)
}
