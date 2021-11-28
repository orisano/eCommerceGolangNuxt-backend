var qa = {
                    id: "filler", afterDatasetsUpdate: function (t, e) {
                        var n, i, a, r, o = (t.data.datasets || []).length, s = e.propagate, l = [];
                        for (i = 0; i < o; ++i) r = null, (a = (n = t.getDatasetMeta(i)).dataset) && a._model && a instanceof Yt.Line && (r = {
                            visible: t.isDatasetVisible(i),
                            fill: Ra(a, i, o),
                            chart: t,
                            el: a
                        }), n.$filler = r, l.push(r);
                        for (i = 0; i < o; ++i) (r = l[i]) && (r.fill = Na(l, i, s), r.boundary = ja(r), r.mapper = La(r))
                    }, beforeDatasetsDraw: function (t) {
                        var e, n, i, a, r, o, s, l = t._getSortedVisibleDatasetMetas(), c = t.ctx;
                        for (n = l.length - 1; n >= 0; --n) (e = l[n].$filler) && e.visible && (a = (i = e.el)._view, r = i._children || [], o = e.mapper, s = a.backgroundColor || Q.global.defaultColor, o && s && r.length && (ot.canvas.clipArea(c, t.chartArea), Ba(c, r, o, a, s, i._loop), ot.canvas.unclipArea(c)))
                    }
                }, za = ot.rtl.getRtlAdapter, Wa = ot.noop, $a = ot.valueOrDefault;

                function Va(t, e) {
                    return t.usePointStyle && t.boxWidth > e ? e : t.boxWidth
                }

                Q._set("global", {
                    legend: {
                        display: !0,
                        position: "top",
                        align: "center",
                        fullWidth: !0,
                        reverse: !1,
                        weight: 1e3,
                        onClick: function (t, e) {
                            var n = e.datasetIndex, i = this.chart, a = i.getDatasetMeta(n);
                            a.hidden = null === a.hidden ? !i.data.datasets[n].hidden : null, i.update()
                        },
                        onHover: null,
                        onLeave: null,
                        labels: {
                            boxWidth: 40, padding: 10, generateLabels: function (t) {
                                var e = t.data.datasets, n = t.options.legend || {},
                                    i = n.labels && n.labels.usePointStyle;
                                return t._getSortedDatasetMetas().map((function (n) {
                                    var a = n.controller.getStyle(i ? 0 : void 0);
                                    return {
                                        text: e[n.index].label,
                                        fillStyle: a.backgroundColor,
                                        hidden: !t.isDatasetVisible(n.index),
                                        lineCap: a.borderCapStyle,
                                        lineDash: a.borderDash,
                                        lineDashOffset: a.borderDashOffset,
                                        lineJoin: a.borderJoinStyle,
                                        lineWidth: a.borderWidth,
                                        strokeStyle: a.borderColor,
                                        pointStyle: a.pointStyle,
                                        rotation: a.rotation,
                                        datasetIndex: n.index
                                    }
                                }), this)
                            }
                        }
                    }, legendCallback: function (t) {
                        var e, n, i, a = document.createElement("ul"), r = t.data.datasets;
                        for (a.setAttribute("class", t.id + "-legend"), e = 0, n = r.length; e < n; e++) (i = a.appendChild(document.createElement("li"))).appendChild(document.createElement("span")).style.backgroundColor = r[e].backgroundColor, r[e].label && i.appendChild(document.createTextNode(r[e].label));
                        return a.outerHTML
                    }
                });
                var Ya = ft.extend({
                    initialize: function (t) {
                        var e = this;
                        ot.extend(e, t), e.legendHitBoxes = [], e._hoveredItem = null, e.doughnutMode = !1
                    }, beforeUpdate: Wa, update: function (t, e, n) {
                        var i = this;
                        return i.beforeUpdate(), i.maxWidth = t, i.maxHeight = e, i.margins = n, i.beforeSetDimensions(), i.setDimensions(), i.afterSetDimensions(), i.beforeBuildLabels(), i.buildLabels(), i.afterBuildLabels(), i.beforeFit(), i.fit(), i.afterFit(), i.afterUpdate(), i.minSize
                    }, afterUpdate: Wa, beforeSetDimensions: Wa, setDimensions: function () {
                        var t = this;
                        t.isHorizontal() ? (t.width = t.maxWidth, t.left = 0, t.right = t.width) : (t.height = t.maxHeight, t.top = 0, t.bottom = t.height), t.paddingLeft = 0, t.paddingTop = 0, t.paddingRight = 0, t.paddingBottom = 0, t.minSize = {
                            width: 0,
                            height: 0
                        }
                    }, afterSetDimensions: Wa, beforeBuildLabels: Wa, buildLabels: function () {
                        var t = this, e = t.options.labels || {}, n = ot.callback(e.generateLabels, [t.chart], t) || [];
                        e.filter && (n = n.filter((function (n) {
                            return e.filter(n, t.chart.data)
                        }))), t.options.reverse && n.reverse(), t.legendItems = n
                    }, afterBuildLabels: Wa, beforeFit: Wa, fit: function () {
                        var t = this, e = t.options, n = e.labels, i = e.display, a = t.ctx,
                            r = ot.options._parseFont(n), o = r.size, s = t.legendHitBoxes = [], l = t.minSize,
                            c = t.isHorizontal();
                        if (c ? (l.width = t.maxWidth, l.height = i ? 10 : 0) : (l.width = i ? 10 : 0, l.height = t.maxHeight), i) {
                            if (a.font = r.string, c) {
                                var d = t.lineWidths = [0], u = 0;
                                a.textAlign = "left", a.textBaseline = "middle", ot.each(t.legendItems, (function (t, e) {
                                    var i = Va(n, o) + o / 2 + a.measureText(t.text).width;
                                    (0 === e || d[d.length - 1] + i + 2 * n.padding > l.width) && (u += o + n.padding, d[d.length - (e > 0 ? 0 : 1)] = 0), s[e] = {
                                        left: 0,
                                        top: 0,
                                        width: i,
                                        height: o
                                    }, d[d.length - 1] += i + n.padding
                                })), l.height += u
                            } else {
                                var h = n.padding, p = t.columnWidths = [], f = t.columnHeights = [], g = n.padding,
                                    m = 0, v = 0;
                                ot.each(t.legendItems, (function (t, e) {
                                    var i = Va(n, o) + o / 2 + a.measureText(t.text).width;
                                    e > 0 && v + o + 2 * h > l.height && (g += m + n.padding, p.push(m), f.push(v), m = 0, v = 0), m = Math.max(m, i), v += o + h, s[e] = {
                                        left: 0,
                                        top: 0,
                                        width: i,
                                        height: o
                                    }
                                })), g += m, p.push(m), f.push(v), l.width += g
                            }
                            t.width = l.width, t.height = l.height
                        } else t.width = l.width = t.height = l.height = 0
                    }, afterFit: Wa, isHorizontal: function () {
                        return "top" === this.options.position || "bottom" === this.options.position
                    }, draw: function () {
                        var t = this, e = t.options, n = e.labels, i = Q.global, a = i.defaultColor,
                            r = i.elements.line, o = t.height, s = t.columnHeights, l = t.width, c = t.lineWidths;
                        if (e.display) {
                            var d, u = za(e.rtl, t.left, t.minSize.width), h = t.ctx,
                                p = $a(n.fontColor, i.defaultFontColor), f = ot.options._parseFont(n), g = f.size;
                            h.textAlign = u.textAlign("left"), h.textBaseline = "middle", h.lineWidth = .5, h.strokeStyle = p, h.fillStyle = p, h.font = f.string;
                            var m = Va(n, g), v = t.legendHitBoxes, b = function (t, e, i) {
                                if (!(isNaN(m) || m <= 0)) {
                                    h.save();
                                    var o = $a(i.lineWidth, r.borderWidth);
                                    if (h.fillStyle = $a(i.fillStyle, a), h.lineCap = $a(i.lineCap, r.borderCapStyle), h.lineDashOffset = $a(i.lineDashOffset, r.borderDashOffset), h.lineJoin = $a(i.lineJoin, r.borderJoinStyle), h.lineWidth = o, h.strokeStyle = $a(i.strokeStyle, a), h.setLineDash && h.setLineDash($a(i.lineDash, r.borderDash)), n && n.usePointStyle) {
                                        var s = m * Math.SQRT2 / 2, l = u.xPlus(t, m / 2), c = e + g / 2;
                                        ot.canvas.drawPoint(h, i.pointStyle, s, l, c, i.rotation)
                                    } else h.fillRect(u.leftForLtr(t, m), e, m, g), 0 !== o && h.strokeRect(u.leftForLtr(t, m), e, m, g);
                                    h.restore()
                                }
                            }, y = function (t, e, n, i) {
                                var a = g / 2, r = u.xPlus(t, m + a), o = e + a;
                                h.fillText(n.text, r, o), n.hidden && (h.beginPath(), h.lineWidth = 2, h.moveTo(r, o), h.lineTo(u.xPlus(r, i), o), h.stroke())
                            }, x = function (t, i) {
                                switch (e.align) {
                                    case"start":
                                        return n.padding;
                                    case"end":
                                        return t - i;
                                    default:
                                        return (t - i + n.padding) / 2
                                }
                            }, _ = t.isHorizontal();
                            d = _ ? {x: t.left + x(l, c[0]), y: t.top + n.padding, line: 0} : {
                                x: t.left + n.padding,
                                y: t.top + x(o, s[0]),
                                line: 0
                            }, ot.rtl.overrideTextDirection(t.ctx, e.textDirection);
                            var w = g + n.padding;
                            ot.each(t.legendItems, (function (e, i) {
                                var a = h.measureText(e.text).width, r = m + g / 2 + a, p = d.x, f = d.y;
                                u.setWidth(t.minSize.width), _ ? i > 0 && p + r + n.padding > t.left + t.minSize.width && (f = d.y += w, d.line++, p = d.x = t.left + x(l, c[d.line])) : i > 0 && f + w > t.top + t.minSize.height && (p = d.x = p + t.columnWidths[d.line] + n.padding, d.line++, f = d.y = t.top + x(o, s[d.line]));
                                var S = u.x(p);
                                b(S, f, e), v[i].left = u.leftForLtr(S, v[i].width), v[i].top = f, y(S, f, e, a), _ ? d.x += r + n.padding : d.y += w
                            })), ot.rtl.restoreTextDirection(t.ctx, e.textDirection)
                        }
                    }, _getLegendItemAt: function (t, e) {
                        var n, i, a, r = this;
                        if (t >= r.left && t <= r.right && e >= r.top && e <= r.bottom) for (a = r.legendHitBoxes, n = 0; n < a.length; ++n) if (t >= (i = a[n]).left && t <= i.left + i.width && e >= i.top && e <= i.top + i.height) return r.legendItems[n];
                        return null
                    }, handleEvent: function (t) {
                        var e, n = this, i = n.options, a = "mouseup" === t.type ? "click" : t.type;
                        if ("mousemove" === a) {
                            if (!i.onHover && !i.onLeave) return
                        } else {
                            if ("click" !== a) return;
                            if (!i.onClick) return
                        }
                        e = n._getLegendItemAt(t.x, t.y), "click" === a ? e && i.onClick && i.onClick.call(n, t.native, e) : (i.onLeave && e !== n._hoveredItem && (n._hoveredItem && i.onLeave.call(n, t.native, n._hoveredItem), n._hoveredItem = e), i.onHover && e && i.onHover.call(n, t.native, e))
                    }
                });

                function Ua(t, e) {
                    var n = new Ya({ctx: t.ctx, options: e, chart: t});
                    $e.configure(t, n, e), $e.addBox(t, n), t.legend = n
                }

                var Ga = {
                    id: "legend", _element: Ya, beforeInit: function (t) {
                        var e = t.options.legend;
                        e && Ua(t, e)
                    }, beforeUpdate: function (t) {
                        var e = t.options.legend, n = t.legend;
                        e ? (ot.mergeIf(e, Q.global.legend), n ? ($e.configure(t, n, e), n.options = e) : Ua(t, e)) : n && ($e.removeBox(t, n), delete t.legend)
                    }, afterEvent: function (t, e) {
                        var n = t.legend;
                        n && n.handleEvent(e)
                    }
                }, Xa = ot.noop;
                Q._set("global", {
                    title: {
                        display: !1,
                        fontStyle: "bold",
                        fullWidth: !0,
                        padding: 10,
                        position: "top",
                        text: "",
                        weight: 2e3
                    }
                });
                var Za = ft.extend({
                    initialize: function (t) {
                        var e = this;
                        ot.extend(e, t), e.legendHitBoxes = []
                    },
                    beforeUpdate: Xa,
                    update: function (t, e, n) {
                        var i = this;
                        return i.beforeUpdate(), i.maxWidth = t, i.maxHeight = e, i.margins = n, i.beforeSetDimensions(), i.setDimensions(), i.afterSetDimensions(), i.beforeBuildLabels(), i.buildLabels(), i.afterBuildLabels(), i.beforeFit(), i.fit(), i.afterFit(), i.afterUpdate(), i.minSize
                    },
                    afterUpdate: Xa,
                    beforeSetDimensions: Xa,
                    setDimensions: function () {
                        var t = this;
                        t.isHorizontal() ? (t.width = t.maxWidth, t.left = 0, t.right = t.width) : (t.height = t.maxHeight, t.top = 0, t.bottom = t.height), t.paddingLeft = 0, t.paddingTop = 0, t.paddingRight = 0, t.paddingBottom = 0, t.minSize = {
                            width: 0,
                            height: 0
                        }
                    },
                    afterSetDimensions: Xa,
                    beforeBuildLabels: Xa,
                    buildLabels: Xa,
                    afterBuildLabels: Xa,
                    beforeFit: Xa,
                    fit: function () {
                        var t, e = this, n = e.options, i = e.minSize = {}, a = e.isHorizontal();
                        n.display ? (t = (ot.isArray(n.text) ? n.text.length : 1) * ot.options._parseFont(n).lineHeight + 2 * n.padding, e.width = i.width = a ? e.maxWidth : t, e.height = i.height = a ? t : e.maxHeight) : e.width = i.width = e.height = i.height = 0
                    },
                    afterFit: Xa,
                    isHorizontal: function () {
                        var t = this.options.position;
                        return "top" === t || "bottom" === t
                    },
                    draw: function () {
                        var t = this, e = t.ctx, n = t.options;
                        if (n.display) {
                            var i, a, r, o = ot.options._parseFont(n), s = o.lineHeight, l = s / 2 + n.padding, c = 0,
                                d = t.top, u = t.left, h = t.bottom, p = t.right;
                            e.fillStyle = ot.valueOrDefault(n.fontColor, Q.global.defaultFontColor), e.font = o.string, t.isHorizontal() ? (a = u + (p - u) / 2, r = d + l, i = p - u) : (a = "left" === n.position ? u + l : p - l, r = d + (h - d) / 2, i = h - d, c = Math.PI * ("left" === n.position ? -.5 : .5)), e.save(), e.translate(a, r), e.rotate(c), e.textAlign = "center", e.textBaseline = "middle";
                            var f = n.text;
                            if (ot.isArray(f)) for (var g = 0, m = 0; m < f.length; ++m) e.fillText(f[m], 0, g, i), g += s; else e.fillText(f, 0, 0, i);
                            e.restore()
                        }
                    }
                })
function Ka(t, e) {
                    var n = new Za({ctx: t.ctx, options: e, chart: t});
                    $e.configure(t, n, e), $e.addBox(t, n), t.titleBlock = n
                }

                var Qa = {}, Ja = qa, tr = Ga, er = {
                    id: "title", _element: Za, beforeInit: function (t) {
                        var e = t.options.title;
                        e && Ka(t, e)
                    }, beforeUpdate: function (t) {
                        var e = t.options.title, n = t.titleBlock;
                        e ? (ot.mergeIf(e, Q.global.title), n ? ($e.configure(t, n, e), n.options = e) : Ka(t, e)) : n && ($e.removeBox(t, n), delete t.titleBlock)
                    }
                };
for (var nr in Qa.filler = Ja, Qa.legend = tr, Qa.title = er, Yn.helpers = ot, Un(), Yn._adapters = Zn, Yn.Animation = mt, Yn.animationService = vt, Yn.controllers = Ce, Yn.DatasetController = St, Yn.defaults = Q, Yn.Element = ft, Yn.elements = Yt, Yn.Interaction = Me, Yn.layouts = $e, Yn.platform = yn, Yn.plugins = xn, Yn.Scale = vi, Yn.scaleService = _n, Yn.Ticks = Kn, Yn.Tooltip = Nn, Yn.helpers.each(Aa, (function (t, e) {
                    Yn.scaleService.registerScaleType(e, t, t._defaults)
                })), Qa) Qa.hasOwnProperty(nr) && Yn.plugins.register(Qa[nr]);
                Yn.platform.initialize();
document.addEventListener("DOMContentLoaded", (function () {
    if (document.getElementsByClassName("js-simplebar")[0]) {
        new nr(document.getElementsByClassName("js-simplebar")[0]);
        var t = document.getElementsByClassName("sidebar")[0];
        console.log(t)
        document.getElementsByClassName("sidebar-toggle")[0].addEventListener("click", (function () {
            t.classList.toggle("collapsed"), t.addEventListener("transitionend", (function () {
                window.dispatchEvent(new Event("resize"))
            }))
        }))
    }
}));
