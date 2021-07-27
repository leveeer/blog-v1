import service from "./https";
import urls from "../utils/urls";

export function getHomeData() {
    return service.get(urls.admin)
}

export function adminLogin(params) {
    return service.post(urls.login, params);
}

export function getUserMenu() {
    return service.get(urls.user_menu)
}

export function getArticleOptions() {
    return service.get(urls.article_options)
}

export function uploadImage(file) {
    const formdata = new FormData();
    formdata.append("file", file);
    let params = {
        index: 1,
    };
    formdata.append("index", params.index);
    return service.post(urls.upload_image, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}

export function addArticle(params) {
    return service.post(urls.articles, params)
}

export function updateArticle(params) {
    return service.put(urls.articles, params)
}

export function updateArticleStatus(params) {
    return service.put(urls.articles_status, params)
}

export function updateArticleTop(id, params) {
    return service.put(urls.articles_top + id, params)
}

export function deleteArticles(params) {
    return service.delete(urls.articles, params)
}

export function refreshToken() {
    return service.get(urls.refresh_token)
}

export function getArticleList(params) {
    return service.get(urls.articles, params)
}

export function getArticleByID(id) {
    return service.get(urls.articles + "/" + id)
}

export function getCategories(params) {
    return service.get(urls.categories, params)
}

export function addCategory(params) {
    return service.post(urls.categories, params)
}

export function updateCategory(params) {
    return service.put(urls.categories, params)
}

export function deleteCategory(params) {
    return service.delete(urls.categories, params)
}

export function getTags(params) {
    return service.get(urls.tags, params)
}

export function addTag(params) {
    return service.post(urls.tags, params)
}

export function updateTag(params) {
    return service.put(urls.tags, params)
}

export function deleteTag(params) {
    return service.delete(urls.tags, params)
}

export function getComments(params) {
    return service.get(urls.comments, params)
}

export function updateCommentStatus(params) {
    return service.put(urls.comments, params)
}

export function deleteComment(params) {
    return service.delete(urls.comments, params)
}

export function getMessages(params) {
    return service.get(urls.messages, params)
}

export function deleteMessage(params) {
    return service.delete(urls.messages, params)
}

export function getUsers(params) {
    return service.get(urls.users, params)
}

export function getUserRoles() {
    return service.get(urls.users_role)
}

export function UpdateUserRoles(params) {
    return service.put(urls.users_role, params)
}

export function UpdateUserStatus(params) {
    return service.put(urls.users_status, params)
}

export function getAdminRoles(params) {
    return service.get(urls.users_role, params)
}

export function getRoleResources() {
    return service.get(urls.role_resources)
}

export function getRoleMenus() {
    return service.get(urls.role_menus)
}

export function deleteRoles(params) {
    return service.delete(urls.roles, params)
}

export function saveOrUpdateRoles(params) {
    return service.post(urls.roles, params)
}



