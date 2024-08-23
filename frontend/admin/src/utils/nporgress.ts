import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

//全局进度条的配置
NProgress.configure({
  easing: 'ease',  // 动画方式，和css动画属性一样（默认：ease）
	speed: 500,  // 递增进度条的速度，单位ms（默认： 200）
	showSpinner: false, // 是否显示加载ico
	trickle: true,//是否自动递增
	trickleSpeed: 200, // 自动递增间隔
	minimum: 0.08, // 初始化时的最小百分比，0-1（默认：0.08）
  parent: 'body', //指定进度条的父容器
})

  // 打开进度条
export const start = () => {
    NProgress.start()
  }
   
  // 关闭进度条
  export const close = () => {
    NProgress.done()
  }