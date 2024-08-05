<template>
  <div class="fp">
    <el-menu router @select="checkMenu">
      <el-menu-item v-for="(item, index) in $router.options.routes" :index="item.path"
        :class="$route.path == item.path ? 'is-active' : ''" :disabled="item.meta?.disable">
        <SvgIcon :name="item.meta?.icon as string"></SvgIcon>
        <template #title>{{ item.meta?.title }}</template>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script lang="ts" setup>
import SvgIcon from './SvgIcon.vue';
import { useRoute } from 'vue-router'

const router = useRoute()

const checkMenu = (e: any) => {
  console.log(e, router.path);

}
</script>

<style lang="scss" scoped>
.el-menu {
  border-right: 0;
}
</style>