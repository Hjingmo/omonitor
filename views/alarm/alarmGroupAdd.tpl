<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-10">
            <div class="ibox float-e-margins">
                <div id="ibox-content" class="ibox-title">
                    <h5> 填写报警组基本信息 </h5>
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
                        <div class="form-group"><label class="col-sm-2 control-label"> 组名 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="groupname" placeholder="" class="form-control"></div>
                        </div>
						<div class="hr-line-dashed"></div>
						<div class="form-group"><label class="col-sm-2 control-label"> 发送短信 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8">
							    <label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio1" value="0" checked> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio2" value="1"> 发送
								</label>
							</div>
                        </div>
                        <div class="hr-line-dashed"></div>
						<div class="form-group"><label class="col-sm-2 control-label"> 发送邮件 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8">
							    <label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio1" value="0"> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio2" value="1" checked> 发送
								</label>
							</div>
                        </div>
						<div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 备注 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="comment" placeholder="" class="form-control"></div>
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
        $("#alarmset").addClass('active');
        $(".alarmgrp").addClass('active');
		$('#userForm').validator({
            timely: 2,
            theme: "yellow_right_effect",
            fields: {
                "groupname": {
                    rule: "required",
                    tip: "输入组名",
                    ok: "",
                    msg: {required: "组名必须填写!"},
                    data: {'data-ok':"组名可以使用"}
                }
            },
            valid: function(form) {
                form.submit();
            }
        });
    });
</script>