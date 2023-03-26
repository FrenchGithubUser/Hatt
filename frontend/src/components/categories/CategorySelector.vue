<template>
  <div class="category-selector shadow-3">
    <div class="categories">
      <div class="title">
        {{ $t('categories.categories') }}
      </div>
      <div class="categories-list">
        <Category
          v-for="category in categories"
          :key="category"
          :category="category"
          ref="categories"
          @selection-updated="$emit('selection-updated')"
        />
      </div>
    </div>
    <div class="custom-lists" v-if="customLists.length !== 0">
      <div class="title">
        {{ $t('settings.custom_lists') }}
      </div>
      <div class="lists">
        <div class="list shadow-3" v-for="list in customLists" :key="list.name">
          <q-checkbox
            v-model="selectedCustomLists"
            :val="list.name"
            class="checkbox"
          >
            <div class="list-name">{{ list.name }}</div>
          </q-checkbox>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import Category from './Category.vue'

export default defineComponent({
  name: 'CategorySelector',
  components: { Category },
  props: {},
  data() {
    return {
      categories: [
        'ebooks',
        'audio_books',
        'courses',
        'tv_shows',
        'movies',
        'console_games',
        'pc_games',
        'pc_software',
        'mobile_apps',
        'music',
        'mainstream',
      ],
      customLists: [],
      selectedCustomLists: [],
      // categories: {
      //   ebooks: ['Magazines', 'Manga'],
      //   tv_shows: ['Anime', 'Netflix-like', 'Cartoons'],
      //   games: ['Switch', 'Wii', 'Nintendo ds', 'Nintendo 3ds', 'PC'],
      //   movies: ['Cartoons', 'Rare-Old', 'Mainstream'],
      //   software: ['PC software', 'Android APKs'],
      //   mainstream: [],
      // },
    }
  },
  created() {
    window['go']['main']['App']['ReadCustomLists']().then((data) => {
      this.customLists = data ?? []
    })
  },
  computed: {
    getSelectedCategories() {
      let categories = []
      this.$refs.categories.forEach((category) => {
        if (category.isSelected) {
          categories.push(category.category)
        }
      })
      return categories
    },
    getSelectedCustomLists() {
      let selectedLists = []
      this.customLists.forEach((list) => {
        if (this.selectedCustomLists.indexOf(list.name) >= 0) {
          selectedLists.push(list)
        }
      })
      return selectedLists
    },
  },
  watch: {
    selectedCustomLists() {
      this.$emit('selection-updated')
    },
  },
})
</script>
<style lang="scss" scoped>
.category-selector {
  max-width: 80%;
  border-radius: 15px;
  padding: 10px;
  .title {
    font-weight: bold;
  }
  .categories {
    .categories-list {
      display: flex;
      justify-content: center;
      flex-wrap: wrap;
    }
  }
  .custom-lists {
    .lists {
      display: flex;
      justify-content: center;
      flex-wrap: wrap;
      .list {
        color: var(--q-primary);
        font-weight: bold;
        padding: 5px;
        border-radius: 10px;
        margin: 10px;
      }
    }
  }
}
</style>
