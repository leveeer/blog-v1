export default {
  getArticlesOnHome(params) {
    return this.https.get("/getArticleList", params);
  }
};
