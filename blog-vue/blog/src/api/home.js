import service from "./https";


export function getArticlesOnHome(params) {
  return service.get("blog/getArticleList", params);
}
