<template>
  <div class="index-page" id="index-page">
    <img src="images/hatt-logo.png" class="logo" />
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
      :doneWebsites="doneWebsites"
      :searching="searching"
      v-if="selectedWebsites.length !== 0"
    />

    <SearchResults v-if="results.length !== 0" :results="results" />
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
      results: [],
      selectedWebsites: [],
      doneWebsites: [],
      searching: false,
    }
  },
  created() {
    runtime.EventsOn('websiteDone', (result) => {
      this.results.push(result)
      this.doneWebsites.push(result.Website)
    })
  },
  methods: {
    updateWebsites() {
      let categories = this.$refs.categories.getSelectedCategories
      window['go']['main']['App']
        ['GetWebsitesWithCategories'](categories)
        .then((data) => {
          this.selectedWebsites = data ?? []
        })
    },
    search() {
      this.searching = true
      this.results = []
      this.doneWebsites = []

      let selectedWebsites = []
      if (this.$refs.selectedWebsitesComponent) {
        selectedWebsites =
          this.$refs.selectedWebsitesComponent.getSelectedWebsites
      }
      if (selectedWebsites.length === 0) {
        this.$q.notify(this.$t('notifications.choose_a_category'))
      }
      window['go']['main']['App']
        ['Search'](this.input, selectedWebsites)
        .then(() => {
          this.searching = false
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
  .logo {
    margin-top: -10px;
    width: 150px;
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
