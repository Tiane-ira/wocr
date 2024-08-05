<template>
    <svg class="icon" aria-hidden="true" :style="`height: ${size}px;width: ${size}px;margin-right: 10px;`">
        <use :xlink:href="`#icon-${name}`"></use>
    </svg>
</template>

<script lang="ts" setup>
defineOptions({
    name: 'Icon',
})
withDefaults(defineProps<{ name: string, size?: number }>(), { size: () => 20 })
</script>

<style lang="scss" scoped>
.icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
}
</style>