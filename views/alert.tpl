{{template "head.tpl" .}}
<div class="container">

    <div class="main">
        <div class="page-header">
            <h1>历史数据</h1>
        </div>
        <p class="text-primary">当前上限:{{.nowuplimit}}&nbsp;&nbsp;当前下限:{{.nowlowlimit}}</p>
                <table class="table table-bordered table-hover table-striped responsive-utilities alert-table">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>时间</th>
                        <th>水位</th>
                        <th>报警类型</th>
                    </tr>
                    </thead>
                    <tbody>
                        {{range .data}}
                        <tr>
                            <th scope="row">{{.Id}}</th>
                            <td>{{.Date}}</td>
                            <td>{{.Value}}</td>
                            <td>
                                {{if gt .Value $.nowuplimit}}
                                    高于上限
                                    {{else if lt .Value $.nowlowlimit}}
                                    低于下限
                                    {{else}}
                                    正常
                                {{end}}
                            </td>
                        </tr>
                    {{end}}
                    </tbody>
                </table>
                <nav class="text-right">
                    <ul class="pagination">
                        {{if ne .nowpn 1}}
                        <li>
                            <a href="/alert/1" aria-label="Previous">
                                <span aria-hidden="true">首页</span>
                            </a>
                        </li>
                        {{end}}

                        {{range $index:=.pn1}}
                            <li {{if eq $.nowpn $index}}class="active"{{end}}><a href="/alert/{{$index}}">{{$index}}</a></li>
                        {{end}}

                        {{if ne .nowpn .maxpagenums}}
                        <li>
                            <a href="/alert/{{.maxpagenums}}" aria-label="Next">
                                <span aria-hidden="true">尾页</span>
                            </a>
                        </li>
                        {{end}}
                    </ul>
                </nav>
            </div>

    </div>



</body>

</html>
