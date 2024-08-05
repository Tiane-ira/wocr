<template>
    <el-card shadow="always" :body-style="{ padding: '20px' }">
        <template #header>
            <div class="tc">
                <span>执行日志(自动清除)</span>
            </div>
        </template>
        <el-scrollbar ref="logRef" :height="logHeight">
            <div ref="innerRef">
                <p v-for="log, index in logList" :key="index" style="margin:0 2px;" class="scrollbar-demo-item">{{ log
                    }}
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