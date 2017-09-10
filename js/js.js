$(function(){
    // input 表单获得焦点
    $('.search_text').focus(function(){
        $(this).attr('value', '');
        $(this).removeClass().addClasss('blue');
    })

    // input 表单失去焦点
    $('.search_text').blur(function(){
        $(this).removeClass('blue').addClasss('search_text');
    })

    // 右边头部隐藏盒子效果
    $('.user_r').hover(function(){
        $('.hidden_r').show();
    }, function(){
        $('.hidden_r').hide();
    })
})