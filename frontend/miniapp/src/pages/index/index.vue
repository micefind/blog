<script setup lang="ts">
import { onShareAppMessage, onShareTimeline } from "@dcloudio/uni-app"
import { ref } from "vue"

const imageSrc = ref<string>("")

const login = () => {
  uni.login({
    provider: "weixin",
    success: (loginRes: any) => {
      const code: string = loginRes.code
      const appid: string = "wx2d5d3ee5a3499a68"
      const secret: string = "4824df0d389d3088aca2487fc8f26839"
      uni.request({
        url: "https://api.weixin.qq.com/sns/jscode2session",
        data: {
          appid: appid,
          secret: secret,
          js_code: code,
        },
        success: (res: any) => {
          console.log("登录成功", res)
          const openid: string = res.data.openid
          const session_key: string = res.data.session_key
          uni.setStorageSync("openid", openid)
          uni.setStorageSync("session_key", session_key)
        },
        fail: () => {},
      })
    },
  })
}
login()

onShareAppMessage(() => {
  return {
    title: "",
    path: "/pages/index/index",
    imageUrl: "",
  }
})
onShareTimeline(() => {
  return {
    title: "",
    path: "/pages/index/index",
    imageUrl: "",
  }
})
</script>

<template>
  <view> 首页 </view>
</template>

<style scoped lang="scss"></style>
