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
    return service.post(urls.upload_image, formdata)
}
