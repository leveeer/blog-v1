<template>
  <div>
    <!-- 遍历菜单 -->
    <template v-for="(item,index) in navMenus">
      <!-- 含有子菜单 -->
      <template v-if="item.subTitle.length > 0">
        <!-- 第一层 含有子菜单菜单 -->
        <el-submenu :index="item.anchor" @click="goAnchor(item.anchor)" :key="item.id">
          <template slot="title">
            <span slot="title" :class="{'highlight-title':item.isActive}">{{ item.text }}</span>
          </template>
          <muti-menu :navMenus="item.subTitle"></muti-menu><!--递归调用-->
        </el-submenu>
      </template>
      <!-- 第一层 不含子菜单  -->
      <template v-else>
        <el-menu-item :index="item.anchor" @click="goAnchor(item.anchor)" :key="item.id">
          <!--<a style="color: #000000" :href="'#' + item.anchor"></a>-->
          <span>{{ item.text }}</span>
        </el-menu-item>
      </template>
    </template>
  </div>
</template>
<script>
export default {
  name: 'mutiMenu',
  props: {
    navMenus: Array,
  },
  data() {
    return {}
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },

    goAnchor(anchor) {
      console.log(anchor)
      // window.location.href = '#' + anchor;
      const element = document.getElementById(anchor);
      console.log(element)
      document.documentElement.scrollTop = element.offsetTop - 60
    }
  },
  mounted: function () {

  }
}
</script>

<style scoped>
.highlight-title {
  border-left: 5px solid rgb(15, 116, 223);
  background-color: rgb(243, 243, 243);
  z-index: -1;
}
</style>
