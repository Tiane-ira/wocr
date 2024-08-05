<template>
    <el-transfer class="el_tsf" v-model="rightValue" :data="data" :titles="['保留字段列表', '待导出字段列表']" @change="change" />
</template>

<script lang="ts" setup>
import { getFields, updateFields } from '@/utils/backend';
import { onMounted, ref } from 'vue';

defineOptions({
    name: 'ExportField',
})
const data = ref<any>([])
const rightValue = ref<any>([])

const change = async (idList) => {
    console.log(idList)
    await updateFields(idList)
}

const refreshData = async () => {
    let fieldList = await getFields()
    console.log(fieldList);
    for (const field of fieldList) {
        if (field.export) {
            rightValue.value.push(field.id)
        }
        data.value.push({
            key: field.id,
            label: field.name,
            disable: false,
        })
    }
}

onMounted(async () => {
    await refreshData()
})
</script>

<style lang="scss" scoped>
.el_tsf {
    :deep(.el-transfer-panel) {
        width: calc((100% - 170px)/2);
    }

    :deep(.el-transfer-panel__body) {
        height: calc((100vh - 105px));
    }

    :deep(.el-transfer-panel__list) {
        height: calc((100vh - 105px));
    }
}
</style>