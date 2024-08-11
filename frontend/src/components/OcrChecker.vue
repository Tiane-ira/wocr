<template>
    <el-card shadow="hover" :body-style="{ padding: '20px' }" class="mb10">
        <el-form :model="dirForm" ref="form" :rules="rules" inline>
            <el-form-item label="扫描路径" prop="ocrPath" required style="display: block;">
                <el-input v-model="dirForm.ocrPath" placeholder="选择扫描目录">
                    <template #suffix>
                        <el-button type="text" icon="FolderOpened" @click="selectOcrPath"></el-button>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item label="保存路径" prop="savePath" required style="display: block;">
                <el-input v-model="dirForm.savePath" placeholder="选择保存目录">
                    <template #suffix>
                        <el-button type="text" @click="selectSavePath" icon="FolderOpened"></el-button>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item label="密钥" prop="id" required>
                <el-select v-model="dirForm.id" placeholder="选择密钥" clearable filterable no-data-text="暂无密钥" no-match-text="无数据" size="small"
                    style="width: 120px;">
                    <el-option v-for="item in configList" :key="item.id" :label="`${item.type}/${item.name}`"
                        :value="item.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="模式" prop="mode" required>
                <el-select v-model="dirForm.mode" placeholder="识别模式" size="small"
                    style="width: 80px;">
                    <el-option v-for="item,index in modeList" :key="index" :label="item"
                        :value="item">
                    </el-option>
                </el-select>

            </el-form-item>
            <el-form-item label="是否递归" prop="recursive">
                <el-switch v-model="dirForm.recursive" />
            </el-form-item>
        </el-form>

        <el-button :loading="state.ocring" type="primary" size="default" @click="doOcr" class="w100">开始扫描</el-button>
    </el-card>

</template>

<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue';
import { eventLogOff, eventLogOn, getConfigs, handleSelectDir, startOcr } from "@/utils/backend";
import { ocrStore } from '@/store/store'
import { ocrFail, ocrSucess } from '@/utils/msgbox';

const state = ocrStore()
const { dirForm, logList } = state

defineOptions({
    name: 'OcrChecker',
})

const modeList = ref(["发票","VIN码","行程单"])

const configList = ref<any[]>([])
const rules = reactive<any>({
    ocrPath: [
        { required: true, message: 'OCR识别目录不能为空', trigger: 'blur' },
    ],
    savePath: [
        { required: true, message: '结果保存目录不能为空', trigger: 'blur' },
    ],
    id: [
        { required: true, message: '密钥必选', trigger: 'blur' },
    ],
    mode: [
        { required: true, message: '模式必选', trigger: 'blur' },
    ],
})

const form = ref<any>(null)

const submit = async () => {
    state.ocring = true
    state.logList.splice(0)
    eventLogOn((log: any) => { logList.push(log) })
    const resultPath = await startOcr(dirForm)
    if (resultPath !== "") {
        ocrSucess(resultPath)
    } else {
        ocrFail()
    }
    eventLogOff()
    state.ocring = false
}

const doOcr = async () => {
    await form.value.validate((valid, fields) => {
        if (valid) {
            submit()
        } else {
            console.log('参数非法:', fields)
        }
    })
}


const selectOcrPath = async () => {
    dirForm.ocrPath = await handleSelectDir()
}
const selectSavePath = async () => {
    dirForm.savePath = await handleSelectDir()
}

onMounted(async () => {
    let configs = await getConfigs()
    dirForm.id = ""
    configList.value = configs
    let defaultSk = configs.filter(item => item.isDefault === true)
    if (defaultSk.length > 0) {
        dirForm.id = defaultSk[0].id
    }else if (configs.length > 0) {
        dirForm.id = configs[0].id
    }
})

</script>

<style lang="scss" scoped></style>