<template>
  <div class="category-selector shadow-1">
    <div class="categories">
      <Category
        v-for="category in categories"
        :key="category"
        :category="category"
        ref="categories"
        @selection-updated="$emit('selection-updated')"
      />
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
        'tv_shows',
        'movies',
        'console_games',
        'pc_games',
        'mobile_apps',
        'pc_software',
        'music',
        'mainstream',
      ],
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
  computed: {
    getSelectedCategories() {
      let categories = []
      this.$refs.categories.forEach((category) => {
        if (category.isSelected) {
          categories.push(category.category)
        }
      })
      return { categories: categories }
    },
  },
})
</script>
<style lang="scss" scoped>
.category-selector {
  padding: 10px;
  border-radius: 15px;
  max-width: 80%;
  .categories {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
  }
}
</style>
