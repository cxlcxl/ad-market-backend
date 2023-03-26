<template>
  <dialog-panel title="课程图上传" :visible="visible" @cancel="cancel"
                confirm-text="确认上传" @confirm="upload" cancel-text="关闭"
                :confirm-loading="loading" width="700px">
    <el-form :model="assetForm" ref="assetForm" label-width="110px" size="mini">
      <el-form-item label="上传提示">
        <p>1. 以图片文件命名</p>
        <p>2. 图片过大可在 <a href="https://tinypng.com/" target="_blank" style="color: #1890ff;">此站</a> 压缩</p>
        <p class="text-error">3. 小程序主页课程列表图要求尺寸：500*300</p>
        <p>4. 其他图片尺寸按具体位置要求上传（配置列表）</p>
      </el-form-item>
      <el-form-item label="选择图片素材" prop="assets" :rules="{required: true}">
        <el-upload :action="uploadUrl" :headers="headers" multiple :limit="5" ref="assetUpload"
                   :on-exceed="overLimit" name="file"
                   :auto-upload="false" :file-list="fileList"
                   :data="assetForm" :on-change="handleChange" :on-remove="handleChange"
                   :accept="accepts"
                   :on-error="uploadErr"
                   :on-success="uploadSuccess">
          <el-button slot="trigger" type="primary" plain icon="el-icon-plus">选取文件（单图 100 KB内，一次性最大上传 5 张）</el-button>
        </el-upload>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
  import DialogPanel from '@c/DialogPanel'
  import { getToken } from '@/utils/auth'
  import {assetUpload} from '@a/asset'

  export default {
    components: {
      DialogPanel
    },
    data() {
      return {
        visible: false,
        loading: false,
        dimensionShow: false,
        assetForm: {},
        accepts: ".jpg,.jpeg,.png",
        headers: {
          "Authorization": 'Bearer ' + getToken()
        },
        fileList: [],
        assetNumbers: {
          total: 0,
          success: 0,
          failed: 0
        }
      }
    },
    computed: {
      uploadUrl() {
        return assetUpload
      }
    },
    methods: {
      initUpload() {
        this.visible = true
      },
      handleChange(file, list) {
        this.assetNumbers.total = list.length
      },
      upload() {
        if (this.assetNumbers.total === 0) {
          this.$message.error("请先选择素材")
          return
        }
        this.loading = true
        this.$refs.assetUpload.submit()
      },
      uploadErr(err, file, list) {
        this.assetNumbers.failed ++
        this.$notify.error({ title: '上传错误提示', message: `${file.name} 上传失败：${err}` , duration: 10000 })
        this.checkUploadComplete()
      },
      uploadSuccess(res, file, list) {
        this.assetNumbers.success ++
        this.checkUploadComplete()
      },
      checkUploadComplete() {
        if (this.assetNumbers.failed + this.assetNumbers.success >= this.assetNumbers.total) {
          this.loading = false
          if (this.assetNumbers.success > 0) {
            // this.$emit('upload-success')
          }
          if (this.assetNumbers.failed === 0) {
            // this.visible = false
          }
        }
      },
      cancel() {
        this.visible = false
        this.$emit('upload-success')
      },
      overLimit() {
        this.$message.error("最多一次选择 5 个素材")
      }
    }
  }
</script>

