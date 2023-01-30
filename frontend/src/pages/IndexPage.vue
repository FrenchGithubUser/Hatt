<template>
  <div class="index-page" id="index-page">
    <q-input
      bg-color="blue-grey-3"
      borderless
      v-model="input"
      class="search-bar"
      @keyup.enter="search"
      filled
    >
      <template v-slot:append>
        <q-icon name="search" @click="search" />
      </template>
    </q-input>

    <CategorySelector ref="categories" />

    <SearchResults :results="results" />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { searchItems } from 'src/helpers/apiCalls.js'
import CategorySelector from 'src/components/categories/CategorySelector.vue'
import SearchResults from 'src/components/results/SearchResults.vue'

export default defineComponent({
  name: 'IndexPage',
  components: { CategorySelector, SearchResults },
  data() {
    return {
      input: '',
      results: {},
    }
  },
  methods: {
    search() {
      let categories = this.$refs.categories.getSelectedCategories
      searchItems({
        input: this.input,
        categories: categories.categories,
      }).then((data) => {
        if (data) {
          this.results = data
        }
      })
    },
  },
})
</script>

<style lang="scss" scoped>
.index-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  .search-bar {
    width: 40%;
    margin-top: 50px;
    margin-bottom: 20px;
  }
  .items-table {
    margin: 20px 0px;
    .thumbnail {
      max-width: 150px;
      max-height: 150px;
    }
  }
}
</style>
<style lang="scss">
#index-page {
  .search-bar {
    .q-field__control {
      border-radius: 15px !important;
    }
  }
}
</style>
