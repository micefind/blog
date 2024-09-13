<script setup lang="ts">
import { onShareAppMessage, onShareTimeline } from "@dcloudio/uni-app"
import { ref } from "vue"
import { onShow } from "@dcloudio/uni-app"
import avatar from "../../static/images/avatar.png"

const userInfo = ref<any>({
  nickName: "",
  avatarUrl: "",
  openid: "",
})

onShow(() => {
  userInfo.value.nickName = uni.getStorageSync("nickName")
  userInfo.value.avatarUrl = uni.getStorageSync("avatarUrl")
  userInfo.value.openid = uni.getStorageSync("openid")
})

const toolList = ref([
  {
    name: "个人信息",
    icon: "account",
    path: "/pages/qrcode/qrcode",
  },
  {
    name: "历史记录",
    icon: "file-text",
    path: "/pages/image/image",
  },
])

const gotoDetails = (item: any) => {
  switch (item.name) {
    case "个人信息":
      uni.showToast({
        title: "暂未开放",
        icon: "none",
      })
      break
    case "历史记录":
      uni.navigateTo({
        url: "/subpkgs/picList/picList",
      })
      break
  }
}

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
  <view class="container-box">
    <view class="header-box">
      <view class="user-img">
        <image :src="userInfo.avatarUrl || avatar" mode=""></image>
      </view>
      <view class="user-info">
        <view class="user-name"> {{ userInfo.nickName || "微信用户" }}</view>
        <view class="other-info">ID：{{ userInfo.openid }}</view>
      </view>
    </view>

    <view class="manage-box">
      <view class="title"> 常用工具 </view>
      <view class="tool-box">
        <view
          class="box-item"
          v-for="(item, index) in toolList"
          @click="gotoDetails(item)"
        >
          <view class="icon">
            <up-icon :name="item.icon" color="#303133" size="30"></up-icon>
          </view>
          <view>{{ item.name }}</view>
        </view>
      </view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.container-box {
  padding: 50rpx 20rpx;
}

.header-box {
  display: flex;
  align-items: center;
}

.user-img {
  height: 100rpx;
  width: 100rpx;
  border-radius: 50%;
}

.user-img image {
  height: 100%;
  width: 100%;
  border-radius: 50%;
}

.user-info {
  flex: 1;
  padding: 0 20rpx;
  display: flex;
  align-items: center;
  flex-direction: column;
  .user-name {
    width: 100%;
    margin-bottom: 10rpx;
    font-size: 32rpx;
    font-weight: 600;
  }
  .other-info {
    width: 100%;
    font-size: 24rpx;
  }
}

.manage-box {
  border-radius: 10rpx;
  background-color: #fff;
  padding: 30rpx;
  margin: 50rpx 0;
  .title {
    font-size: 32rpx;
    margin-bottom: 10rpx;
  }
  .tool-box {
    display: flex;
    flex-wrap: wrap;
    .box-item {
      width: 25%;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      padding: 20rpx;
      border-radius: 10rpx;
      .icon {
        padding: 20rpx;
      }
    }
  }
}
</style>
