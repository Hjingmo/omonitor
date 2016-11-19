<script type="text/javascript">
    function search_project(text, noselect, total){
        $("#" + noselect).children().each(
            function(){
                $(this).remove();
            });

        $("#" + total).children().each(function(){
            if($(this).text().search(text) != -1){
                $("#" + noselect).append($(this).clone())
            }
            })
    }
</script>

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

                <select id="project_no_select_total" name="projects" class="form-control m-b" size="12" multiple style="display: none">
                    {{range $k,$v := .groupNoSelect}}
                        <option value="{{$v.Id}}">{{$v.Username}}</option>
                    {{end}}
                </select>

                <select id="project_select_total" name="om_projects" class="form-control m-b" size="12"  multiple style="display: none">
                    {{range $k,$v := .groupSelect}}
                        <option value="{{$v.Id}}">{{$v.Username}}</option>
                    {{end}}
                </select>

                <div class="ibox-content">
                    {{if .emg}}
                        <div class="alert alert-warning text-center">{{.emg}}</div>
                    {{end}}
                    {{if .smg}}
                        <div class="alert alert-success text-center">{{.smg}}</div>
                    {{end}}
                    <form id="projectForm" method="post" class="form-horizontal">
                        <div class="form-group"><label class="col-sm-2 control-label"> 组名 <span class="red-fonts">*</span></label>
                            <div class="col-sm-8"><input type="text" value="{{.group.Group}}" name="groupname" class="form-control"></div>
                        </div>

                        <div class="hr-line-dashed"></div>
                        <div class="form-group">
                            <label for="group_name" class="col-sm-2 control-label">过滤</label>
                            <div class="col-sm-4">
                                <input id="noselect" class="form-control" oninput="search_project(this.value, 'projects', 'project_no_select_total')">
                            </div>
                            <div class="col-sm-1">
                            </div>
                            <div id="select" class="col-sm-3">
                                <input  class="form-control" oninput="search_project(this.value, 'project_select', 'project_select_total')">
                            </div>
                        </div>


                        <div class="form-group">
                            <label for="" class="col-sm-2 control-label">用户<span class="red-fonts">*</span></label>
                            <div class="col-sm-4">
                                <div>
                                    <select id="projects" name="projects" class="form-control m-b" size="12" multiple>
                                        {{range $k,$v := .groupNoSelect}}
                                            <option value="{{$v.Id}}">{{$v.Username}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            </div>

                            <div class="col-sm-1">
                                <div class="btn-group" style="margin-top: 60px;">
                                    <button type="button" class="btn btn-white" onclick="move('projects', 'project_select', 'project_no_select_total', 'project_select_total'  )"><i class="fa fa-chevron-right"></i></button>
                                    <button type="button" class="btn btn-white" onclick="move_left('project_select', 'projects', 'project_select_total', 'project_no_select_total')"><i class="fa fa-chevron-left"></i> </button>
                                </div>
                            </div>

                            <div class="col-sm-3">
                                <div>
                                    <select id="project_select" name="project_select" class="form-control m-b" size="12"  multiple>
                                        {{range $k,$v := .groupSelect}}
                                            <option value="{{$v.Id}}">{{$v.Username}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            </div>
                        </div>

                        <div class="hr-line-dashed"></div>
						<div class="form-group"><label class="col-sm-2 control-label"> 发送短信 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8">
							{{if eq .group.Sms 0}}
							    <label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio1" value="0" checked> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio2" value="1"> 发送
								</label>
							{{else}}
							    <label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio1" value="0"> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="smsOptions" id="smsRadio2" value="1" checked> 发送
								</label>
							{{end}}
							</div>
                        </div>
                        <div class="hr-line-dashed"></div>
						<div class="form-group"><label class="col-sm-2 control-label"> 发送邮件 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8">
							{{if eq .group.Email 0}}
							    <label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio1" value="0" checked> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio2" value="1"> 发送
								</label>
							{{else}}
							    <label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio1" value="0"> 不发送
								</label>
								<label class="radio-inline">
								    <input type="radio" name="emailOptions" id="emailRadio2" value="1" checked> 发送
								</label>
							{{end}}
							</div>
                        </div>
                        <div class="hr-line-dashed"></div>
                        <div class="form-group"><label class="col-sm-2 control-label"> 备注 <span class="red-fonts">*</span> </label>
                            <div class="col-sm-8"><input type="text" name="comment" value="{{.group.Comment}}" class="form-control"></div>
                        </div>

                        <div class="hr-line-dashed"></div>
                        <div class="form-group">
                            <div class="col-sm-4 col-sm-offset-5">
                                <button class="btn btn-white" type="submit"> 重置 </button>
                                <button class="btn btn-primary" id="submit_button" type="submit" onclick="on_submit('groups_selected')  "> 提交 </button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    $(document).ready(function(){
        $("#alarmset").addClass('active');
        $(".alarmgrp").addClass('active');

        $("#submit_button").click(function(){
            $('#project_select option').each(function(){
                $(this).prop('selected', true);
            });
        });

    });

    $('#projectForm').validator({
            timely: 2,
            theme: "yellow_right_effect",
            rules: {
                check_ip: [/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){3}$/, 'ip地址不正确'],
                check_port: [/^\d{1,5}$/, '端口号不正确']
            },
            fields: {
                "server_ip": {
                    rule: "check_ip;",
                    tip: "输入IP",
                    ok: "",
                    msg: {required: "必须填写!"}
                },
                "server_user": {
                    tip: "输入管理用户",
                    ok: "",
                    msg: {required: "必须填写!"}
                },
                "server_port": {
                    rule: "check_port",
                    tip: "输入端口",
                    ok: "",
                    msg: {required: "必须填写!"}
                }
            },
            valid: function(form) {
                form.submit();
            }
    });

    function on_submit(id){
        search_project('', 'project_select', 'project_select_total');  //提交之前清空过滤框
        $('#'+id+' option').each(
            function(){
                $(this).prop('selected', true)
            })
        }

</script>