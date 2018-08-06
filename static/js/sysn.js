/**

 @Name：layui.sysn
 */
layui.define(['jquery', 'layer'], function (exports) {
    var $ = layui.jquery,
        layer = layui.layer,
        ajaxObj = {async: true},
        sysn = {};

    function ajax(url, type, timeout, data, async, success, error, complete) {
        var error2 = function (ret) {
            var msg = ret.responseText || ret.msg || ret;
            if (ret.status != undefined && ret.status == 0) {
                msg = "网络异常";
            }
            if (!error) {
                sysn.sayError(msg);
            } else {
                error(ret)
            }
        };
        $.ajax({
            url: url,
            type: type,
            timeout: timeout || 5000,
            data: data,
            async: async,
            success: function (ret) {
                if (ret.code == 0) {
                    success(ret)
                } else {
                    error2(ret)
                }
            },
            error: error2,
            complete: complete
        });
    }

    function ajax2(url, type, timeout, data, async, success, error, complete) {
        var error2 = function (ret) {
            var msg = ret.responseText || ret.msg || ret;
            if (ret.status != undefined && ret.status == 0) {
                msg = "网络异常";
            }
            sysn.sayError(msg);
        };
        $.ajax({
            url: url,
            type: type,
            timeout: timeout || 5000,
            data: data,
            async: async,
            success: function (ret) {
                success(ret)
            },
            error: error2,
            complete: complete
        });
    }


    sysn.get = function (url, data) {
        ajaxObj.url = url;
        ajaxObj.data = data;
        ajaxObj.method = "GET";
        return this;
    };


    sysn.post = function (url, data) {
        ajaxObj.url = url;
        ajaxObj.data = data;
        ajaxObj.method = "POST";
        return this;
    };

    sysn.success = function (success) {
        ajaxObj.success = success;
        return this;
    };

    sysn.success = function (success) {
        ajaxObj.success = success;
        return this;
    };

    sysn.async = function (async) {
        ajaxObj.async = async;
        return this;
    };

    sysn.error = function (error) {
        ajaxObj.error = error;
        return this;
    };

    sysn.complete = function (complete) {
        ajaxObj.complete = complete;
        return this;
    };

    sysn.run = function (option) {
        if (ajaxObj.async == null || ajaxObj.async == undefined) {
            ajaxObj.async = true;
        }
        if (option && option.novaild) {
            ajax2(ajaxObj.url, ajaxObj.method, ajaxObj.timeout, ajaxObj.data, ajaxObj.async, ajaxObj.success, ajaxObj.error, ajaxObj.complete);

        } else {
            ajax(ajaxObj.url, ajaxObj.method, ajaxObj.timeout, ajaxObj.data, ajaxObj.async, ajaxObj.success, ajaxObj.error, ajaxObj.complete);
        }
    };


    sysn.sayOk = function (msg) {
        layer.msg(msg, {icon: 6});
    };
    sysn.sayError = function (msg) {
        layer.msg(msg, {icon: 5})
    };

    sysn.setTimeout = function (timeout) {
        ajaxObj.timeout = timeout;
        return this;
    };

    //输出test接口
    exports('sysn', sysn);
});  