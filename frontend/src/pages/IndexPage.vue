<template>
  <div class="index-page" id="index-page">
    <div class="title">Hatt</div>
    <div class="quote">{{ $t('home.quote') }}</div>

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

    <CategorySelector ref="categories" @selection-updated="updateWebsites" />

    <SelectedWebsites
      ref="selectedWebsitesComponent"
      :websites="selectedWebsites"
      v-if="selectedWebsites.length !== 0"
    />

    <SearchResults :results="results" />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import CategorySelector from 'src/components/categories/CategorySelector.vue'
import SelectedWebsites from 'src/components/websites/SelectedWebsites.vue'
import SearchResults from 'src/components/results/SearchResults.vue'

export default defineComponent({
  name: 'IndexPage',
  components: { CategorySelector, SearchResults, SelectedWebsites },
  data() {
    return {
      input: '',
      results: {},
      selectedWebsites: [],
      searching: false,
    }
  },
  methods: {
    updateWebsites() {
      let categories = this.$refs.categories.getSelectedCategories
      window['go']['main']['App']['GetWebsites'](categories).then((data) => {
        this.selectedWebsites = data ?? []
      })
    },
    search() {
      this.searching = true
      let selectedWebsites = []
      if (this.$refs.selectedWebsitesComponent) {
        selectedWebsites =
          this.$refs.selectedWebsitesComponent.getSelectedWebsites
      }
      window['go']['main']['App']
        ['Search'](this.input, selectedWebsites)
        .then((data) => {
          console.log(data)
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
