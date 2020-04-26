/** layuiAdmin.std-v1.2.1 LPPL License By http://www.layui.com/admin/ */
;layui.define(function (e) {
    layui.use(["carousel"], function () {
        var e = layui.$, a = layui.carousel;
        e(".layadmin-carousel").each(function () {
            var l = e(this);
            a.render({
                elem: this,
                width: "100%",
                arrow: "none",
                interval: l.data("interval"),
                autoplay: l.data("autoplay") === !0,
                anim: l.data("anim")
            })
        })
    }), layui.use(["echarts"], function () {
        var e = layui.$, a = layui.echarts, xData = [], l = [];

        e.ajaxSettings.async = false;
        var date = e("#date-range").val();
        var sid = e("#sid").val();
        e.post('/monitor/onlineavg', {date: date, sid: sid}, function (res) {
            xData = res.xData;
            yDataNew = res.yDataNew;
            yDataActive = res.yDataActive;
        }, "json");

        var tp = [{
            tooltip: {trigger: "axis"},
            dataZoom: {show: !0, realtime: !0, start: 0, end: 100},
            legend: {data: ["新增用户平均在线时长", "活跃用户平均在线时长"]},
            xAxis: [{
                type: "category",
                boundaryGap: false,
                axisLine: {onZero: !1},
                data: xData
            }],
            yAxis: [
                {
                    type: "value", name: "平均在线时长(秒)", axisLabel: {formatter: "{value}"}
                }
            ],
            series: [
                {
                    name: "新增用户平均在线时长",
                    type: "line",
                    data: yDataNew
                },
                {
                    name: "活跃用户平均在线时长",
                    type: "line",
                    data: yDataActive
                }
            ]
        }];

        var ip = e("#LAY-index-avg").children("div"), np = function (e) {
            l[e] = a.init(ip[e], layui.echartsTheme), l[e].setOption(tp[e]), window.onresize = l[e].resize
        };
        ip[0] && np(0);


    }), e("avg", {})
});