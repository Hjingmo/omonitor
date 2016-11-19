<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> 登录 </title>
    <link rel="shortcut icon" href="/static/img/facio.ico" type="image/x-icon">
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/font-awesome/css/font-awesome.css" rel="stylesheet">
    <link href="/static/css/animate.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">
</head>

<body class="black-bg">
    <div class="middle-box text-center loginscreen  animated fadeInDown">
        <div>
            <div>
                <h1 class="logo-name"><img src="/static/img/logo.png"></h1>
            </div>
			{{if .error}}
                <div class="alert alert-danger text-center">{{.error}}</div>
            {{end}}
            <h2>欢迎登录</h2>
            <form class="m-t" role="form" method="post" action="">
                <div class="form-group">
                    <input type="text" name="username" class="form-control" placeholder="Username" required="length[6~50]">
                </div>
                <div class="form-group">
                    <input type="password" name="password" class="form-control" placeholder="Password" required="">
                </div>
                <button type="submit" class="btn btn-primary block full-width m-b">登录</button>

            </form>
            <p class="m-t"> <small><b>Copyright</b> omserver.com Organization © 2015-2016</small> </p>
        </div>
    </div>

    <!-- Mainly scripts -->
    <script src="/static/js/jquery-2.1.1.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>

</body>

</html>
