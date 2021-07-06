<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-title">管理员登录</div>
      <!-- 登录表单 -->
      <el-form
        :model="loginForm"
        :rules="rules"
        class="login-form"
        ref="ruleForm"
        status-icon
      >
        <!-- 用户名输入框 -->
        <el-form-item prop="username">
          <el-input
            @keyup.enter.native="login"
            placeholder="用户名"
            prefix-icon="el-icon-user-solid"
            v-model="loginForm.username"
          />
        </el-form-item>
        <!-- 密码输入框 -->
        <el-form-item prop="password">
          <el-input
            @keyup.enter.native="login"
            placeholder="密码"
            prefix-icon="iconfont el-icon-mymima"
            show-password
            v-model="loginForm.password"
          />
        </el-form-item>
      </el-form>
      <!-- 登录按钮 -->
      <el-button type="primary" @click="login">登录</el-button>
    </div>
  </div>
</template>

<script>
import { generaMenu } from "../../assets/js/menu";
import {adminLogin, getUserMenu} from "../../api/api";
import {getResultCode} from "../../utils/util";
import {resultMap} from "../../utils/constant";

export default {
  data: function() {
    return {
      loginForm: {
        username: "",
        password: ""
      },
      rules: {
        username: [
          { required: true, message: "用户名不能为空", trigger: "blur" }
        ],
        password: [{ required: true, message: "密码不能为空", trigger: "blur" }]
      }
    };
  },
  methods: {
    login() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          const that = this;
          // eslint-disable-next-line no-undef
          const captcha = new TencentCaptcha(this.config.TENCENT_CAPTCHA, function(res) {
              if (res.ret === 0) {
                //发送登录请求
                adminLogin({
                  user: {
                    username: that.loginForm.username,
                    password: that.loginForm.password
                  }
                }).then(( data ) => {
                  console.log(data)
                  if (data.code === getResultCode(resultMap.SuccessOK)) {
                    // 登录后保存用户信息
                    that.$store.commit("login", data.loginResponse);
                    // 加载用户菜单
                    generaMenu();
                    that.$message.success("登录成功");
                    that.$router.push({ path: "/" });
                  } else {
                    that.$message.error(data.message);
                  }
                });
              }
            }
          );
          // 显示验证码
          captcha.show();
        } else {
          return false;
        }
      });
    }
  }
};
</script>

<style scoped>
.login-container {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  background: url(https://www.static.talkxj.com/0w3pdr.jpg) center center /
    cover no-repeat;
}

.login-card {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  background: #fff;
  padding: 170px 60px 180px;
  width: 350px;
}

.login-title {
  color: #303133;
  font-weight: bold;
  font-size: 1rem;
}

.login-form {
  margin-top: 1.2rem;
}

.login-card button {
  margin-top: 1rem;
  width: 100%;
}
</style>
