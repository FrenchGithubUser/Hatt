<template>
  <div class="index-page" id="index-page">
    <div class="title">Hatt</div>
    <div class="quote">Stop searching, start finding</div>

    <q-input
      bg-color="blue-grey-3"
      v-model="input"
      class="search-bar"
      @keyup.enter="search"
      outlined
    >
      <template v-slot:append>
        <q-icon
          v-if="!searching"
          class="cursor-pointer"
          name="search"
          @click="search"
        />
        <q-spinner-puff v-else color="primary" size="1.5em" />
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
      searching: false,
    }
  },
  methods: {
    search() {
      this.searching = true
      let categories = this.$refs.categories.getSelectedCategories
      window['go']['main']['App']['Search'](this.input, categories).then((data) =>{
        this.searching = false
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
  .title {
    font-size: 3em;
    color: $primary;
    font-weight: bold;
    margin-top: 15px;
  }
  .quote {
    color: $primary;
  }
  .search-bar {
    width: 40%;
    margin-top: 20px;
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
