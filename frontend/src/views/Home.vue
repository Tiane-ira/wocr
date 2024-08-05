<template>
    <Checker />
    <LogWin />
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue';
import LogWin from '@/components/LogWin.vue';
import Checker from '@/components/OcrChecker.vue';
defineOptions({
    name: 'Home',
})

onMounted(() => {

})
</script>

<style lang="scss" scoped></style>