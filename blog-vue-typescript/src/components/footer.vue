<template>
  <el-container>
    <el-footer>酒馆酒 ©2021 Created by Einson</el-footer>
  </el-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Footer extends Vue {}
</script>
<style>
.el-footer {
  font-size: 14px;
  /*font-weight: bold;*/
  color: #f8f8f2;
  background-color: #000000;
  text-align: center;
  line-height: 60px;
  position: fixed;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 999;
}
</style>
