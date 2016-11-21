package controllers

import (
	"fmt"
	"strings"

	"github.com/yangji168/omonitor/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Layout = "inc/base.tpl"
	this.TplName = "index.tpl"
}

type SkinConfigController struct {
	BaseController
}

func (this *SkinConfigController) Get() {
	this.TplName = "skin_config.html"
}

type PermissionController struct {
	BaseController
}

func (this *PermissionController) Get() {
	this.Layout = "inc/base.tpl"
	this.TplName = "users/permission.tpl"
}

type LoginUserController struct {
	BaseController
}

func (this *LoginUserController) Get() {
	check := this.BaseController.IsLogin
	if check {
		this.Abort("401")
	} else {
		this.TplName = "users/login.tpl"
	}
}

func (this *LoginUserController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	err, users := models.LoginUser(username, password)
	if err == nil {
		this.SetSession("userInfo", fmt.Sprintf("%d", users.Id)+"||"+users.Username+"||"+users.Firstname+users.Lastname+"||"+users.Avatar+"||"+fmt.Sprintf("%d", users.Superuser))
		models.UPdateUserLastLogin(users.Id)
		this.Ctx.Redirect(302, "/")
	} else {
		this.Data["error"] = "用户名或密码错误"
		this.TplName = "users/login.tpl"
	}
}

type LogoutUserController struct {
	BaseController
}

func (c *LogoutUserController) Get() {
	c.DelSession("userInfo")
	c.Ctx.Redirect(302, "/login")
}

type UserListController struct {
	BaseController
}

func (c *UserListController) Get() {
	keyword := c.GetString("keyword")
	page, err := c.GetInt("p")
	if err != nil {
		page = 1
	}
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	countUsers := models.CountUsers()
	pagination.SetPaginator(c.Ctx, offset, countUsers)
	users, err := models.GetUserList(page, offset, keyword)
	if err != nil {
		c.Abort("401")
	}
	c.Data["users"] = users
	c.Data["activeid"] = "users"
	c.Data["class"] = "userlist"
	c.Layout = "inc/base.tpl"
	c.TplName = "users/userList.tpl"
}

type UserAddController struct {
	BaseController
}

func (c *UserAddController) Get() {
	c.Layout = "inc/base.tpl"
	c.TplName = "users/userAdd.tpl"
}

func (c *UserAddController) Post() {
	c.Layout = "inc/base.tpl"
	c.TplName = "users/userAdd.tpl"
	userName := c.GetString("username")
	password := c.GetString("password")
	firstName := c.GetString("firstname")
	lastName := c.GetString("lastname")
	avatar := c.GetString("avatar")
	status, serr := c.GetInt("status")
	if serr != nil {
		status = 1
	}
	superuser, serr := c.GetInt("superuser")
	if serr != nil {
		superuser = 1
	}
	mobile, merr := c.GetInt64("mobile")
	if merr != nil {
		mobile = 0
	}
	email := c.GetString("email")
	if userName == "" {
		c.Data["emg"] = "用户名不能为空"
		return
	}
	userTest := models.GetUserByName(userName)
	if userTest {
		var user models.Users
		user.Username = userName
		user.Password = password
		user.Firstname = firstName
		user.Lastname = lastName
		user.Avatar = avatar
		user.Status = status
		user.Superuser = superuser
		user.Mobile = mobile
		user.Email = email
		err := models.AddUser(user)
		if err != nil {
			c.Abort("401")
		} else {
			c.Data["smg"] = userName + "用户添加成功"
		}
	} else {
		c.Data["emg"] = userName + "用户名已存在"
	}
}

type UserEditController struct {
	BaseController
}

func (c *UserEditController) Get() {
	userId, err := c.GetInt64("id")
	if err != nil {
		c.Abort("401")
	}
	projectNoSelect := models.GetUserPermissionNoSelect(userId)
	projectSelect := models.GetUserPermissionSelect(userId)
	user, err1 := models.GetUserById(userId)
	if err1 != nil {
		c.Abort("401")
	}
	c.Data["user"] = user
	c.Data["project_no_select"] = projectNoSelect
	c.Data["project_select"] = projectSelect
	c.Layout = "inc/base.tpl"
	c.TplName = "users/userEdit.tpl"
}

func (c *UserEditController) Post() {
	userId, err := c.GetInt64("id")
	if err != nil {
		c.Abort("401")
	}
	userName := c.GetString("username")
	firstName := c.GetString("firstname")
	lastName := c.GetString("lastname")
	avatar := c.GetString("avatar")
	status, serr := c.GetInt("status")
	if serr != nil {
		c.Abort("401")
	}
	superuser, serr := c.GetInt("superuser")
	if serr != nil {
		superuser = 1
	}
	mobile, merr := c.GetInt64("mobile")
	if merr != nil {
		mobile = 0
	}
	email := c.GetString("email")
	projectSelect := c.GetStrings("project_select")
	var user models.Users
	user.Username = userName
	user.Firstname = firstName
	user.Lastname = lastName
	user.Avatar = avatar
	user.Status = status
	user.Superuser = superuser
	user.Mobile = mobile
	user.Email = email
	err1 := models.UpdateUserPermission(userId, user, projectSelect)
	if err1 != nil {
		c.Abort("401")
	}
	c.Ctx.Redirect(302, "/user/manage")
}

type UserDelController struct {
	BaseController
}

func (c *UserDelController) Get() {
	userId := c.GetString("id")
	userIdList := strings.Split(userId, ",")
	for _, uid := range userIdList {
		models.DelUser(uid)
	}
	c.Ctx.Redirect(302, "/user/manage")
}

type PermissionListController struct {
	BaseController
}

func (c *PermissionListController) Get() {
	page, err := c.GetInt("p")
	if err != nil {
		page = 1
	}
	keyword := c.GetString("keyword")
	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}
	countPermission := models.CountPermission()
	pagination.SetPaginator(c.Ctx, offset, countPermission)
	permissions, err2 := models.GetPermissionList(page, offset, keyword)
	if err2 != nil {
		c.Abort("401")
	}
	c.Data["permissions"] = permissions
	c.Layout = "inc/base.tpl"
	c.TplName = "users/permissionList.tpl"
}

type PermissionAddController struct {
	BaseController
}

func (c *PermissionAddController) Get() {
	c.Layout = "inc/base.tpl"
	c.TplName = "users/permissionAdd.tpl"
}

func (c *PermissionAddController) Post() {
	c.Layout = "inc/base.tpl"
	c.TplName = "users/permissionAdd.tpl"
	codeName := c.GetString("codename")
	comment := c.GetString("comment")
	_, terr := models.GetPermissionByName(codeName)
	if terr == nil {
		c.Data["emg"] = codeName + "已存在"
		return
	}
	ierr := models.AddPermission(codeName, comment)
	if ierr != nil {
		c.Abort("401")
	} else {
		c.Data["smg"] = codeName + "添加成功"
	}
}

type PermissionEditController struct {
	BaseController
}

func (c *PermissionEditController) Get() {
	pId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	permission, err1 := models.GetPermissionById(pId)
	if err1 != nil {
		c.Abort("401")
	}
	c.Data["permission"] = permission
	c.Layout = "inc/base.tpl"
	c.TplName = "users/permissionEdit.tpl"
}

func (c *PermissionEditController) Post() {
	pId, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
	}
	codeName := c.GetString("codename")
	comment := c.GetString("comment")
	err1 := models.UpdatePermission(pId, codeName, comment)
	if err1 != nil {
		c.Abort("401")
	}
	c.Ctx.Redirect(302, "/user/permission")
}

type PermissionDelController struct {
	BaseController
}

func (c *PermissionDelController) Get() {
	pId := c.GetString("id")
	pIdList := strings.Split(pId, ",")
	for _, pid := range pIdList {
		models.DelPermission(pid)
	}
	c.Ctx.Redirect(302, "/user/permission")
}
