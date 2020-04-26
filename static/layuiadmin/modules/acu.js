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
        e.post('/monitor/acu', {date: date, sid: sid}, function (res) {
            xData = res.xData;
            yDataAcu = res.yDataAcu;
            yDataPcu = res.yDataPcu;
            yDataAcuPcu = res.yDataAcuPcu;
        }, "json");

        var tp = [{
            tooltip: {trigger: "axis"},
            dataZoom: {show: !0, realtime: !0, start: 0, end: 100},
            legend: {data: ["最高同时在线人数", "平均同时在线人数"]},
            xAxis: [{
                type: "category",
                boundaryGap: false,
                axisLine: {onZero: !1},
                data: xData
            }],
            yAxis: [
                {
                    type: "value", name: "ACU、PCU", axisLabel: {formatter: "{value}"}
                }
            ],
            series: [
                {
                    name: "最高同时在线人数",
                    type: "line",
                    data: yDataPcu
                },
                {
                    name: "平均同时在线人数",
                    type: "line",
                    data: yDataAcu
                }
            ]
        }];

        var tap = [{
            tooltip: {trigger: "axis"},
            dataZoom: {show: !0, realtime: !0, start: 0, end: 100},
            legend: {data: ["ACU/PCU"]},
            xAxis: [{
                type: "category",
                boundaryGap: false,
                axisLine: {onZero: !1},
                data: xData
            }],
            yAxis: [{
                type: "value", name: "ACU/PCU", axisLabel: {formatter: "{value}"}
            }],
            series: [{
                name: "ACU/PCU",
                type: "line",
                data: yDataAcuPcu
            }]
        }];

        var ip = e("#LAY-index-pcu").children("div"), np = function (e) {
            l[e] = a.init(ip[e], layui.echartsTheme), l[e].setOption(tp[e]), window.onresize = l[e].resize
        };
        ip[0] && np(0);

        var iap = e("#LAY-index-acu-pcu").children("div"), nap = function (e) {
            l[e] = a.init(iap[e], layui.echartsTheme), l[e].setOption(tap[e]), window.onresize = l[e].resize
        };
        iap[0] && nap(0)


    }), e("acu", {})
});