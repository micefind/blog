<script setup lang="ts">
import "vue-cropper/dist/index.css";
import { VueCropper } from "vue-cropper";
import { ref, reactive } from "vue";
import type { UploadProps, UploadFile } from "element-plus";
import { defineEmits } from "vue";

const emit = defineEmits(["clickChild"]);

const cropper = ref();
const option = reactive({
  img: "",
  outputSize: 1,
  outputType: "jpeg",
  info: true,
  canScale: false,
  autoCrop: true,
  autoCropWidth: 150,
  autoCropHeight: 150,
  fixed: false,
  fixedNumber: [1, 1],
  canMove: true,
  canMoveBox: true,
  centerBox: true,
  full: true,
  enlarge: "1",
  mode: "contain",
});

const imgUrl = ref("");
const fileList = ref<UploadFile[]>([]);

// 获取裁剪后的图片
const getCropDataBase64 = () => {
  cropper.value.getCropData((data: string) => {
    imgUrl.value = data;
    emit("clickChild", imgUrl.value);
  });
};

// 处理上传的文件并生成裁剪图
const handleChange: UploadProps["onChange"] = (uploadFile: UploadFile) => {
  option.img = URL.createObjectURL(uploadFile.raw!);
};
</script>

<template>
  <div class="main-box">
    <el-row class="cropper-box">
      <VueCropper
        style="width: 100%; height: 100%"
        ref="cropper"
        v-bind="option"
      />
    </el-row>
    <el-row class="btn-box">
      <el-upload
        v-model:file-list="fileList"
        action=""
        :auto-upload="false"
        :show-file-list="false"
        :on-change="handleChange"
      >
        <el-button type="primary">点击上传图片</el-button>
      </el-upload>
      <el-button type="primary" @click="getCropDataBase64" style="margin-left: 10px">
        获取截取的图片
      </el-button>
    </el-row>
    <div class="pic-box" v-if="imgUrl">
      <div class="image-box">
        <img :src="imgUrl" class="img0" />
        <div class="img-box">
          <el-image :src="imgUrl" class="img1" fit="cover" />
          <el-image :src="imgUrl" class="img2" fit="cover" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.main-box {
  width: 100%;
  height: auto;
}

.cropper-box {
  height: 300px;
}

.btn-box {
  padding: 20px;
  display: flex;
  justify-content: center;
}

.pic-box {
  .image-box {
    display: flex;
    flex-wrap: wrap;

    .img0 {
      max-height: 150px;
    }

    .img-box {
      padding: 10px 20px;
      display: flex;
      align-items: center;

      .img1 {
        width: 100px;
        height: 100px;
        border-radius: 50%;
        margin-right: 20px;
      }

      .img2 {
        width: 50px;
        height: 50px;
        border-radius: 50%;
      }
    }
  }
}
</style>
