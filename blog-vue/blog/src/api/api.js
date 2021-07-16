import service from "./https";
import urls from "../utils/urls";


export function getArticlesOnHome(params) {
  return service.get(urls.articles, params);
}

export function getBlogInfo() {
  return service.get(urls.blog_info);
}

export function getArticleById(routePath) {
  return service.get(urls.article + routePath);
}

export function getArchiveList(params) {
  return service.get(urls.archive, params);
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
  return service.post(urls.messages, params);
}

export function getLinks() {
  return service.get(urls.links);
}

export function getComments(params) {
  return service.get(urls.comments, params);
}

export function addComments(params) {
  return service.post(urls.comments, params);
}

export function likeArticle(params) {
  return service.post(urls.like_article, params);
}

export function getReplies(commentId, params) {
  return service.get(urls.replies + commentId, params);
}

export function getLoginCode(params) {
  return service.get(urls.code, params);
}

export function register(params) {
  return service.post(urls.register,params)
}

export function login(params) {
  return service.post(urls.login,params)
}
