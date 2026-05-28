<template>
  <aside>
    <component :is="ui.rightPanelContext.component" v-bind="ui.rightPanelContext.props" />
  </aside>
</template>

<script setup>
import { useUiStore } from '@/stores/ui'

const ui = useUiStore()

</script>
