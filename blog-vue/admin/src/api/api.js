import service from "./https";
import urls from "../utils/urls";
import store from "../store";


export function adminLogin(params) {
  return service.post(urls.login, params);
}

export function getUserMenu() {
  return service.get(urls.user_menu, {
        params: {token: store.state.token}
      }
  )
}
