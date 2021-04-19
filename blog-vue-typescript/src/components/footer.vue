<template>
  <div class="footer">酒馆酒 ©2021 Created by Einson</div>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Footer extends Vue {}
</script>
<style scoped>
.footer {
  text-align: center;
  padding: 20px;
  font-weight: bold;
  margin-bottom: 0;
  width: 100%;
}
</style>
