<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>OMonitor</title>

    <link rel="shortcut icon" href="/static/img/facio.ico" type="image/x-icon">
    {{template "inc/link_css.html" .}}
	{{template "inc/head_script.html" .}}
</head>

<body>

<div id="wrapper">
    {{template "inc/nav.html" .}}
    <div id="page-wrapper" class="gray-bg">
        <div class="row border-bottom">
			{{template "inc/nav_bar_header.html" .}}
        </div>
		{{.LayoutContent}}
    </div>
</div>

</body>
    {{template "inc/foot_script.html" .}}
</html>
