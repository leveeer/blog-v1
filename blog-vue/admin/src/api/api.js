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
    var formdata = new FormData();
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

export function refreshToken() {
    return service.get(urls.refresh_token)
}
