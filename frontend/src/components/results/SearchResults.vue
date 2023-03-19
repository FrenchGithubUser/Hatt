<template>
  <div class="results shadow-6">
    <div class="search-bar-container">
      <q-input
        bg-color="blue-grey-1"
        class="search-bar"
        v-model="input"
        :label="$t('results.filter_results')"
        @keyup.enter="filterResults"
        outlined
      >
        <template v-slot:append>
          <q-icon class="cursor-pointer" name="search" @click="filterResults" />
        </template>
      </q-input>
    </div>
    <WebsiteResults
      class="result"
      v-for="source in filteredResults"
      :key="source.Website"
      :result="source"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import WebsiteResults from 'src/components/results/WebsiteResults.vue'

export default defineComponent({
  name: 'SearchResults',
  components: { WebsiteResults },
  data() {
    return {
      filteredResults: [],
      input: '',
    }
  },
  props: {
    results: { type: Array },
  },
  created() {
    this.filterResults()
  },
  methods: {
    filterResults() {
      this.filteredResults = JSON.parse(JSON.stringify(this.results))

      let matchingInput = false
      let formattedInput = this.input.toLowerCase()
      let matchingItems = []

      this.filteredResults.forEach((website, websiteIndex) => {
        matchingItems = []

        if (website.Items !== null && website.Items.length !== 0) {
          website.Items.forEach((item, itemIndex) => {
            matchingInput = false

            if (item.Name.toLowerCase().indexOf(formattedInput) >= 0) {
              matchingInput = true
            }

            if (item.Metadata) {
              Object.values(item.Metadata).forEach((metadata) => {
                if (metadata.toLowerCase().indexOf(formattedInput) >= 0) {
                  matchingInput = true
                }
              })
            }

            if (matchingInput) {
              // this.filteredResults[websiteIndex].Items.splice(itemIndex, 1)
              matchingItems.push(item)
            }
          })
        }

        if (matchingItems.length === 0) {
          this.filteredResults[websiteIndex].Items = null
        } else {
          this.filteredResults[websiteIndex].Items = [...matchingItems]
        }
      })

      this.filteredResults.sort((a, b) => {
        return (b.Items ? b.Items.length : 0) - (a.Items ? a.Items.length : 0)
      })
    },
  },
  watch: {
    results: {
      handler: function () {
        this.filterResults()
      },
      deep: true,
    },
    // input() {
    //   this.filterResults()
    // },
  },
})
</script>
<style lang="scss" scoped>
.results {
  width: 97vw;
  margin-top: 15px;
  padding: 15px;
  border-radius: 15px;
  margin-top: 30px;
  margin-bottom: 20px;
  .search-bar-container {
    display: flex;
    justify-content: center;
    width: 100%;
    .search-bar {
      width: 25%;
    }
  }
  .result {
    margin-bottom: 15px;
  }
}
</style>
