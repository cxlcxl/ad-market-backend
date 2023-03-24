<template>
  <dialog-panel title="用户修改" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save" :confirm-loading="loading" width="388px">
    <el-form :model="userForm" ref="userForm" label-width="90px" size="small" :rules="userRules">
      <el-form-item label="用户名称" prop="username">
        <el-input v-model="userForm.username" placeholder="请填写用户名" />
      </el-form-item>
      <el-form-item label="手机号" prop="mobile">
        <el-input v-model="userForm.mobile" placeholder="请填写手机号码" />
      </el-form-item>
      <el-form-item label="状态" prop="state">
        <el-switch v-model="userForm.state" :active-value="1" :inactive-value="0" />
      </el-form-item>
      <el-form-item label="登录密码" prop="pass">
        <el-input v-model="userForm.pass" placeholder="字母开头，数字特殊字符 [@.&!#?,%$] 的 6 - 18 位" />
      </el-form-item>
    </el-form>
  </dialog-panel>
</template>

<script>
import DialogPanel from "@c/DialogPanel"
import { getUserInfo, userUpdate } from "@a/user"
import { validPass, validMobile } from "@/utils/validate"

export default {
  components: {
    DialogPanel,
  },
  data() {
    var checkPass = (rule, value, callback) => {
      if (value === "") {
        callback()
      } else {
        if (!validPass(value)) {
          callback(new Error("密码格式不符合要求"))
        } else {
          callback()
        }
      }
    }
    var checkMobile = (rule, value, callback) => {
      if (!validMobile(value)) {
        callback(new Error("手机号格式不正确"))
      } else {
        callback()
      }
    }
    return {
      visible: false,
      loading: false,
      remoteLoading: false,
      userForm: {
        id: 0,
        username: "",
        mobile: "",
        state: 1,
        pass: "",
      },
      userRules: {
        username: { required: true, message: "请填写用户名称" },
        mobile: [{ required: true, message: "请填写手机号" }, { validator: checkMobile }],
        pass: { validator: checkPass },
      },
    }
  },
  methods: {
    initUpdate(user_id) {
      getUserInfo(user_id)
        .then((res) => {
          this.userForm = res.data
          this.visible = true
        })
        .catch(() => {})
    },
    cancel() {
      this.$refs.userForm.resetFields()
      this.visible = false
    },
    save() {
      this.$refs.userForm.validate((v) => {
        if (v) {
          this.loading = true
          userUpdate(this.userForm)
            .then((res) => {
              this.$message.success("修改成功")
              this.$emit("success")
              this.loading = false
              this.cancel()
            })
            .catch((err) => {
              this.loading = false
              console.log(err)
            })
        } else {
          return false
        }
      })
    },
  },
}
</script>
