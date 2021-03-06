<template>
  <div>
    <div class="top-title" @click="linkTo('', 'home')">上弦月</div>

    <div class="top-links">
      <div
        v-show="$store.state.isLogin"
        class="tl-item"
        @click="linkTo('files', 'files')"
      >
        云文件
      </div>

      <div class="tl-item">
        <span
          rel="bookmark"
          title="https://github.com/mats9693/unnamed_plan"
        >
          <a href="https://github.com/mats9693/unnamed_plan" target="_blank">本站代码</a>
        </span>
      </div>
    </div>

    <div class="top-login-entrance">
      <div v-if="!$store.state.isLogin" class="tle-text" @click="openLoginDialog">登录</div>

      <el-dropdown v-if="$store.state.isLogin" class="tle-text">
        <span>{{ $store.state.nickname }}<i class="el-icon-arrow-down el-icon--right"/></span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item @click.native="exit">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <el-dialog
      class="top-login-dialog"
      :visible.sync="loginDialogController"
      append-to-body
      :modal-append-to-body="false"
    >
      <div slot="title" class="tld-title">登录</div>

      <div class="tld-content">
        <div class="tldc-item">
          <span class="tldci-label">用户名&#58;</span>
          <el-input v-model="userName" placeholder="请输入用户名"/>
        </div>

        <div class="tldc-item">
          <span class="tldci-label">密码&#58;</span>
          <el-input v-model="password" type="password" placeholder="请输入密码" clearable/>
        </div>
      </div>

      <div slot="footer">
        <el-button @click="cancelLogin">取消</el-button>
        <el-button type="info" @click="login">登录</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import axios from "axios";
import { calcSHA256 } from "@/ts/sha256";

@Component
export default class Top extends Vue {
  private userName = "";
  private password = "";

  private loginDialogController = false;

  private mounted() {
    // placeholder
  }

  private login(): void {
    const pwd = calcSHA256(this.password);
    this.password = "";

    let data: FormData = new FormData();
    data.append("userName", this.userName);
    data.append("password", pwd);

    axios.post(process.env.VUE_APP_login_url, data).then(
      response => {
        if (response.data.hasError) {
          throw response.data.data;
        }

        sessionStorage.setItem("auth", "passed");

        const payload = JSON.parse(response.data.data as string);
        this.$store.state.userID = payload.userID;
        this.$store.state.nickname = payload.nickname;
        this.$store.state.permission = payload.permission;

        this.userName = "";

        this.$store.state.isLogin = true;
        this.loginDialogController = false;
      }
    ).catch(
      err => {
        console.log("login failed, error:", err);
      }
    );
  }

  private exit(): void {
    this.$store.state.isLogin = false;

    sessionStorage.removeItem("auth");

    this.linkTo("", "home")
  }

  private linkTo(path: string, name: string): void {
    if (location.href.split("#/")[1] !== path) {
      this.$router.push({ name: name });
    }
  }

  private openLoginDialog(): void {
    this.userName = "";
    this.password = "";

    this.loginDialogController = true;
  }

  private cancelLogin(): void {
    this.loginDialogController = false;
  }
}
</script>

<style lang="scss">
.top-title {
  width: 10vw;
  line-height: 10rem;

  font-size: 4rem;
  font-weight: 600;
}

.top-title:hover {
  cursor: pointer;
}

.top-links {
  width: 65vw;
  display: flex;
  justify-content: flex-end;

  font-size: 2rem;

  .tl-item {
    width: 7vw;
    margin: auto 0;

    a {
      color: black;
      text-decoration: none;
    }
  }

  .tl-item:hover {
    color: darkgray;
    cursor: pointer;
    text-decoration-line: underline;
  }
}

.top-login-entrance {
  width: 10vw;
  margin: auto 0;
  padding: 0 2.5vw;

  .tle-text {
    font-size: 2rem;
  }
  .tle-text:hover {
    cursor: pointer;
  }
}

.top-login-dialog {
  text-align: left;

  .tld-title {
    font-size: 3rem;
    font-weight: 600;
  }

  .tld-content {
    padding: 3vh 20%;
    font-size: 2.5rem;

    .tldc-item {
      display: flex;
      padding: 2vh 0;

      .tldci-label {
        width: 20%;
        align-self: center;
      }

      .el-input {
        width: 80%;
      }
    }
  }
}
</style>
