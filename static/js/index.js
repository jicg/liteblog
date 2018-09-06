layui.use(['jquery', 'util'], function () {
    var util = layui.util,
        $ = layui.jquery;
    var option = {
        bar1: '&#xe770;',
        click: function (type) {
            if (type === 'bar1') {
                window.location.href = "/user";
            }
        }
    };
    if (user && user.id >= 0 && user.role === 0) {
        option = {
            bar1: '&#xe654;',
            click: function (type) {
                console.log(type);
                if (type === 'bar1') {
                    window.location.href = "/note/new";
                }
            }
        };
    }
    util.fixbar(option);
});