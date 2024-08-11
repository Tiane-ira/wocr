import { getConfigCount } from '@/utils/backend'
import { defineStore } from 'pinia'

export const ocrStore = defineStore('ocrStore', {
    state: () => {
        return {
            dirForm: {
                ocrPath: "",
                savePath: "",
                mode: "发票",
                recursive: false,
                id: "",
            },
            ocring: false,
            logList: [] as any[],
            foot: {
                skCount: 0,
            }
        }
    },
    actions: {
        async getSkCount() {
            this.foot.skCount = await getConfigCount()
        },
        async skCount(num) {
            this.foot.skCount += num
        }
    }
})