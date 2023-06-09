<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.username" class="w150" clearable placeholder="用户名" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="search.mobile" class="w230" clearable placeholder="手机号" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" placeholder="用户状态" class="w110">
            <el-option :key="1" label="正常" :value="1" />
            <el-option :key="0" label="停用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="add">添加用户</el-button>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="userList.list" highlight-current-row stripe border size="mini" style="margin-top: 15px">
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="username" label="用户名"/>
        <el-table-column prop="mobile" label="手机号" />
        <el-table-column label="状态" width="90" align="center">
          <template slot-scope="scope">
            {{ scope.row.state === 1 ? '正常' : '停用' }}
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{scope.row.created_at|timeFormat}}</template>
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
      <page ref="page" :page="search.page" :total="userList.total" @current-change="handlePage" @size-change="handlePageSize" />

      <user-create ref="userCreate" @success="getUserList" />
      <user-update ref="userUpdate" @success="getUserList" />
    </el-col>
  </el-row>
</template>

<script>
import { list } from "@a/user"
import UserCreate from "./components/add-user"
import UserUpdate from "./components/edit-user"
import Page from "@c/Page"
import { parseTime } from "@/utils"

export default {
  // name: 'UserList',
  components: {
    UserCreate,
    UserUpdate,
    Page,
  },
  data() {
    return {
      loadings: {
        pageLoading: false,
      },
      userList: {
        total: 0,
        list: [],
      },
      search: {
        mobile: "",
        username: "",
        state: 1,
        page: 1,
        page_size: 10,
      },
    }
  },
  computed: {},
  filters: {
    timeFormat(t) {
      return parseTime(t)
    },
  },
  mounted() {
    this.getUserList()
  },
  methods: {
    getUserList() {
      this.loadings.pageLoading = true
      list(this.search)
        .then((res) => {
          this.userList = res.data
          this.loadings.pageLoading = false
        })
        .catch(() => {
          this.loadings.pageLoading = false
        })
    },
    add() {
      this.$refs.userCreate.visible = true
    },
    editRow(row) {
      this.$refs.userUpdate.initUpdate(row.id)
    },
    doSearch() {
      this.search.page = 1
      this.getUserList()
    },
    handlePage(p) {
      this.search.page = p
      this.getUserList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getUserList()
    },
  },
}
</script>
