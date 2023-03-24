<template>
  <div class="custom-lists">
    <div class="explanation">
      {{ $t('custom_lists.explanation') }}
    </div>
    <div class="lists">
      <div
        class="list shadow-3 cursor-pointer"
        v-for="list in lists"
        :key="list.name"
        @click="editList(list)"
      >
        {{ list.name }}
        <q-icon name="edit" />
      </div>
    </div>
    <q-btn
      class="new-list-btn"
      :label="$t('custom_lists.new_list')"
      color="primary"
      no-caps
      @click="listPopup = true"
    />
    <q-dialog v-model="listPopup">
      <CustomList :list="currentList" @validate="updateLists" />
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import CustomList from './CustomList.vue'

export default defineComponent({
  name: 'CustomLists',
  components: { CustomList },
  data() {
    return {
      listPopup: false,
      currentList: null,
      lists: [],
    }
  },
  props: {},
  methods: {
    editList(list) {
      this.currentList = list
      this.listPopup = true
    },
    updateLists(newList) {
      let updatedLists = []
      this.lists.forEach((list) => {
        if (list.name !== newList.name) {
          updatedLists.push(list)
        }
      })
      updatedLists.push(newList)
      window['go']['main']['App']
        ['UpdateCustomLists'](updatedLists)
        .then(() => {
          this.listPopup = false
          this.lists = updatedLists
          this.$q.notify(this.$t('notifications.custom_lists_updated'))
        })
    },
  },
  created() {
    window['go']['main']['App']['ReadCustomLists']().then((data) => {
      this.lists = data ?? []
      console.log(this.lists)
    })
  },
  watch: {
    listPopup(newVal) {
      if (!newVal) {
        this.currentList = null
      }
    },
  },
})
</script>
<style lang="scss" scoped>
.explanation {
  margin-bottom: 15px;
}
.lists {
  display: flex;
  flex-wrap: wrap;
  .list {
    padding: 10px;
    border-radius: 10px;
    margin-right: 15px;
  }
}
.new-list-btn {
  margin-top: 15px;
}
</style>
