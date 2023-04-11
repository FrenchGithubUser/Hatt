<template>
  <div class="results shadow-6">
    <div class="search-bar-container">
      <SearchBar
        v-model="input"
        @search="filterResults"
        :label="$t('results.filter_results')"
      >
        <q-icon
          class="cursor-pointer"
          name="help"
          @click="filterExplanation = true"
        />
      </SearchBar>
      <q-select
        class="sort-options"
        outlined
        v-model="sortMethod"
        :options="sortMethods"
        :label="$t('results.sort_by')"
        @update:model-value="filterResults"
      >
        <template v-slot:option="scope">
          <q-item v-bind="scope.itemProps">
            <q-item-section>
              <q-item-label>{{ $t('results.' + scope.opt) }}</q-item-label>
            </q-item-section>
          </q-item>
        </template>
        <template v-slot:selected-item="method">
          {{ $t('results.' + method.opt) }}
        </template>
      </q-select>
    </div>
    <WebsiteResults
      class="result"
      v-for="source in filteredResults"
      :key="source.Website"
      :result="source"
    />
    <q-dialog v-model="filterExplanation">
      <div class="popup">
        {{ $t('results.filter_explanation') }}
      </div>
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import WebsiteResults from 'src/components/results/WebsiteResults.vue'
import SearchBar from '../SearchBar.vue'

export default defineComponent({
  name: 'SearchResults',
  components: { WebsiteResults, SearchBar },
  data() {
    return {
      filteredResults: [],
      input: '',
      filterExplanation: false,
      sortMethod: 'relevance',
      sortMethods: ['results_amount', 'relevance'],
    }
  },
  props: {
    results: { type: Array },
    query: { type: String },
  },
  created() {
    this.filterResults()
  },
  methods: {
    filterResults() {
      let formattedInput = this.input.toLowerCase()
      this.filteredResults = JSON.parse(JSON.stringify(this.results))

      let matchingInput = false
      let matchingItems = []

      this.filteredResults.forEach((website, websiteIndex) => {
        matchingItems = []

        if (formattedInput === '' && website.Items !== null) {
          matchingItems = [...website.Items]
        } else {
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
        }

        if (matchingItems.length === 0) {
          this.filteredResults[websiteIndex].Items = null
        } else {
          this.filteredResults[websiteIndex].Items = [...matchingItems]
        }
      })

      if (this.sortMethod === 'results_amount') {
        this.filteredResults.sort((a, b) => {
          return (b.Items ? b.Items.length : 0) - (a.Items ? a.Items.length : 0)
        })
      } else if (this.sortMethod === 'relevance') {
        let formattedQuery = this.query.toLowerCase()
        this.filteredResults.sort((a, b) => {
          let aTotalScore = 0
          let bTotalScore = 0
          if (a.Items !== null) {
            a.Items.forEach((item) => {
              aTotalScore += this.calculateRelevanceScore(
                item.Name,
                formattedQuery
              )
            })
          }
          if (b.Items !== null) {
            b.Items.forEach((item) => {
              bTotalScore += this.calculateRelevanceScore(
                item.Name,
                formattedQuery
              )
            })
          }
          return (
            bTotalScore / (b.Items ? b.Items.length : 1) -
            aTotalScore / (a.Items ? a.Items.length : 1)
          )
        })
      }
    },
    calculateRelevanceScore(item, query) {
      const itemLower = item.toLowerCase() // convertir l'élément en minuscules pour ignorer la casse
      // const queryLower = query.toLowerCase() // convertir la chaîne de caractères recherchée en minuscules pour ignorer la casse

      const queryKeywords = query.split(/[^\w]+/) // découper la chaîne de caractères recherchée en mots clés

      let score = 0
      let position = itemLower.indexOf(query)

      if (position !== -1) {
        // si la chaîne de caractères recherchée est présente dans l'élément
        score += 100 - position * 10 // ajouter un score basé sur la position de la chaîne de caractères recherchée dans l'élément
      } else {
        // si la chaîne de caractères recherchée n'est pas présente dans l'élément
        // ajouter un score basé sur la présence de sous-chaînes similaires aux mots clés de la recherche
        for (let i = 0; i < queryKeywords.length; i++) {
          if (itemLower.includes(queryKeywords[i])) {
            score += 50
            break
          }
        }
      }

      // ajouter un score basé sur la longueur de l'élément par rapport à la longueur de la chaîne de caractères recherchée
      const lengthRatio = itemLower.length / query.length
      if (lengthRatio >= 0.5 && lengthRatio <= 2) {
        score += 10
      }

      // ajouter un score basé sur la correspondance des majuscules/minuscules
      // let isCaseSensitiveMatch = true
      // for (let i = 0; i < query.length; i++) {
      //   if (query[i] !== item[i] && query[i] === itemLower[i]) {
      //     isCaseSensitiveMatch = false
      //     break
      //   }
      // }
      // if (isCaseSensitiveMatch) {
      //   score += 5
      // }

      return score
    },
  },
  watch: {
    results: {
      handler: function () {
        this.filterResults()
      },
      deep: true,
    },
  },
})
</script>
<style lang="scss" scoped>
.results {
  width: 98vw;
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
    .sort-options {
      margin-left: 10px;
    }
  }
  .result {
    margin-bottom: 15px;
  }
}
</style>
