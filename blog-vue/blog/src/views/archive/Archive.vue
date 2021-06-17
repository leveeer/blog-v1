<template>
  <div>
    <!-- banner -->
    <div class="archive-banner banner">
      <h1 class="banner-title">归档</h1>
    </div>
    <!-- 归档列表 -->
    <v-card class="blog-container">
      <timeline>
        <timeline-title> 目前共计{{ count }}篇文章，继续加油</timeline-title>
        <timeline-item v-for="item of archiveList" :key="item.id">
          <!-- 日期 -->
          <span class="time">{{ item.createTime | date }}</span>
          <!-- 文章标题 -->
          <router-link
            :to="'/articles/' + item.id"
            style="color:#666;text-decoration: none"
          >
            {{ item.articleTitle }}
          </router-link>
        </timeline-item>
      </timeline>
      <!-- 分页按钮 -->
      <v-pagination
        color="#00C4B6"
        v-model="current"
        :length="Math.ceil(count / 10)"
        total-visible="7"
      />
    </v-card>
  </div>
</template>

<script>
  import { Timeline, TimelineItem, TimelineTitle } from "vue-cute-timeline";
  import { getArchiveList } from "../../api/api";

  export default {
    created() {
      this.listArchives();
    },
    components: {
      Timeline,
      TimelineItem,
      TimelineTitle
    },
    data: function() {
      return {
        current: 1,
        count: 0,
        archiveList: []
      };
    },
    methods: {
      listArchives() {
        getArchiveList({
          cmdId: 2,
          currentPage: this.current,
        })
          .then((data) => {
            console.log(data)
            this.archiveList = data.archiveInfo.archiveList;
            this.count = data.archiveInfo.count;
          });
      }
    },
    watch: {
      current(value) {
        getArchiveList({
          cmdId: 2,
          current: value
        })
          .then((data) => {
            this.archiveList = data.archiveInfo.archiveList;
            this.count = data.archiveInfo.count;
          });
      }
    }
  };
</script>

<style scoped>
  .archive-banner {
    background: #49b1f5 url(https://www.static.talkxj.com/wallroom-1920x1080-bg-338d7bc.jpg) no-repeat center center;
  }

  .time {
    font-size: 0.75rem;
    color: #555;
    margin-right: 1rem;
  }
</style>
