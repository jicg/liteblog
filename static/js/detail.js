layui.use(['form', 'jquery', 'layer', 'sysn', 'util'], function () {
    var form = layui.form,
        sysn = layui.sysn,
        $ = layui.jquery,
        layer = layui.layer,
        util = layui.util;
    //执行
    var option = {
        bar1: '&#xe770;',
        click: function (type) {
            if (type === 'bar1') {
                window.location.href = "/user";
            }
        }
    };
    if (user.id > 0 && note.userid === user.id) {
        option = {
            bar1: '&#xe642;',
            bar2: '&#xe640;',
            click: function (type) {
                if (type === 'bar1') {
                    window.location.href = "/note/edit/"+note.key;
                } else if (type === 'bar2') {
                    layer.confirm('您确定要删除当前文章？', {
                        btn: ['确定', '取消'] //按钮
                    }, function () {
                        sysn.post("/note/del/"+note.key)
                            .success(function (data) {
                                layer.msg(data.msg);
                                if (data.action) {
                                    setTimeout(function () {
                                        window.location.href = data.action;
                                    }, 300)
                                }
                            }).run();
                    }, function () {
                    });

                }
            }

        };
    }
    util.fixbar(option);
    var fielstemp = $("#files").text(), filelists = $("#demoList");
    if (fielstemp) {
        var otherfiles = fielstemp.split(";");
        for (var i = 0; i < otherfiles.length; i++) {
            var index = i;
            sysn.get(otherfiles[index] + ".json")
                .async(false)
                .success(function (file) {
                    var tr = $(['<tr id="upload-' + index + '">'
                        , '<td><a  href="' + otherfiles[index] + '" target="_blank">' + file.Filename + '</a></td>'
                        , '</tr>'].join(''));
                    tr.addClass("done").data("link", otherfiles[index]);
                    filelists.append(tr);
                }).run({novaild: true});
        }
    }
});