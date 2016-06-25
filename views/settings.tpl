{{template "head.tpl" .}}
<div class="container">
    <div class="main">
        <div class="page-header">
            <h1>设置</h1>
        </div>
        {{if .message}}
        <div class="alert {{.messagetype}}" role="alert">
            {{.message}}
        </div>
        {{end}}

        <p class="text-primary">当前上限:{{.nowuplimit}}&nbsp;&nbsp;当前下限:{{.nowlowlimit}}</p>
        <br/>
        <form class="form-inline" action="/settings" method="post">
          <div class="form-group">
            <label for="uplimt">上限(mm)</label>
            <input type="text" class="form-control" name="uplimit" id="uplimit" placeholder="">
          </div>
          &nbsp;
          <div class="form-group">
            <label for="lowlimt">下限(mm)</label>
            <input type="text" class="form-control" name="lowlimit" id="lowlimit" placeholder="">
          </div>
          <button type="submit" class="btn btn-default">设置</button>
        </form>
    </div>
</div>
</body>
</html>
