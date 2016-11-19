<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-10">
            <div class="ibox float-e-margins">
                <div id="ibox-content" class="ibox-title">
                    <h5> 填写用户基本信息 </h5>
                    <div class="ibox-tools">
                        <a class="collapse-link">
                            <i class="fa fa-chevron-up"></i>
                        </a>
                        <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                            <i class="fa fa-wrench"></i>
                        </a>
                        <ul class="dropdown-menu dropdown-user">
                        </ul>
                        <a class="close-link">
                            <i class="fa fa-times"></i>
                        </a>
                    </div>
                </div>
                <div class="ibox-content">
                    {{if .emg}}
                        <div class="alert alert-warning text-center">{{.emg}}</div>
                    {{end}}
                    {{if .smg}}
                        <div class="alert alert-success text-center">{{.smg}}</div>
                    {{end}}
                    <form id="userForm" method="post" class="form-horizontal">
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 用户名 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="username" placeholder="" class="form-control"></div>
                        </div>
						<div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 密码 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="password" name="password" placeholder="" class="form-control"></div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group">
						    <label class="col-sm-2 control-label"> 姓 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-3"><input type="text" name="firstname" placeholder="" class="form-control"></div>
							<label class="col-sm-2 control-label"> 名 <span class="red-fonts">*</span> </label>
							<div class="col-sm-3"><input type="text" name="lastname" placeholder="" class="form-control"></div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 头像 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="avatar" placeholder="" class="form-control"></div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 状态 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="status" placeholder="1" class="form-control"></div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 是否管理 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="superuser" placeholder="0 or 1" class="form-control"></div>
                        </div>
						<div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 手机号码 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="mobile" placeholder="" class="form-control"></div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 邮箱 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="email" name="email" placeholder="" class="form-control"></div>
                        </div>

                        <div class="form-group">
                            <div class="col-sm-4 col-sm-offset-5">
                                <button class="btn btn-white" type="reset"> 重置 </button>
                                <button class="btn btn-primary" type="submit"> 提交 </button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    $(function(){
        $("#users").addClass('active');
        $(".userlist").addClass('active');
		$('#userForm').validator({
            timely: 2,
            theme: "yellow_right_effect",
            fields: {
                "username": {
                    rule: "required",
                    tip: "输入用户名",
                    ok: "",
                    msg: {required: "用户名必须填写!"},
                    data: {'data-ok':"用户名可以使用"}
                }
            },
            valid: function(form) {
                form.submit();
            }
        });
    });
</script>