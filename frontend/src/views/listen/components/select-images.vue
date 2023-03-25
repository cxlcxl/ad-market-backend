<template>
  <page-drawer :visible="visible" title="图片选择" @cancel="cancel">
    <el-row style="padding: 0 15px;">
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
      <el-col :span="24" class="select-box">
        <el-table v-loading="loading" :data="assets.list" highlight-current-row stripe border size="mini">
          <el-table-column prop="name" label="图片" width="170" align="center">
            <template slot-scope="scope">
              <img :src="scope.row.f_code|lsImg" class="list-img"/>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="图片名"/>
          <el-table-column align="center" label="操作" fixed="right" width="70">
          <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-thumb" circle plain @click="handleSelect(scope.row.f_code)"/>
          </template>
        </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="24" style="margin: 15px 0; text-align: center;">
        <el-button-group>
          <el-button type="primary" icon="el-icon-arrow-left" @click="handlePage(-1)"
                     :disabled="search.page === 1" plain>上一页</el-button>
          <el-button disabled plain>{{search.page}} / {{max_page}}</el-button>
          <el-button type="primary" @click="handlePage(1)" :disabled="search.page === max_page" plain>
            下一页<i class="el-icon-arrow-right el-icon--right"></i>
          </el-button>
        </el-button-group>
      </el-col>
    </el-row>
  </page-drawer>
</template>

<script>
  import { assetList, assetInfo } from '@a/asset'
  import PageDrawer from "@c/Drawer"

  export default {
    components: { PageDrawer },
    data() {
      return {
        loading: false,
        visible: false,
        search: {
          name: '',
          page: 1,
          page_size: 6
        },
        assets: {
          total: 0,
          list: []
        },
        max_page: 1
      }
    },
    filters: {
      lsImg(code) {
        return assetInfo+code
      }
    },
    methods: {
      initSelect() {
        this.search.page = 1
        assetList(this.search).then(res => {
          this.visible = true
          this.assets = res.data
          this.max_page = Math.ceil(this.assets.total/this.search.page_size)
          console.log(this.max_page)
        }).catch(err => {})
      },
      fetchData() {
        this.loading = true
        assetList(this.search).then(res => {
          this.loading = false
          this.assets = res.data
        }).catch(err => {
          this.loading = false
        })
      },
      handlePage(v) {
        if (this.search.page <= 1 && v === -1) {
          return
        }
        if (v === 1 && this.max_page <= this.search.page) {
          return
        }
        this.search.page += v
        this.fetchData()
      },
      handleSelect(v) {
        this.$emit('handle-select', v)
        this.cancel()
      },
      cancel() {
        this.visible = false
      }
    }
  }
</script>

<style lang="scss">
.select-box .list-img {
  max-height: 70px;
  max-width: 150px;
}
</style>
