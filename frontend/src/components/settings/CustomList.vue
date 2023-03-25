<template>
  <div class="popup" id="new-custom-list">
    <q-input
      class="name"
      v-model="newList.name"
      outlined
      :label="$t('custom_lists.list_name')"
    >
    </q-input>
    <SearchBar
      :label="$t('custom_lists.filter_sources')"
      v-model="searchInput"
    />

    <div class="selected-sources-amount custom-chip">
      {{ newList.sources.length }} {{ $t('custom_lists.sources_amount') }}
    </div>
    <div class="source-picker shadow-3">
      <div class="source" v-for="source in filteredSources" :key="source">
        <q-checkbox
          v-model="newList.sources"
          :val="source"
          class="source-inner shadow-3"
        >
          <div class="source-name">
            {{ source }}
          </div>
        </q-checkbox>
      </div>
    </div>
    <q-btn
      class="validate"
      :label="$t('expressions.validate')"
      color="primary"
      no-caps
      @click="validate"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SearchBar from '../SearchBar.vue'

export default defineComponent({
  name: 'CustomList',
  components: { SearchBar },
  data() {
    return {
      newList: {
        name: '',
        sources: [],
      },
      availableSources: [],
      filteredSources: [],
      searchInput: '',
    }
  },
  props: {
    list: { type: Array },
  },
  methods: {
    validate() {
      this.$emit('validate', this.newList)
    },
  },
  computed: {},
  created() {
    window['go']['main']['App']
      ['GetWebsitesWithCategories']({ categories: ['all'] })
      .then((data) => {
        this.availableSources = data
        this.filteredSources = data
      })
    if (this.list) {
      this.newList = { ...this.list }
    }
  },
  watch: {
    searchInput() {
      this.filteredSources = []
      this.availableSources.forEach((source) => {
        if (source.indexOf(this.searchInput) >= 0)
          this.filteredSources.push(source)
      })
    },
  },
})
</script>
<style lang="scss" scoped>
.popup {
  height: 90vh;
  width: 400px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  align-items: center;
  .search-bar {
    margin: 10px 0px;
    border-radius: 15px !important;
  }
  .source-picker {
    margin-top: 5px;
    flex-grow: 1;
    width: 98%;
    overflow-y: scroll;
    padding: 5px;
    border-radius: 10px;
    .source {
      margin: 7px;
      .source-inner {
        width: 100%;
        display: flex;
        flex-direction: row-reverse;
        align-items: center;
        justify-content: space-between;
        padding: 5px;
        border-radius: 10px;
      }
    }
  }
  .validate {
    margin-top: 15px;
  }
}
</style>
<style lang="scss">
#new-custom-list {
  .search-bar {
    .q-field__control {
      border-radius: 15px !important;
    }
  }
}
</style>
