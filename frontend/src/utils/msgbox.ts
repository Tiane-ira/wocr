import { ElMessage, ElMessageBox } from "element-plus"
import { OpenFile } from "../../wailsjs/go/service/SystemService";

export const ocrSucess = (path: string) => {
    ElMessageBox.confirm(
        `扫描成功,保存路径为:${path},确认打开吗?`,
        '请确认',
        {
            confirmButtonText: '打开',
            cancelButtonText: '关闭',
            type: 'success',
        }
    )
        .then(() => {
            // 打开文件
            console.log("打开文件: ", path);
            OpenFile(path)
        })
}

export const ocrFail = () => {
    ElMessageBox.alert(
        `扫描失败,请查看执行日志确认原因！`,
        '注意',
        {
            confirmButtonText: '好的',
            callback: () => { },
            type: "error"
        }
    )
}


