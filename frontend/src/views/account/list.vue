<template>
  <el-row>
    <el-col :span="24" class="search-container">
      <el-form ref="_search" :model="search" inline size="small">
        <el-form-item>
          <el-input v-model="search.mobile" class="w200" clearable placeholder="手机号" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="search.state" class="w120">
            <el-option label="全部" :value="0" />
            <el-option v-for="(key, val) in accountList.state" :label="key" :value="Number(val)" />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" icon="el-icon-search" class="item" @click="doSearch">查询</el-button>
          <el-button icon="el-icon-download" class="item" @click="download" :loading="loadings.downloadLoading">导出</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-table v-loading="loadings.pageLoading" :data="accountList.list" highlight-current-row stripe border size="mini">
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column label="手机号" prop="mobile" width="130" />
        <el-table-column prop="account_name" label="名称" width="200"/>
        <el-table-column label="状态" width="130" align="center">
          <template slot-scope="scope">{{scope.row.state|stateFilter(accountList.state)}}</template>
        </el-table-column>
        <el-table-column prop="remark" label="备注"/>
        <el-table-column prop="created_at" label="添加时间" width="140" align="center">
          <template slot-scope="scope">{{scope.row.created_at|timeFormat}}</template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="130">
          <template slot-scope="scope">
            <el-button-group class="table-operate">
              <el-button type="primary" plain @click.native.prevent="editRow(scope.row.id)">编辑</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
    <el-col :span="24" style="margin-top: 15px;">
      <page ref="page" :page="search.page" :total="accountList.total" @current-change="handlePage" @size-change="handlePageSize" />

      <account-update ref="accountUpdate" @success="getAccountList" :account-state="accountList.state" />
    </el-col>
  </el-row>
</template>

<script>
import { accountList, accountAuth, refreshAuth } from "@a/account"
import AccountUpdate from "./components/edit-act"
import Page from "@c/Page"
import { parseTime } from "@/utils"

export default {
  // name: "AccountList",
  components: {
    AccountUpdate,
    Page,
  },
  data() {
    return {
      loadings: {
        pageLoading: false,
        downloadLoading: false
      },
      accountList: {
        total: 0,
        list: [],
        state: {},
      },
      roles: [],
      search: {
        mobile: "",
        state: 0,
        page: 1,
        page_size: 10,
      },
    }
  },
  mounted() {
    this.getAccountList()
  },
  filters: {
    stateFilter(s, state) {
      return state[s]
    },
    timeFormat(timestamp) {
      return parseTime(timestamp)
    },
  },
  methods: {
    getAccountList() {
      this.loadings.pageLoading = true
      Promise.all([])
        .then((response) => {
          accountList(this.search)
            .then((res) => {
              const { list, total, state } = res.data
              this.accountList.list = list
              this.accountList.total = total
              this.accountList.state = state
              this.loadings.pageLoading = false
            })
            .catch(() => {
              this.loadings.pageLoading = false
            })
        })
        .catch(() => {
          this.loadings.pageLoading = false
        })
    },
    add() {
      this.$refs.accountCreate.initCreate()
    },
    editRow(id) {
      this.$refs.accountUpdate.initUpdate(id)
    },
    doSearch() {
      this.search.page = 1
      this.getAccountList()
    },
    handlePage(p) {
      this.search.page = p
      this.getAccountList()
    },
    handlePageSize(p) {
      this.search.page_size = p
      this.getAccountList()
    },
    download() {
      if (this.accountList.total === 0) {
        this.$message.info('没有需导出的数据')
        return
      }
      this.loadings.downloadLoading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = ['mobile', 'account_name', 'state', 'remark', 'created_at']
        const data = excel.formatJson(tHeader, this.accountList)
        if (data.length === 0) {
          this.$message.info('没有筛选到需导出的数据')
          return
        }
        excel.export_json_to_excel({ header: tHeader, data, filename: '商机' })
        this.loadings.downloadLoading = false
      }).catch(err => {
        console.log(err)
        this.loadings.downloadLoading = false
      })
    },
  },
}
</script>
