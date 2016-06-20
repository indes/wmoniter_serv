{{template "head.tpl" .}}
<div class="container">

    <div class="main">
        <div class="page-header">
            <h1>历史数据</h1>
        </div>
        <p class="text-primary">当前上限:{{.nowuplimit}}&nbsp;&nbsp;当前下限:{{.nowlowlimit}}</p>
        
                <div id="container" style="min-width:400px;height:500px"></div>
                <table class="table table-bordered table-hover table-striped responsive-utilities history-table">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>时间</th>
                        <th>水位</th>
                        <th>状态</th>

                    </tr>
                    </thead>
                    <tbody>
                        {{range .data}}
                        <tr {{if gt .Value $.nowuplimit}}class="danger"{{end}} {{if lt .Value $.nowlowlimit}}class="danger"{{end}}>
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
                            <a href="/historydata/1" aria-label="Previous">
                                <span aria-hidden="true">首页</span>
                            </a>
                        </li>
                        {{end}}

                        {{range $index:=.pn1}}
                            <li {{if eq $.nowpn $index}}class="active"{{end}}><a href="/historydata/{{$index}}">{{$index}}</a></li>
                        {{end}}

                        {{if ne .nowpn .maxpagenums}}
                        <li>
                            <a href="/historydata/{{.maxpagenums}}" aria-label="Next">
                                <span aria-hidden="true">尾页</span>
                            </a>
                        </li>
                        {{end}}
                    </ul>
                </nav>
            </div>

    </div>


<script>
    $(function () {
        $.getJSON('/data/all', function (data) {
            $('#container').highcharts('StockChart', {

                exporting: {
                    enabled: false
                },

                rangeSelector : {
                    selected : 1
                },
                credits: {
                    enabled: false,
                },
                yAxis: {
                    title: {
                        text: '水位'                  //指定y轴的标题
                    }
                },
                title : {
                    text : '历史数据'
                },

                series : [{
                    name : '水位',
                    data :data,
                    tooltip: {
                        valueDecimals: 2
                    }
                }]
            });
        });

    });

</script>
</body>

</html>
