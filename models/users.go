package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yangji168/omonitor/utils"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id        int64
	Username  string
	Firstname string
	Lastname  string
	Password  string
	Avatar    string
	Status    int
	Superuser int
	Lastlogin string
	Mobile    int64
	Email     string
}

type Permissions struct {
	Id       int
	Codename string
	Comment  string
}

type User_permissions struct {
	Id           int
	Userid       int64
	Permissionid int
}

func init() {
	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Permissions))
	orm.RegisterModel(new(User_permissions))
}

func LoginUser(username, password string) (err error, user Users) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	cond := orm.NewCondition()
	pwdmd5 := utils.Md5(password)
	cond = cond.And("Username", username)
	cond = cond.And("Password", pwdmd5)
	cond = cond.And("Status", 1)

	qs = qs.SetCond(cond)
	var users Users
	err = qs.Limit(1).One(&users)
	return err, users
}

func CountUsers() int64 {
	var user []Users
	num, _ := orm.NewOrm().QueryTable("users").All(&user)
	return num
}

func GetUserList(page int, offset int, keyword string) (user []Users, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("Username__icontains", keyword)).Or("Lastname__icontains", keyword)
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()
	var userlist []Users
	_, err = qs.Limit(offset, start).All(&userlist)
	return userlist, err
}

func GetPageUserListByIds(userIds []int64, page int, offset int, keyword string) (user []Users, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("Username__icontains", keyword)).Or("Lastname__icontains", keyword)
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()
	var userlist []Users
	if len(userIds) < 1 {
		return userlist, nil
	}
	_, err = qs.Filter("id__in", userIds).Limit(offset, start).All(&userlist)
	return userlist, err
}

func GetAllUsers() (all []Users, err error) {
	var userlist []Users
	_, err = orm.NewOrm().QueryTable("users").All(&userlist)
	return userlist, err
}

func GetUserListByIds(userIds []int64) (ulist []Users, err error) {
	var userlist []Users
	if len(userIds) < 1 {
		return userlist, nil
	}
	_, err = orm.NewOrm().QueryTable("users").Filter("id__in", userIds).All(&userlist)
	return userlist, err
}

func GetUserByName(userName string) bool {
	result := false
	var user Users
	err := orm.NewOrm().QueryTable("users").Filter("Username", userName).One(&user)
	if err != nil {
		result = true
	}
	return result
}

func GetUserById(userId int64) (u Users, err error) {
	var user Users
	err = orm.NewOrm().QueryTable("users").Filter("Id", userId).One(&user)
	return user, err
}

func AddUser(user Users) error {
	o := orm.NewOrm()
	add := new(Users)
	add.Username = user.Username
	add.Firstname = user.Firstname
	add.Lastname = user.Lastname
	add.Password = utils.Md5(user.Password)
	add.Avatar = user.Avatar
	add.Status = user.Status
	add.Superuser = user.Superuser
	add.Email = user.Email
	add.Mobile = user.Mobile
	_, err := o.Insert(add)
	return err
}

func UPdateUserLastLogin(userId int64) {
	timeNow := time.Now().String()
	user := Users{Id: userId}
	user.Lastlogin = timeNow
	_, err := orm.NewOrm().Update(&user, "Lastlogin")
	if err != nil {
		fmt.Println(err)
	}
}

func DelUser(userId string) {
	idInt, _ := strconv.ParseInt(userId, 10, 64)
	o := orm.NewOrm()
	_, err := o.Delete(&Users{Id: idInt})
	if err != nil {
		fmt.Println(err)
	}
}

func GetUserPermissions(userId int64) []int {
	var userPermission []User_permissions
	var result []int
	_, err := orm.NewOrm().QueryTable("user_permissions").Filter("Userid", userId).All(&userPermission)
	if err != nil {
		return result
	}
	for _, p := range userPermission {
		result = append(result, p.Permissionid)
	}
	return result
}

func CheckUserPermission(p []Permissions, pid int) bool {
	result := true
	for _, i := range p {
		if i.Id == pid {
			result = false
		}
	}
	return result
}

func GetUserPermissionNoSelect(userId int64) (np []Permissions) {
	userPermissions := GetUserPermissions(userId)
	userPermissionList, err := GetPermissionListByIds(userPermissions)
	if err != nil {
		return
	}
	pers := &userPermissionList
	allPermissionList, err1 := GetAllPermissions()
	if err1 != nil {
		return
	}
	var noselect []Permissions
	var p Permissions
	for _, a := range allPermissionList {
		chk := CheckUserPermission(*pers, a.Id)
		if chk {
			p.Id = a.Id
			p.Codename = a.Codename
			p.Comment = a.Comment
			noselect = append(noselect, p)
		}
	}
	return noselect
}

func GetUserPermissionSelect(userId int64) (sp []Permissions) {
	userPermissions := GetUserPermissions(userId)
	userPermissionList, err := GetPermissionListByIds(userPermissions)
	if err != nil {
		return
	}
	return userPermissionList
}

func CountPermission() int64 {
	var projects []Permissions
	num, _ := orm.NewOrm().QueryTable("permissions").All(&projects)
	return num
}

func GetAllPermissions() (p []Permissions, err error) {
	var permissions []Permissions
	_, err = orm.NewOrm().QueryTable("permissions").All(&permissions)
	return permissions, err
}

func GetPermissionList(page int, offset int, keyword string) (permiss []Permissions, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("permissions")
	cond := orm.NewCondition()
	if keyword != "" {
		cond = cond.AndCond(cond.And("Codename__icontains", keyword))
	}
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	if page < 1 {
		page = 1
	}
	start := (page - 1) * offset
	var permission []Permissions
	_, err = qs.Limit(offset, start).OrderBy("Codename").All(&permission)
	return permission, err
}

func GetPermissionListByIds(pids []int) (p []Permissions, err error) {
	var permissions []Permissions
	if len(pids) < 1 {
		return permissions, nil
	}
	_, err = orm.NewOrm().QueryTable("permissions").Filter("Id__in", pids).OrderBy("Codename").All(&permissions)
	return permissions, err
}

func GetPermissionByName(codeName string) (perId int, err error) {
	var permission Permissions
	err = orm.NewOrm().QueryTable("permissions").Filter("Codename", codeName).One(&permission)
	return permission.Id, err
}

func GetPermissionById(pid int) (p Permissions, err error) {
	var permission Permissions
	err = orm.NewOrm().QueryTable("permissions").Filter("Id", pid).One(&permission)
	return permission, err
}

func AddPermission(codeName string, comment string) error {
	o := orm.NewOrm()
	permission := new(Permissions)
	permission.Codename = codeName
	permission.Comment = comment
	_, err := o.Insert(permission)
	return err
}

func UpdatePermission(pid int, codeName string, comment string) error {
	o := orm.NewOrm()
	var permission Permissions
	permission = Permissions{Id: pid}
	permission.Codename = codeName
	permission.Comment = comment
	_, err := o.Update(&permission)
	return err
}

func DelPermission(pid string) {
	pidInt, _ := strconv.Atoi(pid)
	o := orm.NewOrm()
	_, err := o.Delete(&Permissions{Id: pidInt})
	if err != nil {
		fmt.Println(err)
	}
}

func CheckPermission(userId int, codeName string) (bool, error) {
	permissionId, err := GetPermissionByName(codeName)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	var userPermission User_permissions
	err1 := orm.NewOrm().QueryTable("user_permissions").Filter("Userid", userId).Filter("Permissionid", permissionId).One(&userPermission)
	if err1 != nil {
		fmt.Println(err1)
		return false, err1
	}
	return true, nil
}

func UpdateUserPermission(userId int64, user Users, selectIds []string) error {
	userPermissions := GetUserPermissions(userId)
	selectIdInts := utils.StringToInt(selectIds)
	adds, dels := utils.UpdateCompareInt(userPermissions, selectIdInts)
	if len(adds) > 0 {
		for _, add := range adds {
			aerr := AddUserPermission(userId, add)
			if aerr != nil {
				fmt.Println(aerr)
			}
		}
	}
	if len(dels) > 0 {
		for _, del := range dels {
			derr := DelUserPermission(userId, del)
			if derr != nil {
				fmt.Println(derr)
			}
		}
	}
	var u Users
	u = Users{Id: userId}
	u.Username = user.Username
	u.Firstname = user.Firstname
	u.Lastname = user.Lastname
	u.Avatar = user.Avatar
	u.Status = user.Status
	u.Superuser = user.Superuser
	u.Mobile = user.Mobile
	u.Email = user.Email
	_, err := orm.NewOrm().Update(&u, "Username", "Firstname", "Lastname", "Avatar", "Status", "Superuser", "Mobile", "Email")
	return err
}

func AddUserPermission(userId int64, pid int) error {
	o := orm.NewOrm()
	userperm := new(User_permissions)
	userperm.Userid = userId
	userperm.Permissionid = pid
	_, err := o.Insert(userperm)
	return err
}

func DelUserPermission(userId int64, pid int) error {
	var userperm User_permissions
	rerr := orm.NewOrm().QueryTable("user_permissions").Filter("Userid", userId).Filter("Permissionid", pid).One(&userperm)
	if rerr != nil {
		return rerr
	}
	o := orm.NewOrm()
	_, err := o.Delete(&User_permissions{Id: userperm.Id})
	return err
}
