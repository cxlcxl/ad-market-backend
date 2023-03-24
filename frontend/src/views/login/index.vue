<template>
  <div class="login-container">
    <el-row style="height: 100%;">
      <el-col :span="24" :style="{height: '100%', position: 'relative'}">
        <el-form ref="login_form" :model="loginForm" :rules="loginRules" class="login-form" size="small">
          <div class="title-container">
            <h3 class="title">{{ SiteTitle }} Login</h3>
          </div>
          <el-form-item prop="mobile">
            <el-input v-model="loginForm.mobile" placeholder="请输入手机号" />
          </el-form-item>

          <el-form-item prop="pass">
            <el-input :key="passwordType" v-model="loginForm.pass" :type="passwordType" placeholder="请输入密码" @keyup.native="checkCapslock"
              @blur="capsTooltip = false" @keyup.enter.native="handleLogin" />
          </el-form-item>

          <el-form-item>
            <div style="text-align: center;">
              <el-button :loading="loading" type="primary" size="medium" @click.native.prevent="handleLogin">Login
              </el-button>
            </div>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { validMobile } from "@/utils/validate"
import Banner from "./components/banner"
import SocialSign from "./components/SocialSignin"
import Settings from "@/settings"

export default {
  name: "Login",
  components: { SocialSign, Banner },
  data() {
    const validateUsername = (rule, value, callback) => {
      if (!validMobile(value)) {
        callback(new Error("请填写正确的手机号"))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error("密码不可少于 6 个字符"))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        mobile: "",
        pass: "",
      },
      loginRules: {
        mobile: [{ required: true, trigger: "blur", validator: validateUsername }],
        pass: [{ required: true, trigger: "blur", validator: validatePassword }],
      },
      passwordType: "password",
      capsTooltip: false,
      loading: false,
      redirect: undefined,
      otherQuery: {},
    }
  },
  computed: {
    SiteTitle() {
      return Settings.title
    },
  },
  watch: {
    $route: {
      handler: function (route) {
        const query = route.query
        if (query) {
          this.redirect = query.redirect
          this.otherQuery = this.getOtherQuery(query)
        }
      },
      immediate: true,
    },
  },
  mounted() {
    if (this.loginForm.username === "") {
      this.$refs.username.focus()
    } else if (this.loginForm.password === "") {
      this.$refs.password.focus()
    }
  },
  methods: {
    checkCapslock(e) {
      const { key } = e
      this.capsTooltip = key && key.length === 1 && key >= "A" && key <= "Z"
    },
    showPwd() {
      if (this.passwordType === "password") {
        this.passwordType = ""
      } else {
        this.passwordType = "password"
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.login_form.validate((valid) => {
        if (valid) {
          this.loading = true
          this.$store
            .dispatch("user/login", this.loginForm)
            .then(() => {
              this.$router.push({
                path: this.redirect || "/",
                query: this.otherQuery,
              })
              this.loading = false
            })
            .catch((err) => {
              this.loading = false
            })
        } else {
          console.log("error submit!!")
          return false
        }
      })
    },
    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== "redirect") {
          acc[cur] = query[cur]
        }
        return acc
      }, {})
    },
  },
}
</script>

<style lang="scss" scoped>
.login-container {
  height: 100%;
  width: 100%;
  background: url("/images/login-bg.jpg");
  overflow: hidden;
  background-size: cover;

  .login-form {
    position: relative;
    top: 50%;
    left: 60%;
    width: 350px;
    max-width: 100%;
    padding: 20px 35px 0;
    margin-top: -180px;
    overflow: hidden;
    background: #fff;
    -webkit-border-radius: 5px;
    -moz-border-radius: 5px;
    border-radius: 5px;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      margin: 0 auto 20px auto;
      text-align: center;
      font-weight: bold;
    }
  }
}
</style>
