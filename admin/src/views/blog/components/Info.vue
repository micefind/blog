<script setup>
import { ref } from 'vue'
import github from '@/assets/image/github.png'
import gitee from '@/assets/image/gitee.png'
import bilibili from '@/assets/image/bilibili.png'
import museum from '@/assets/image/museum.png'
import avatar from '@/assets/image/avatar.png'

const year = ref(new Date().getFullYear())

const objectList = ref([
  {
    name: '线上博物馆',
    intro: '博物馆微信小程序，在线展览历史文物',
    link: 'https://gitee.com/micefind/museum',
    img: museum,
  },
])
</script>

<template>
  <div class="info-page">
    <div class="header">
      <div class="avatar">
        <div class="image">
          <img :src="avatar" alt="" />
        </div>
      </div>
      <div class="name">micefind</div>
      <div class="link">
        <el-link :underline="false" target="_blank" href="https://github.com/micefind">
          <img :src="github" alt="" />
        </el-link>
        <el-link :underline="false" target="_blank" href="https://gitee.com/micefind">
          <img :src="gitee" alt="" />
        </el-link>
        <el-link :underline="false" target="_blank" href="https://space.bilibili.com/321496512">
          <img :src="bilibili" alt="" />
        </el-link>
      </div>
    </div>
    <div class="content">
      <el-scrollbar class="object">
        <el-link
          :underline="false"
          target="_blank"
          :href="item.link"
          v-for="item in objectList"
          :key="item.link"
        >
          <div class="object-item">
            <img :src="item.img" alt="" />
            <div class="detail">
              <div class="title">{{ item.name }}</div>
              <div class="desc">{{ item.intro }}</div>
            </div>
          </div>
        </el-link>
      </el-scrollbar>

      <div class="footer">
        <el-link :underline="false" type="info">© {{ year }} MiceFind.com</el-link>
        <el-link :underline="false" type="info" target="_blank" href="https://beian.miit.gov.cn/"
          >渝ICP备2022007989号</el-link
        >
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.info-page {
  height: 100%;
  color: var(--el-text-color-primary);
  .header {
    padding: 20px 20px 0;
    border-bottom: 1px solid var(--el-border-color);

    .avatar {
      display: flex;
      justify-content: center;
      .image {
        width: 100px;
        height: 100px;
        border-radius: 50%;
        overflow: hidden;

        &:hover {
          img {
            transform: rotateY(180deg);
          }
        }
        img {
          width: 100px;
          height: 100px;
          border-radius: 50%;
          /* 开启3D环境 */
          transform-style: preserve-3d;
          transition: transform 0.6s ease;
          border: 3px solid var(--el-border-color-lighter);
        }
      }
    }
    .name {
      padding: 10px 0;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 24px;
      font-weight: 700;
    }
    .link {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 10px;
      padding: 0 0 20px;

      img {
        width: 20px;
        height: 20px;
      }
    }
  }
  .content {
    height: calc(100% - 216px);

    .object {
      height: calc(100% - 80px);
      padding: 10px 20px;
      border-bottom: 1px solid var(--el-border-color);

      .object-item {
        padding: 10px 10px;
        display: flex;
        align-items: center;
        gap: 20px;

        img {
          width: 40px;
          height: 40px;
          object-fit: cover;
          border-radius: 5px;
          border: 1px solid var(--el-border-color);
        }
        .detail {
          height: 50px;
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: space-between;
          overflow: hidden;

          .title {
            font-size: 16px;
            font-weight: 700;
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
          }
          .desc {
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
          }
        }
      }
    }
    .footer {
      height: 80px;
      padding: 0 20px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      gap: 5px;

      .el-link {
        font-size: 13px;
        display: block;
      }
    }
  }
}
</style>
