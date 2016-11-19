<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-10">
            <div class="ibox float-e-margins">
                <div id="ibox-content" class="ibox-title">
                    <h5> 填写权限基本信息 </h5>
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
                        <div class="form-group"><label class="col-sm-2 control-label"> 名称 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="codename" placeholder="" class="form-control"></div>
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
        $("#users").addClass('active');
        $(".permission").addClass('active');
		$('#userForm').validator({
            timely: 2,
            theme: "yellow_right_effect",
            fields: {
                "codename": {
                    rule: "required",
                    tip: "输入名称",
                    ok: "",
                    msg: {required: "名称必须填写!"},
                    data: {'data-ok':"名称可以使用"}
                }
            },
            valid: function(form) {
                form.submit();
            }
        });
    });
</script>