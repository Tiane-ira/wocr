<template>
    <el-tag type="primary" round class="tag">当前密钥数:{{ store.foot.skCount }}</el-tag>
</template>

<script lang="ts" setup>
import { ocrStore } from '@/store/store'
import { onMounted } from 'vue';
defineOptions({
    name: 'LayoutFooter',
})

const store = ocrStore()

onMounted(async () => {
    await store.getSkCount()
})

</script>

<style lang="scss" scoped>
.tag {
    margin: 2px 5px;
}
</style>