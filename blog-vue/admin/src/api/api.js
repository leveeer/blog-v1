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
