<template>
  <el-card class="main-card">
    <div class="title">{{ this.$route.name }}</div>
    <!-- 表格操作 -->
    <div class="operation-container">
      <el-button
              v-if="isDelete === 0"
              type="danger"
              size="small"
              icon="el-icon-deleteItem"
              :disabled="articleIdList.length === 0"
              @click="updateIsDelete = true"
      >
        批量删除
      </el-button>
      <el-button
              v-else
              type="danger"
              size="small"
              icon="el-icon-deleteItem"
              :disabled="articleIdList.length === 0"
              @click="remove = true"
      >
        批量删除
      </el-button>
      <!-- 条件筛选 -->
      <div style="margin-left:auto">
        <el-select
          v-model="condition"
          placeholder="请选择"
          size="small"
          style="margin-right:1rem"
        >
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-input
          v-model="keywords"
          prefix-icon="el-icon-search"
          size="small"
          placeholder="请输入文章名"
          style="width:200px"
          @keyup.enter.native="listArticles"
        />
        <el-button
          type="primary"
          size="small"
          icon="el-icon-search"
          style="margin-left:1rem"
          @click="listArticles"
        >
          搜索
        </el-button>
      </div>
    </div>
    <!-- 表格展示 -->
    <el-table
      border
      :data="articleList"
      @selection-change="selectionChange"
      v-loading="loading"
    >
      <!-- 表格列 -->
      <el-table-column type="selection" width="55" />
      <el-table-column prop="articleTitle" label="标题" align="center" />
      <!-- 文章分类 -->
      <el-table-column
        prop="categoryName"
        label="分类"
        width="120"
        align="center"
      />
      <!-- 文章标签 -->
      <el-table-column
        prop="tagDTOList"
        label="标签"
        width="180"
        align="center"
      >
        <template slot-scope="scope">
          <el-tag
                  v-for="item of scope.row.tagList"
                  :key="item.id"
                  style="margin-right:0.2rem;margin-top:0.2rem"
          >
            {{ item.tagName }}
          </el-tag>
        </template>
      </el-table-column>
      <!-- 文章浏览量 -->
      <el-table-column
        prop="viewsCount"
        label="浏览量"
        width="80"
        align="center"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.viewsCount">
            {{ scope.row.viewsCount }}
          </span>
          <span v-else>0</span>
        </template>
      </el-table-column>
      <!-- 文章点赞量 -->
      <el-table-column
        prop="likeCount"
        label="点赞量"
        width="80"
        align="center"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.likeCount">
            {{ scope.row.likeCount }}
          </span>
          <span v-else>0</span>
        </template>
      </el-table-column>
      <!-- 文章发表时间 -->
      <el-table-column
        prop="createTime"
        label="发表时间"
        width="140"
        align="center"
      >
        <template slot-scope="scope">
          <i class="el-icon-time" style="margin-right:5px" />
          {{ scope.row.createTime | date }}
        </template>
      </el-table-column>
      <!-- 文章修改时间 -->
      <el-table-column
        prop="updateTime"
        label="更新时间"
        width="140"
        align="center"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.updateTime">
            <i class="el-icon-time" style="margin-right:5px" />
            {{ scope.row.updateTime | date }}
          </span>
          <span v-else>无</span>
        </template>
      </el-table-column>
      <!-- 文章置顶 -->
      <el-table-column prop="isTop" label="置顶" width="100" align="center">
        <template slot-scope="scope">
          <el-switch
                  v-model="scope.row.isTop"
                  active-color="#13ce66"
                  inactive-color="#F4F4F5"
                  :disabled="(scope.row.isDelete === 1 || scope.row.isPublish === 0) || (scope.row.isDelete === 1 || scope.row.isPublish == null)"
                  :active-value="1"
                  :inactive-value="0"
                  @change="changeTop(scope.row)"
          />
        </template>
      </el-table-column>
      <!-- 列操作 -->
      <el-table-column label="操作" align="center" width="160">
        <template slot-scope="scope">
          <el-button
                  type="primary"
                  size="mini"
                  @click="editArticle(scope.row.id)"
                  v-if="scope.row.isDelete === 0 || scope.row.isDelete == null "
          >
            编辑
          </el-button>
          <el-popconfirm
                  title="确定删除吗？"
                  style="margin-left:10px"
                  @onConfirm="updateArticleStatus(scope.row.id)"
                  v-if="scope.row.isDelete === 0 || scope.row.isDelete == null"
          >
            <el-button size="mini" type="danger" slot="reference">
              删除
            </el-button>
          </el-popconfirm>
          <el-popconfirm
                  title="确定恢复吗？"
                  v-if="scope.row.isDelete === 1"
                  @onConfirm="updateArticleStatus(scope.row.id)"
          >
            <el-button size="mini" type="success" slot="reference">
              恢复
            </el-button>
          </el-popconfirm>
          <el-popconfirm
                  style="margin-left:10px"
                  v-if="scope.row.isDelete === 1"
                  title="确定彻底删除吗？"
                  @onConfirm="deleteArticles(scope.row.id)"
          >
            <el-button size="mini" type="danger" slot="reference">
              删除
            </el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页 -->
    <el-pagination
      class="pagination-container"
      background
      @size-change="sizeChange"
      @current-change="currentChange"
      :current-page="current"
      :page-size="size"
      :total="count"
      :page-sizes="[10, 20]"
      layout="total, sizes, prev, pager, next, jumper"
    />
    <!-- 批量逻辑删除对话框 -->
    <el-dialog :visible.sync="updateIsDelete" width="30%">
      <div class="dialog-title-container" slot="title">
        <i class="el-icon-warning" style="color:#ff9900" />提示
      </div>
      <div style="font-size:1rem">是否删除选中项？</div>
      <div slot="footer">
        <el-button @click="updateIsDelete = false">取 消</el-button>
        <el-button type="primary" @click="updateArticleStatus(null)">
          确 定
        </el-button>
      </div>
    </el-dialog>
    <!-- 批量彻底删除对话框 -->
    <el-dialog :visible.sync="remove" width="30%">
      <div class="dialog-title-container" slot="title">
        <i class="el-icon-warning" style="color:#ff9900" />提示
      </div>
      <div style="font-size:1rem">是否彻底删除选中项？</div>
      <div slot="footer">
        <el-button @click="remove = false">取 消</el-button>
        <el-button type="primary" @click="deleteArticles(null)">
          确 定
        </el-button>
      </div>
    </el-dialog>
  </el-card>
</template>

<script>
  import {deleteArticles, getArticleList, updateArticleStatus, updateArticleTop} from "../../api/api";
  import {getResultCode} from "../../utils/util";
  import {resultMap} from "../../utils/constant";

  export default {
    created() {
      this.listArticles();
    },
    data: function () {
      return {
        loading: true,
        updateIsDelete: false,
        remove: false,
        options: [
          {
            value: '{"isDelete":0,"isPublish":1}',
            label: "已发布"
          },
          {
            value: '{"isDelete":1}',
            label: "回收站"
          },
          {
            value: '{"isDelete":0,"isPublish":0}',
            label: "草稿箱"
          }
        ],
        condition: '{"isDelete":0,"isPublish":1}',
        articleList: [],
        articleIdList: [],
        keywords: "",
        isDelete: 0,
        isPublish: 1,
        current: 1,
        size: 10,
        count: 0
      };
    },
    methods: {
      selectionChange(articleList) {
        this.articleIdList = [];
        articleList.forEach(item => {
          this.articleIdList.push(item.id);
        });
      },
      editArticle(id) {
        this.$router.push({path: "/articles/" + id});
      },
      notify(data) {
        if (data.code === getResultCode(resultMap.SuccessOK)) {
          this.$notify.success({
            title: "成功",
            message: data.message
          });
        } else {
          this.$notify.error({
            title: "失败",
            message: data.message
          });
        }
        this.addOrEdit = false;
      },
      updateArticleStatus(id) {
        if (id != null) {
          this.articleIdList = [];
          this.articleIdList.push(id)
        }
        updateArticleStatus({
          articleStatus: {
            articleIdList: this.articleIdList,
            isDelete: this.isDelete === 0 ? 1 : 0
          }
        }).then(data => {
          this.articleIdList = [];
          this.notify(data);
          this.updateIsDelete = false;
        });
      },
      deleteArticles(id) {
        if (id != null) {
          this.articleIdList = [];
          this.articleIdList.push(id)
        }
        deleteArticles({
          data: {
            articleIds: {
              articleIdList: this.articleIdList,
            }
          }
        }).then(data => {
          this.articleIdList = [];
          this.notify(data);
          this.remove = false;
        });
      },
      sizeChange(size) {
        this.size = size;
        this.listArticles();
      },
      currentChange(current) {
        this.current = current;
        this.listArticles();
      },
      changeTop(article) {
        updateArticleTop(article.id, {
          articleTop: {
            isTop: article.isTop
          }
        }).then(data =>{
          this.notify(data)
        })
      },
      listArticles() {
        getArticleList({
          params: {
            current: this.current,
            size: this.size,
            keywords: this.keywords,
            isDelete: this.isDelete,
            isPublish: this.isPublish
          }
        }).then((data) => {
          console.log(data);
          this.articleList = data.adminArticle.articleList;
          this.count = data.adminArticle.count;
          this.loading = false;
        });
      }
    },
    watch: {
      condition() {
        const condition = JSON.parse(this.condition);
        this.isDelete = condition.isDelete;
        this.isPublish = condition.isPublish;
        this.listArticles();
      }
    }
  };
</script>
