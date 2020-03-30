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
        var e = layui.$, a = layui.echarts, xData = [], yData = [], l = [];

        e.ajaxSettings.async = false;
        var date = e("#date-range").val();
        var sid = e("#sid").val();
        e.post('/monitor/online', {date: date, sid: sid}, function (res) {
            xData = res.xData;
            yData = res.yData;
        }, "json");

        var t = [{
            tooltip: {trigger: "axis"},
            dataZoom: {show: !0, realtime: !0, start: 0, end: 100},
            legend: {data: ["在线玩家数"]},
            xAxis: [{
                type: "category",
                boundaryGap: false,
                axisLine: {onZero: !1},
                data: xData
            }],
            yAxis: [{
                type: "value", name: "在线玩家数", axisLabel: {formatter: "{value}"}
            }],
            series: [{
                name: "玩家数",
                type: "line",
                data: yData
            }]
        }];
        var i = e("#LAY-index-pagetwo").children("div"), n = function (e) {
            l[e] = a.init(i[e], layui.echartsTheme), l[e].setOption(t[e]), window.onresize = l[e].resize
        };
        i[0] && n(0)
    }), e("sample", {})
});