import service from "./https";
import urls from "../utils/urls";


export function getArticlesOnHome(params) {
  return service.post(urls.articles, params);
}

export function getBlogInfo() {
  return service.get(urls.blog_info);
}

export function getArticleById(routePath) {
  return service.get(urls.article + routePath);
}

export function getArchiveList(params) {
  return service.post(urls.archive, params);
}

export function getCategories() {
  return service.get(urls.categories);
}

export function getTags() {
  return service.get(urls.tags);
}

export function getCategoryOrTagArticleList(path, params) {
  return service.get("/blog" + path, params);
}

export function getAbout() {
  return service.get(urls.about);
}

export function getMessages() {
  return service.get(urls.messages);
}

export function addMessages(params) {
  return service.post(urls.messages,params);
}
