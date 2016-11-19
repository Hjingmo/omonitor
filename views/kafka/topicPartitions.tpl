<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-10">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5> {{.topic}} 详细信息</h5>
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
                    {{template "kafka/kafkaNav.tpl" .}}
                    {{template "headbtn" .}}
                    <form id="search_form" method="get" action="" class="pull-right mail-search">
                        <div class="input-group">
                            <input type="text" class="form-control input-sm" id="search_input" name="keyword" placeholder="Search">
                            <input type="text" style="display: none" name="env" value="{{.env}}">
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
                                <th class="text-center"> Partition </th>
								<th class="text-center"> Offset </th>
								<th class="text-center"> LogSize </th>
								<th class="text-center"> Lag </th>
								<th class="text-center"> Owner </th>
                            </tr>
                        </thead>
                        <tbody>
						    {{range $k,$v := .partitions}}
                            <tr class="gradeX">
                                <td class="text-center"> {{ $v.Partition }}</td>
								<td class="text-center"> {{ $v.Offset }}</td>
								<td class="text-center"> {{ $v.LogSize }}</td>
								<td class="text-center"> {{ $v.Lag }}</td>
								<td class="text-center"> {{ $v.Owner }}</td>
                            </tr>
                            {{end}}
                    </table>
                   </form>
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    $(document).ready(function(){
        $("#kafka").addClass('active');
        $(".kafka{{.env}}").addClass('active');
    });
</script>