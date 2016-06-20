{{template "head.tpl" .}}
<div class="container">


    <div class="main">
        <div class="page-header">
            <h1>实时数据</h1>
        </div>
        <p class="text-primary">当前上限:{{.nowuplimit}}&nbsp;&nbsp;当前下限:{{.nowlowlimit}}</p>
        <div id="container" style="min-width:200px;height:500px;"></div>
    </div>

</div>



<script>
    $(function () {
        $(document).ready(function () {
            Highcharts.setOptions({
                global: {
                    useUTC: false
                }
            });

            var chart;
            //图形样式
            chart : {

            }
            $('#container').highcharts({

                chart: {
                    type: 'spline',
                    animation: Highcharts.svg,
                    marginRight: 10,
                    events: {
                        load: function () {

                            // 每秒通过ajax获取新数据
                            var series = this.series[0];
                            setInterval(function () {
                                // $.getJSON('/dev/gtdata', function (data) {});
                                $.getJSON('/data/now', function (data) {
                                    console.log("获取实时数据",data.x,data.y);
                                    series.addPoint([data.x*1000, data.y], true, true);
                                });
                            }, 1000);
                        }
                    }
                },
                title: {
                    text: '实时水位数据'
                },
                xAxis: {
                    type: 'datetime',
                    tickPixelInterval: 500,

                },
                yAxis: {
                    title: {
                        text: '水位'
                    },
                    plotLines: [{
                        value: 0,
                        width: 1,
                        color: '#808080'
                    }],
                    //警告线
                    // plotLines: [{
                    //     color: 'red',            //线的颜色，定义为红色
                    //     dashStyle: 'Dot',//标示线的样式，默认是solid（实线），这里定义为长虚线
                    //     value: {{.nowuplimit}},                //定义在哪个值上显示标示线，这里是在x轴上刻度为3的值处垂直化一条线
                    //     width: 2                 //标示线的宽度，2px
                    // }],
                },
                tooltip: {
                    formatter: function () {
                        return '<b>' + this.series.name + '</b><br/>' +
                                Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                                Highcharts.numberFormat(this.y, 2);
                    },

                },
                legend: {
                    enabled: false
                },
                exporting: {
                    enabled: false
                },
                //版权信息
                credits: {
                    enabled: false,
                },
                series: [{

                    name: '实时水位',
                    zones: [{
                        value: {{.nowlowlimit}},
                        color: '#f7a35c',
                    }, {
                        value: {{.nowuplimit}},
                        color: '#7cb5ec'
                    }, {
                        color: '#f00'
                    }],
                    data: (function () {
                        // 初始化20个数据点
                        var data = [],
                                time = (new Date()).getTime(),
                                i;

                        for (i = -19; i <= 0; i++) {
                            data.push({
                                x: time + i * 1000,
                                // y: Math.random()
                                y: 0
                            });
                        }
                        return data;
                    })()
                }]
            });
        });

    });
</script>
</body>
</html>
