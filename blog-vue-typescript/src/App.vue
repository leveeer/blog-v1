<template>
  <div id="app"
       class="container">
    <Nav v-if="isShowNav" />
    <div class="layout">
      <router-view />
      <Slider v-if="isShowSlider"></Slider>
    </div>
    <ArrowUp></ArrowUp>
    <Footer v-show="isShowNav"></Footer>
  </div>
</template>
<script lang="ts">
import { Vue, Watch } from "vue-property-decorator";
import Component from "vue-class-component";
import { Route } from "vue-router";
import Nav from "@/components/nav.vue"; // @ is an alias to /src
import Slider from "@/components/slider.vue"; // @ is an alias to /src
import Footer from "@/components/footer.vue"; // @ is an alias to /src
import ArrowUp from "@/components/arrowUp.vue"; // @ is an alias to /src

@Component({
  components: {
    Nav,
    Slider,
    ArrowUp,
    Footer
  }
})
export default class App extends Vue {
  private isShowNav: boolean = false;
  private isShowSlider: boolean = false;
  mounted(): void {
    this.routeChange(this.$route, this.$route);
  }
  @Watch("$route")
  routeChange(val: Route, oldVal: Route): void {
    const referrer: any = document.getElementById("referrer");
    if (val.path === "/") {
      this.isShowNav = false;
      referrer.setAttribute("content", "always");
    } else {
      this.isShowNav = true;
      referrer.setAttribute("content", "never");
    }
    this.isShowSlider = val.path === "/articles" ||
        val.path === "/archive" ||
        val.path === "/project" ||
        val.path === "/timeline" ||
        val.path === "/message";
  }
}
</script>

<style lang="less">
@import url("./less/index.less");
@import url("./less/mobile.less");
#app {
  font-family: "PingFang SC", Monaco,serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 1400px;
  margin: 0 auto;
  padding-top: 60px;
}
img {
  vertical-align: bottom;
}
</style>
