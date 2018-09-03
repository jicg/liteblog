/**
 *
 */
window.wangEditor.fullscreen = {

    // editor create之后调用
    init: function (editorSelector) {
        var $ = layui.jquery;
        $(editorSelector + " .w-e-toolbar").append('<div class="w-e-menu"><a class="_wangEditor_btn_fullscreen" href="###" onclick="window.wangEditor.fullscreen.toggleFullscreen(\'' + editorSelector + '\')">全屏</a></div>');
    },
    toggleFullscreen: function (editorSelector) {
        var $ = layui.jquery;
        $(editorSelector).toggleClass('fullscreen-editor');
        if ($(editorSelector + ' ._wangEditor_btn_fullscreen').text() == '全屏') {
            $(editorSelector + ' ._wangEditor_btn_fullscreen').text('退出全屏');
        } else {
            $(editorSelector + ' ._wangEditor_btn_fullscreen').text('全屏');
        }
    }
};
