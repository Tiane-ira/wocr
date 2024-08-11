<template>
    <el-card shadow="hover" :body-style="{ padding: '20px' }">
        <template #header>
            <div class="tc">
                <el-tag type="info" size="large"  effect="dark">执行日志(自动清除)</el-tag>
            </div>
        </template>
        <el-scrollbar ref="logRef" :height="logHeight">
            <div ref="innerRef">
                <p v-for="log, index in logList" :key="index" style="margin:0 2px;" class="scrollbar-demo-item">
                    <el-text :type="log.isError?'danger':'success'" size="small">{{ log.msg }}</el-text>
                </p>
            </div>
        </el-scrollbar>
    </el-card>
</template>

<script lang="ts" setup>
import { nextTick, onMounted, ref, watch } from 'vue';
import { ocrStore } from '@/store/store'

defineOptions({
    name: 'LogWin',
})

const { logList } = ocrStore()
let logHeight = ref(260)
const logRef = ref<any>(null)
const innerRef = ref<any>(null)

const setLogHeight = () => {
    logHeight.value = window.innerHeight - 470
    console.log(logHeight.value);
}

watch(() => logList.length, () => {
    setScrollBottom()
}, { deep: true })

const setScrollBottom = async () => {
    await nextTick()
    const max = innerRef.value!.clientHeight
    logRef.value!.setScrollTop(max)
}

onMounted(() => {
    window.addEventListener('resize', setLogHeight);
})
</script>

<style lang="scss" scoped></style>