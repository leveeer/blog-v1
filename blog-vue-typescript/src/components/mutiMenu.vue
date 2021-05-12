<template>
  <div>
    <template v-for="item in navMenus">
      <!-- 含有子菜单 -->
      <template v-if="item.subTitle.length > 0">
        <!-- 第一层 含有子菜单菜单 -->
        <el-submenu
          :index="item.index"
          @click="goAnchor(item.anchor)"
          :key="item.index"
          @select="select"
        >
          <template slot="title">
            <!--<i class="el-icon-document"></i>-->
            <span slot="title">{{ item.text }}</span>
          </template>
          <muti-menu :navMenus="item.subTitle"></muti-menu
          ><!--递归调用-->
        </el-submenu>
      </template>
      <!-- 第一层 不含子菜单  -->
      <template v-else>
        <el-menu-item
          :index="item.index"
          @click="goAnchor(item.anchor)"
          :key="item.index"
          @select="select"
        >
          <!--<i class="el-icon-document"></i>-->
          <span slot="title">{{ item.text }}</span>
        </el-menu-item>
      </template>
    </template>
  </div>
</template>
<script>
export default {
  components: {},
  name: "mutiMenu",
  props: {
    navMenus: Array
  },
  data() {
    return {};
  },
  /* created() {
    this.navMenus.forEach((element, index) => {
      var elementById = document.getElementById(element.anchor);
      element.offsetTop = elementById.offsetTop
    })
    /!*window.addEventListener("scroll", this.handleScroll);*!/
  },*/

  activated(){
    console.log("111111111111111111111111")
    // window.addEventListener('scroll',this.handleScroll)
  },
  methods: {
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },

    /*handleScroll(e) {
      let scrollTop = document.documentElement.scrollTop + 120 //当前滚动距离
      this.navMenus.forEach((element, index) => {
        if ((scrollTop) >= element.offsetTop) {//当前滚动距离大于某一目录项时。
          for (let i = 0; i < index; i++) {
            document.getElementById(this.navMenus[i].id).className = "el-menu-item"  //同一时刻，只能有一个目录项的状态位为Active，即此时其他目录项的isActive = false
          }
          document.getElementById(element.id).className = "el-menu-item is-active" //将对应的目录项状态位置为true
        } else {
          document.getElementById(element.id).className = "el-menu-item"
        }
      })
    },*/

    goAnchor(anchor) {
      const element = document.getElementById(anchor);
      document.documentElement.scrollTop = element.offsetTop - 80;
    },

    select(index, indexPath) {
      console.log(index, "-", indexPath);
    }
  },
  mounted: function() {}
};
</script>

<style lang="less" scoped>
//深度作用选择器 /deep/  or  >>>

/* 覆盖导航样式 */
//.el-menu-item >>> .is-active {
/deep/.el-menu-item.is-active {
  background: #c6cbcd;
  border-left: 3px solid #8e96e7;
  font-weight: bold;
}

/deep/.el-submenu__title {
  display: flex;
  align-items: center;
}

/deep/.el-submenu__title span {
  white-space: normal;
  word-break: break-word;
  line-height: 15px;
}

/deep/.el-menu-item {
  display: flex;
  align-items: center;
  padding-right: 20px !important;
}

/deep/.el-submenu{
}

/deep/.el-menu-item span {
  white-space: normal;
  word-break: break-word;
  line-height: 15px;
}
</style>
