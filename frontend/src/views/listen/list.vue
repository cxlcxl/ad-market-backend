<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.title" class="w150" clearable placeholder="课程名" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" class="w120" placeholder="是否展示">
            <el-option :value="1" label="展示"/>
            <el-option :value="0" label="不展示"/>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加课程</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="listenList.list" highlight-current-row stripe border size="mini" style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="title" label="课程标题"/>
        <el-table-column prop="title" label="小标题"/>
        <el-table-column prop="order_by" label="排序" width="90"/>
        <el-table-column prop="state" label="是否展示" width="90">
          <template slot-scope="scope">
            {{ scope.row.state === 1 ? '展示' : '不展示' }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" width="100">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="listenList.total" @current-change="handlePage" @size-change="handlePageSize" />

      <listen-create ref="listenCreate" @success="getList" />
      <listen-update ref="listenUpdate" @success="getList" />
    </el-col>
  </el-row>
</template>

<script>
import { listenList } from "@a/listen"
import ListenCreate from "./components/add-listen"
import ListenUpdate from "./components/edit-listen"
import Page from "@c/Page"

export default {
  // name: 'ListenList',
  components: {
    ListenCreate,
    ListenUpdate,
    Page,
  },
  data() {
    return {
      loadings: {
        pageLoading: false,
      },
      listenList: {
        total: 0,
        list: [],
      },
      search: {
        title: "",
        state: 1,
        page: 1,
        page_size: 10,
      },
    }
  },
  mounted() {
    this.getList()
  },
  methods: {
    getList() {
      this.loadings.pageLoading = true
      listenList(this.search)
        .then((res) => {
          this.listenList = res.data
          this.loadings.pageLoading = false
        })
        .catch(() => {
          this.loadings.pageLoading = false
        })
    },
    add() {
      this.$refs.listenCreate.initCreate()
    },
    editRow(row) {
      this.$refs.listenUpdate.initUpdate(row.id)
    },
    doSearch() {
      this.search.page = 1
      this.getList()
    },
    handlePage(p) {
      this.search.page = p
      this.getList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getList()
    },
  },
}
</script>
