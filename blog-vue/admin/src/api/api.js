import service from "./https";
import urls from "../utils/urls";


export function adminLogin(params) {
  return service.post(urls.login, params);
}
