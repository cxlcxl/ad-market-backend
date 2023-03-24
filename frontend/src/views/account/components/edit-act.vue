<template>
  <dialog-panel title="信息修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading">
    <el-form :model="accountForm" ref="accountForm" label-width="120px" size="small" :rules="userRules">
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="accountForm.mobile" placeholder="手机号" disabled/>
      </el-form-item>
      <el-form-item label="名称" prop="account_name">
        <el-input v-model="accountForm.account_name" placeholder="请填写账号名称" />
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-select v-model="accountForm.state" class="w120">
          <el-option v-for="(key, val) in AccountState" :label="key" :value="Number(val)" />
        </el-select>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="accountForm.remark" type="textarea"/>
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { accountInfo, accountUpdate } from "@a/account"

export default {
  components: {
    DialogPanel,
  },
  props: {
    AccountState: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      visible: false,
      loading: false,
      accountForm: {
        id: 0,
        account_name: "",
        mobile: "",
        remark: "",
        state: 1,
      },
      userRules: {
        account_name: { required: true, message: "请填写账户名称" },
        state: { required: true, message: "请选择" },
      },
    }
  },
  methods: {
    initUpdate(id) {
      accountInfo(id)
        .then((res) => {
          this.accountForm = res.data
          this.visible = true
        })
        .catch(() => {
          this.$message.error("用户信息请求错误")
        })
    },
    cancel() {
      this.$refs.accountForm.resetFields()
      this.visible = false
    },
    save() {
      this.$refs.accountForm.validate((v) => {
        if (v) {
          this.loading = true
          accountUpdate(this.accountForm)
            .then((res) => {
              this.$message.success("修改成功")
              this.$emit("success")
              this.loading = false
              this.cancel()
            })
            .catch((err) => {
              this.loading = false
            })
        } else {
          return false
        }
      })
    },
  },
}
</script>
