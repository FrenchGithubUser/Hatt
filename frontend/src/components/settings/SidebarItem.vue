<template>
  <div
    :class="{
      'sidebar-item': true,
      'cursor-pointer': true,
      selected: selected,
      'shadow-1': hovering,
    }"
    @mouseover="hovering = true"
    @mouseout="hovering = false"
  >
    <div class="name">
      {{ $t('settings.' + name) }}
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SidebarItem',
  data() {
    return {
      hovering: false,
    }
  },
  props: {
    name: { type: String },
    selected: { type: Boolean },
  },
  methods: {},
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped>
.sidebar-item {
  padding: 10px;
  margin: 5px 0px;
  border-radius: 15px;
  &.selected {
    background-color: var(--q-primary);
    color: white;
  }
}
</style>
