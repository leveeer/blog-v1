import service from "./https";
import urls from "../utils/urls";


export function getArticlesOnHome(params) {
  return service.get(urls.articles, params);
}

export function getBlogInfo() {
  return service.get(urls.blogInfo)
}
