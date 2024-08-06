<template>
    <el-dialog v-model="show" title="添加密钥" width="500">
        <el-form ref="formRef" :model="form" :rules="rules" label-position="left" label-width="100">
            <el-form-item label="厂商" prop="type">
                <el-select v-model="form.type" placeholder="选择厂商">
                    <el-option label="百度云" value="百度云" />
                    <el-option label="腾讯云" value="腾讯云" />
                    <el-option label="阿里云" value="阿里云" />
                </el-select>
            </el-form-item>
            <el-form-item label="名称" prop="name" required>
                <el-input v-model="form.name" placeholder="填入密钥的名称" />
            </el-form-item>
            <el-form-item label="AccessKey" prop="ak" required>
                <el-input v-model="form.ak" placeholder="填入云厂商AccessKey" />
            </el-form-item>
            <el-form-item label="SecretKey" prop="sk" required>
                <el-input v-model="form.sk" placeholder="填入云厂商SecretKey" />
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="cancle">取消</el-button>
                <el-button type="primary" @click="saveConf">
                    保存
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import { checkConfig, saveConfig } from '@/utils/backend';
import { ElMessage } from 'element-plus';
import { reactive, ref } from 'vue';
defineOptions({
    name: 'SkDialog',
})
let show = ref(false)
const form = reactive<any>({
    id: '',
    name: '',
    ak: '',
    sk: '',
    type: "百度云",
})
const rules = reactive<any>({
    name: [
        { required: true, message: '名称不能为空', trigger: 'blur' },
        { min: 2, max: 8, message: '名称长度必须为2-8', trigger: 'blur' },
    ],
    ak: [
        { required: true, message: 'AccessKey不能为空', trigger: 'blur' },
    ],
    sk: [
        { required: true, message: 'SecretKey不能为空', trigger: 'blur' },
    ],
})
const emit = defineEmits(['save'])

const formRef = ref<any>(null)
const saveConf = async () => {
    await formRef.value.validate((valid, fields) => {
        if (valid) {
            doSave()
        }
    })

}

const doSave = () => {
    checkConfig(form).then(res => {
        if (res !== "") {
            ElMessage.error(res)
        } else {
            saveConfig(form).then(success => {
                if (success) {
                    ElMessage.success("添加成功")
                    show.value = false
                    emit('save')
                    resetForm()
                } else {
                    ElMessage.error("添加失败")
                }
            })
        }
    })
}

const cancle = () => {
    show.value = false
    resetForm()
}

const resetForm = () => {
    form.id = ""
    form.name = ""
    form.ak = ""
    form.sk = ""
    form.type = "百度云"
}

defineExpose({
    show
});

</script>

<style lang="scss" scoped></style>