import Layout from "@/layout/index.vue";
import router from "../../router";
import store from "../../store";
import Vue from "vue";
import {getUserMenu} from "../../api/api";
import {getResultCode} from "../../utils/util";
import {resultMap} from "../../utils/constant";

export function generaMenu() {
  // 查询用户菜单
  getUserMenu().then((data) => {
    console.log(data)
    if (data.code === getResultCode(resultMap.SuccessOK)) {
      const userMenuList = data.userMenu;
      userMenuList.forEach(item => {
        console.log(item)
        if (item.icon != null) {
          item.icon = "iconfont " + item.icon;
        }
        if (item.component === "Layout") {
          item.component = Layout;
        }
        if (item.children && item.children.length > 0) {
          item.children.forEach(route => {
            route.icon = "iconfont " + route.icon;
            route.component = loadView(route.component);
          });
        }
      });
      // 添加侧边栏菜单
      store.commit("saveUserMenuList", userMenuList);
      console.log(store.state.userMenuList)
      // 添加菜单到路由
      router.addRoutes(userMenuList);
    } else {
      Vue.prototype.$message.error(data.message);
      Vue.prototype.router.push({path: "/login"})
    }
  });
}

function loadView(view) {
  // 路由懒加载
  return (resolve) => require([`@/views${view}`], resolve);
  // return () => import(`@/views${view}`)
}
