<template>
  <div class="custom-lists">
    <div class="explanation">
      {{ $t('custom_lists.explanation') }}
    </div>
    <q-btn :label="$t('custom_lists.new_list')" color="primary" no-caps />
    <q-dialog v-model="newListPopup">
      <NewCustomList />
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import NewCustomList from './NewCustomList.vue'

export default defineComponent({
  name: 'CustomLists',
  components: { NewCustomList },
  data() {
    return {
      newListPopup: true,
    }
  },
  props: {},
  methods: {},
  computed: {},
  created() {},
})
</script>
<style lang="scss" scoped></style>
