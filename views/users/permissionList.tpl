<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-10">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5> 权限信息列表</h5>
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
                    <div class="">
                    <a href="/user/addpermission" class="btn btn-sm btn-primary "> 添加权限 </a>
                    <input type="button" id="del_check" class="btn btn-danger btn-sm"  name="del_button" value="删除所选"/>
                    <form id="search_form" method="get" action="" class="pull-right mail-search">
                        <div class="input-group">
                            <input type="text" class="form-control input-sm" id="search_input" name="keyword" placeholder="Search">
                            <input type="text" style="display: none">
                            <div class="input-group-btn">
                                <button id='search_btn' type="submit" class="btn btn-sm btn-primary">
                                    - 搜索 -
                                </button>
                            </div>
                        </div>
                    </form>
                    </div>

                    <form id="contents_form" name="contents_form">
                    <table class="table table-striped table-bordered table-hover " id="editable" >
                        <thead>
                            <tr>
                                <th class="text-center"><input id="checkall" type="checkbox" class="i-checks" name="checkall" value="checkall" data-editable='false' onclick="check_all('contents_form')"></th>
                                <th class="text-center"> 名称 </th>
								<th class="text-center"> 备注 </th>
                                <th class="text-center"> 操作 </th>
                            </tr>
                        </thead>
                        <tbody>
                        {{range $k,$v := .permissions}}
                        <tr class="gradeX">
							<td class="text-center" name="om_id" value="{{$v.Id}}" data-editable='false'><input name="id" value="{{$v.Id}}" type="checkbox" class="i-checks"></td>
							<td class="text-center"> {{$v.Codename}} </td>
							<td class="text-center"> {{$v.Comment}} </td>
							<td class="text-center">
	                            <a href="/user/editpermission?id={{$v.Id}}" class="btn btn-xs btn-info">编辑</a>
	                            <a value="/user/delpermission?id={{$v.Id}}" class="btn btn-xs btn-danger server_del">删除</a>
	                        </td>
                        </tr>
						{{end}}
                        </tbody>
                    </table>
					<div class="row">
                        <div class="col-sm-6">
                            <div class="dataTables_info" id="editable_info" role="status" aria-live="polite">
                                Showing start to end of count entries
                            </div>
                        </div>
						<div class="col-sm-6">
                            {{template "inc/paginator.html" .}}
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
        $("#users").addClass('active');
        $(".permission").addClass('active');

        $('.server_del').click(function(){
            var row = $(this).closest('tr');
            if (confirm('确定删除?')) {
                $.get(
                        $(this).attr('value'),
                        {},
                        function (data) {
                            row.remove();
                        }
                );
                return false
            }
        });

        $('#del_check').click(function(){
            var check_array = [];
            if (confirm('确定删除?')){
                $('tr.gradeX input:checked').each(function(){
                    check_array.push($(this).attr('value'))
                });
                $.get(
                        '/user/delpermission',
                        {id: check_array.join(',')},
                        function(data){
                            $('tr.gradeX input:checked').closest('tr').remove();
                        }
                );
                return false;
            }
        })
    });
</script>