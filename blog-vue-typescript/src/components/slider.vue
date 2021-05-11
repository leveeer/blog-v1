<template>
  <div class="right slider" style="position: fixed;right: 20px">
    <img class="right-logo" src="../assets/user_logo.jpg" alt="" />
    <div class="title">Einson</div>
    <div class="right-content"></div>
    <div class="tags">
      <div class="title">标签云</div>
      <router-link
        v-for="item in list"
        class="item"
        :key="item.uid"
        :to="
          `/articles?tagUid=${item.uid}&tag_name=${item.content}&category_id=`
        "
      >
        <span :key="item._id">{{ item.content }}</span>
      </router-link>
    </div>
    <div class="introduce">
      <div class="title">技术以内的 BB</div>
      <div class="content">
        <!--        <img style="width:100%;" src="../assets/BiaoChenXuYing.png" alt="全栈修炼" />-->
      </div>
    </div>
    <div class="introduce">
      <div class="title">技术以外的 BB</div>
      <div class="content">
        <!--        <img style="width:100%;" src="../assets/YingHeZaHuoPu.png" alt="硬核杂货铺" />-->
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { PageParams, Tag, TagsData } from "@/types/index";

@Component
export default class Slider extends Vue {
  private isLoadEnd: boolean = false;
  private isLoading: boolean = false;
  private list: Array<object> = [];
  private total: number = 0;
  private params: PageParams = {
    current: 1,
    size: 100,
    records: {},
    total: 0,
    optimizeCountSql: false,
    isSearchCount: false
  };

  mounted(): void {
    this.handleSearch();
  }

  private async handleSearch(): Promise<void> {
    this.isLoading = true;
    const data: TagsData = await this.$https.get(this.$urls.getTagList, {
      params: this.params
    });
    this.isLoading = false;

    this.list = [...this.list, ...data.list];
    this.total = data.count;
    this.params.current++;
    if (this.total === this.list.length) {
      this.isLoadEnd = true;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.slider {
  padding-top: 50px;
}

.right {
  text-align: center;

  .right-logo {
    width: 130px;
    border-radius: 50%;
    animation: mylogo 3s;
    -moz-animation: mylogo 3s; /* Firefox */
    -webkit-animation: mylogo 3s; /* Safari and Chrome */
    -o-animation: mylogo 3s; /* Opera */
    animation-iteration-count: infinite;
  }

  .title {
    font-size: 25px;
    font-weight: bold;
  }

  .right-content {
    padding: 10px 0 20px 0;
    margin-bottom: 10px;
    border-bottom: 1px solid #eee;

    .item {
      display: inline-block;
      padding: 0 10px;
      color: #969696;
      border-right: 1px solid #eee;

      &:last-child {
        border-right: none;
      }

      .num {
        color: #333;
      }
    }
  }

  .introduce {
    margin-bottom: 10px;
    border-bottom: 1px solid #eee;

    .title {
      font-size: 14px;
      color: #969696;
    }

    .content {
      color: #333;
      line-height: 26px;
      text-align: left;
      padding: 20px 0;
    }

    .footer {
      padding-bottom: 10px;
    }
  }

  .tags {
    min-height: 200px;
    padding: 5px 0 20px 0;
    margin-bottom: 10px;
    border-bottom: 1px solid #eee;

    .title {
      font-size: 14px;
      color: #969696;
    }

    .item {
      display: inline-block;
      cursor: pointer;
      padding: 5px 10px;
      border-radius: 5px;
      background-color: #eee;
      color: #333;
      margin: 10px 10px 0 0;
      text-decoration: none;

      &:hover {
        color: #409eff;
      }
    }
  }

  .classification {
    padding: 5px 0 20px 0;

    .title {
      font-size: 14px;
      color: #969696;
    }

    .item {
      text-align: center;
      padding: 10px;
      border-bottom: 1px solid #eee;
      color: #666;
      margin: 10px 10px 0 0;
    }
  }
}

@keyframes mylogo {
  0% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
  25% {
    transform: rotate(0deg) scale(1, 1);
    opacity: 1;
  }
  100% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
}

@-moz-keyframes mylogo {
  0% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
  25% {
    transform: rotate(0deg) scale(1, 1);
    opacity: 1;
  }
  100% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
}

@-webkit-keyframes mylogo {
  0% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
  25% {
    transform: rotate(0deg) scale(1, 1);
    opacity: 1;
  }
  100% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
}

@-o-keyframes mylogo {
  0% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
  25% {
    transform: rotate(0deg) scale(1, 1);
    opacity: 1;
  }
  100% {
    transform: rotate(0deg) scale(0.8, 0.8);
    opacity: 1;
  }
}
</style>
