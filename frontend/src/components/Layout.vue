<template>
  <el-row class="top">
    <el-col :span="4">
      <Aside />
    </el-col>
    <div class="bgg" style="height: 100%;width:1px;"></div>
    <el-col style="margin-left: -1px;" :span="20" class="fp main">
      <RouterView></RouterView>
    </el-col>
  </el-row>
  <el-row class="footer">
    <Footer />
  </el-row>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router'
import Aside from '@/components/LayoutAside.vue';
import Footer from '@/components/LayoutFooter.vue';
</script>

<style lang="scss" scoped>
.top {
  height: calc(100vh - 30px);
}

.footer {
  border-top: 1px solid #EBEEF5;
  height: 30px;
}
</style>