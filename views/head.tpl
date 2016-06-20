<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.pagetitle}} | 水位监控平台</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/style.css">
    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/highstock.js"></script>

    <script src="/static/js/plugins/highcharts-zh_CN.js"></script>
</head>
<body>
    <div class="container">
        <nav class="navbar navbar-default">
            <div class="container-fluid">
                <div class="navbar-header">

                    <a class="navbar-brand" href="/"><b>水位监测系统</b></a>
                </div>
                <div id="navbar" class="navbar-collapse collapse">

                    <ul class="nav navbar-nav navbar-right">
                        <li {{if .IsHome}}class="active"{{end}}><a href="/">实时数据</a></li>
                        <li {{if .IsHistory}}class="active"{{end}}><a href="/historydata">历史数据</a></li>
                        <li {{if .IsAlert}}class="active"{{end}}><a href="/alert">超限记录</a></li>
                        <li {{if .IsSettings}}class="active"{{end}}><a href="/settings">设置</a></li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>
