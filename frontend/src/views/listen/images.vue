<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.name" class="w220" clearable placeholder="图片名"/>
        </el-form-item>
        <el-form-item>
          <el-button @click="fetchData" type="primary" icon="el-icon-search">筛选</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-upload" size="mini" @click="uploadResource">上传图片</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loading" :data="assets.list" highlight-current-row stripe border size="mini" style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="name" label="图片">
          <template slot-scope="scope">
            <img :src="scope.row.f_code|lsImg" class="list-img"/>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="图片名"/>
        <el-table-column align="center" label="操作" fixed="right" width="100">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="danger" plain @click.native.prevent="delRow(scope.row.id)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="assets.total" @current-change="handlePage" @size-change="handlePageSize" />

      <image-upload ref="upload" @upload-success="fetchData"/>
    </el-col>
  </el-row>
</template>

<script>
  import ImageUpload from './components/upload'
  import { assetList, assetDelete,assetInfo } from '@a/asset'
  import Page from "@c/Page"

  export default {
    components: { ImageUpload,Page },
    data() {
      return {
        loading: false,
        search: {
          name: '',
          page: 1,
          page_size: 10
        },
        assets: {
          total: 0,
          list: []
        }
      }
    },
    filters: {
      lsImg(code) {
        return assetInfo+code
      }
    },
    mounted() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        this.loading = true
        assetList(this.search).then(res => {
          this.loading = false
          this.assets = res.data
        }).catch(err => {
          this.loading = false
        })
      },
      delRow(id) {
        this.$confirm('已设置的课程图不会丢失，新课程无法选此图片，确定删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          this.loading = true
          assetDelete(id).then(res => {
            this.loading = false
            this.$message.success("删除成功")
            this.fetchData()
          }).catch(err => {
            this.loading = false
          })
        }).catch(() => {
        })
      },
      uploadResource() {
        this.$refs.upload.initUpload()
      },
      handlePage(p) {
        this.search.page = p
        this.fetchData()
      },
      handlePageSize(p) {
        this.search.page_size = p
        this.fetchData()
      }
    }
  }
</script>

<style lang="scss" scoped>
.list-img {
  max-height: 50px;
  max-width: 150px;
}
</style>
