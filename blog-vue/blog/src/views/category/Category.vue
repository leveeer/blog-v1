<template>
  <div>
    <!-- banner -->
    <div class="category-banner banner">
      <h1 class="banner-title">分类</h1>
    </div>
    <!-- 分类列表 -->
    <v-card class="blog-container">
      <!--  词云  -->
      <div id="char2" ref="wordCloud"></div>
      <div class="category-title">分类<span style="font-size: medium;color: #6c9d8f"><{{ count }}个></span></div>
      <ul class="category-list">
        <li
          class="category-list-item"
          v-for="item of categoryList"
          :key="item.id"
        >
          <router-link :to="'/categories/' + item.id">
            {{ item.categoryName }}
            <span class="category-count">({{ item.articleCount ? item.articleCount : 0 }})</span>
          </router-link>
        </li>
      </ul>
    </v-card>
  </div>
</template>

<script>
  import { getCategories } from "../../api/api";
  let echarts = require('echarts/lib/echarts');
  require('echarts-wordcloud');
  import 'echarts/theme/macarons.js'

  export default {
    mounted() {
      getCategories().then(data => {
        this.categoryList = data.categories;
        this.count = data.categories.length;
        if (data.categories.length != null) {
          data.categories.forEach(item => {
            this.option.series[0].data.push({
              name: item.categoryName,
              value: item.articleCount,
            })
          });
          echarts.init(this.$refs.wordCloud).setOption(this.option)
        }
      });
    },
    data() {
      return {
        wordList: [],
        categoryList: [],
        count: 0,
        option: {
          title: {
            x: "center"
          },
          tooltip: {
            pointFormat: "{series.name}: <b>{point.percentage:.1f}%</b>"
          },
          series: [{
            type: "wordCloud",
            //用来调整词之间的距离
            gridSize: 10,
            //用来调整字的大小范围
            sizeRange: [14, 60],
            // Text rotation range and step in degree. Text will be rotated randomly in range [-90,90] by rotationStep 45
            //用来调整词的旋转方向，，[0,0]--代表着没有角度，也就是词为水平方向，需要设置角度参考注释内容
            // rotationRange: [-45, 0, 45, 90],
            // rotationRange: [ 0,90],
            rotationRange: [0, 0],
            //随机生成字体颜色
            textStyle: {
              normal: {
                color: function() {
                  return (
                    "rgb(" +
                    Math.round(Math.random() * 255) +
                    ", " +
                    Math.round(Math.random() * 255) +
                    ", " +
                    Math.round(Math.random() * 255) +
                    ")"
                  );
                }
              }
            },
            //位置相关设置
            left: "center",
            top: "center",
            right: null,
            bottom: null,
            width: "200%",
            height: "200%",
            //数据
            data: []
          }]
        }
      };
    },
    methods: {},
  };
</script>

<style scoped>
.category-banner {
  background: #49b1f5 url(https://www.static.talkxj.com/wallhaven-13mk9v.jpg) no-repeat center center;
}
.category-title {
  text-align: center;
  font-size: 36px;
  line-height: 2;
}
@media (max-width: 759px) {
  .category-title {
    font-size: 28px;
  }
}
.category-list {
  margin: 0 1.8rem;
  list-style: none;
}
.category-list-item {
  padding: 8px 1.8rem 8px 0;
}
.category-list-item:before {
  display: inline-block;
  position: relative;
  left: -0.75rem;
  width: 12px;
  height: 12px;
  border: 0.2rem solid #49b1f5;
  border-radius: 50%;
  background: #fff;
  content: "";
  transition-duration: 0.3s;
}
.category-list-item:hover:before {
  border: 0.2rem solid #ff7242;
}
.category-list-item a:hover {
  transition: all 0.3s;
  color: #8e8cd8;
}

.category-list-item a:not(:hover) {
  transition: all 0.3s;
}

.category-count {
  margin-left: 0.5rem;
  font-size: 0.75rem;
  color: #858585;
}

#char2 {
  float: left;
  width: 100%;
  height: 300px;
  margin-top: -20px;
  border: 0 solid #E4E4E4;
  background: #ffffff;
  border-radius: 6px;
}
</style>
