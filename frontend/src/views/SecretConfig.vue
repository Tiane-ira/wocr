<template>
    <el-button type="primary" class="w100 mb10" @click="onAddItem">
        添加密钥
    </el-button>
    <el-table :data="tableData" border style="width: 100%" :height="tableHeight" :cell-style="{ textAlign: 'center' }"
        :header-cell-style="{ 'text-align': 'center' }">
        <el-table-column type="index" label="序号" width="60" />
        <el-table-column prop="name" label="名称" width="80" />
        <el-table-column prop="type" label="厂商" width="80" />
        <el-table-column prop="ak" label="AK" show-overflow-tooltip />
        <el-table-column prop="sk" label="SK" show-overflow-tooltip />
        <el-table-column prop="date" label="添加时间" width="160" />
        <el-table-column fixed="right" label="操作" width="60">
            <template #default="scope">
                <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.row.id)">
                    移除
                </el-button>
            </template>
        </el-table-column>
    </el-table>
    <SkDialog ref="SL" @save="refreshList" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import SkDialog from '@/components/SkDialog.vue';
import { getConfigs, removeConfig } from '@/utils/backend';
import { ElMessage } from 'element-plus';
import { ocrStore } from '@/store/store'

defineOptions({
    name: "SecretConfig"
})
const store = ocrStore()
const tableHeight = ref(625)
const SL = ref<any>(null)

const tableData = ref<any[]>([])

const deleteRow = async (id: string) => {
    const success = await removeConfig(id)
    if (success) {
        ElMessage.success("删除成功")
        store.skCount(-1)
        await refreshList()
    }
}
const onAddItem = () => {
    store.skCount(1)
    SL.value.show = true
}

const refreshList = async () => {
    tableData.value = await getConfigs()
}

const setTableHeight = () => {
    tableHeight.value = window.innerHeight - 100
}

onMounted(async () => {
    await refreshList()
    console.log(tableData.value);

    window.addEventListener("resize", setTableHeight)
})

</script>

<style scoped></style>