<template>
  <div id="pro_form" style="width: 90%;margin-bottom: 100px">
    <div style="position: relative;top: 20px;left:100px;">
      <el-row>
        <el-col :span="6" v-for="(article, index) in articlesList" :key="article.uid" :offset="1"
                style="margin-bottom:30px" type="flex">
          <el-card :body-style="{ padding: '0px', height:'350px'}"
                   shadow="hover"
                   @click.native="getBlogByUid(article.uid)"
                   style="width: 270px;height: 350px">
            <img style="width: 100%;display: block"
                 src="../assets/logo5.jpg"
                 class="image"
                 alt=""/>
            <p>{{ article.title }}</p>
            <p>{{ article.createTime }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div style="display: flex;justify-content: center; margin-top: 25px">
      <el-pagination
          background
          @current-change="this.handleCurrentChange"
          @size-change="this.handleSizeChange"
          layout="sizes, prev, pager, next, jumper, ->, total, slot"
          :current-page="this.params.currentPage"
          :page-sizes="this.pageSizeList"
          :page-size="this.params.pageSize"
          :total="this.total">
      </el-pagination>
    </div>
    <LoadingCustom v-if="isLoading"></LoadingCustom>
    <LoadEnd v-if="isLoadEnd"></LoadEnd>
  </div>
</template>
<script lang="ts">
import {Component, Vue, Watch} from "vue-property-decorator";
import LoadEnd from "@/components/loadEnd.vue";
import LoadingCustom from "@/components/loading.vue";
import {ArticlesParams} from "@/types";

@Component({
  components: {
    LoadEnd,
    LoadingCustom
  }
})
export default class Articles extends Vue {
  isLoading = false;
  isLoadEnd = false;
  private articlesList: Array<object> = [];
  private total: number = 0;
  private pageSizeList: Array<number> = [];
  private currentPage: number = 1;
  private pageSize: number = 6
  private params: ArticlesParams = {
    adminUid: "",
    articlesPart: "",
    author: "",
    blogSort: "",
    blogSortUid: "",
    content: "",
    copyright: "",
    fileUid: "",
    isOriginal: "",
    isPublish: "",
    level: "",
    levelKeyword: undefined,
    orderByAscColumn: "",
    orderByDescColumn: "",
    outsideLink: "",
    parseCount: "",
    photoList: [],
    sort: 0,
    summary: "",
    tagList: "",
    tagUid: "",
    title: "",
    type: "",
    uid: "",
    userSort: 0,
    keyword: "",
    status: 1, // 文章发布状态 => 0 草稿，1 已发布,'' 代表所有文章
    currentPage: this.currentPage ? this.currentPage : 1,
    pageSize: this.pageSize ? this.pageSize : 6,
  };
  private href: string =
      process.env.NODE_ENV === "development"
          ? "http://localhost:3001/articleDetail?article_id="
          : "https://biaochenxuying.cn/articleDetail?article_id=";

  // lifecycle hook
  mounted(): void {
    // this.handleSearch();
    this.pageSizeList = [6, 12, 18, 24];
  }

  // @Watch("$route")
  // routeChange(val: Route, oldVal: Route): void {
  //   this.tag_name = decodeURI(getQueryStringByName("tag_name"));
  //   this.params.tagUid = getQueryStringByName("tag_id");
  //   this.articlesList = [];
  //   this.params.currentPage = 1;
  //   this.handleSearch();
  // }

  @Watch("params", {immediate: true, deep: true})
  onParamsChange(newVal: ArticlesParams, oldValue: ArticlesParams) {
    // this.pageSizeList = [6, 12, 18, 24];
    this.articlesList = []
    this.handleSearch();
  }

  // method
  getBlogByUid(id: string): void {
    this.$router.push({
      path:'/articleDetail',
      query:{
        id:id
      }
    })
  }

  handleSizeChange(val: number) {
    this.params.pageSize = val
  }

  handleCurrentChange(val: number) {
    this.params.currentPage = val
  }

  private async handleSearch(): Promise<void> {
    this.isLoading = true;
    const data = await this.$https.get(this.$urls.getArticleList, {
      params: this.params
    });
    this.isLoading = false;
    this.articlesList = [...this.articlesList, ...data.list];
    this.total = data.count;
  }
}
</script>

<style lang="less" scoped>

.left {
  .articles-list {
    margin: 0;
    padding: 0;
    list-style: none;

    .title {
      color: #333;
      margin: 7px 0 4px;
      display: inherit;
      font-size: 18px;
      font-weight: 700;
      line-height: 1.5;
    }

    .item > div {
      padding-right: 140px;
    }

    .item .wrap-img {
      position: absolute;
      top: 50%;
      margin-top: -50px;
      right: 0;
      width: 125px;
      height: 100px;

      img {
        width: 100%;
        height: 100%;
        border-radius: 4px;
        border: 1px solid #f0f0f0;
      }
    }

    li {
      line-height: 20px;
      position: relative;
      // width: 100%;
      padding: 15px 150px 15px 0;
      border-bottom: 1px solid #f0f0f0;
      word-wrap: break-word;
      cursor: pointer;

      &:hover {
        .title {
          color: #000;
        }
      }

      .abstract {
        min-height: 30px;
        margin: 0 0 8px;
        font-size: 13px;
        line-height: 24px;
        color: #555;
      }

      .meta {
        padding-right: 0 !important;
        font-size: 12px;
        font-weight: 400;
        line-height: 20px;

        a {
          margin-right: 10px;
          color: #b4b4b4;

          &:hover {
            transition: 0.1s ease-in;
            -webkit-transition: 0.1s ease-in;
            -moz-transition: 0.1s ease-in;
            -o-transition: 0.1s ease-in;
            -ms-transition: 0.1s ease-in;
          }
        }

        span {
          margin-right: 10px;
          color: #666;
        }
      }
    }
  }
}
</style>
