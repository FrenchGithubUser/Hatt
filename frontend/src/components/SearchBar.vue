<template>
  <q-input
    :bg-color="$q.dark.isActive ? bgColorDark : bgColorWhite"
    :label="label"
    class="search-bar"
    outlined
    @keyup.enter="$emit('search')"
  >
    <template v-slot:append>
      <q-icon
        class="cursor-pointer"
        name="search"
        v-if="!searching"
        @click="$emit('search')"
      />
      <q-spinner-puff v-else color="primary" size="1.5em" />
      <slot></slot>
    </template>
  </q-input>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SearchBar',
  components: {},
  props: {
    label: { type: String },
    searching: { type: Boolean, default: false },
    bgColorWhite: { type: String, default: 'blue-grey-1' },
    bgColorDark: { type: String, default: 'blue-grey-7' },
  },
  data() {
    return {}
  },
  created() {},
  methods: {},
  watch: {},
  computed: {},
})
</script>
<style lang="scss" scoped></style>
