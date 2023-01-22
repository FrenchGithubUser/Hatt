<template>
  <div class="category-selector shadow-3">
    <div class="categories">
      <Category
        v-for="(subcategories, category) in categories"
        :key="category"
        :category="category"
        :subcategories="subcategories"
        ref="categories"
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
      categories: {
        Ebooks: ['Magazines', 'Manga'],
        'TV shows': ['Anime', 'Netflix-like', 'Cartoons'],
        Games: ['Switch', 'Wii', 'Nintendo ds', 'Nintendo 3ds', 'PC'],
        Movies: ['Cartoons', 'Rare-Old', 'Mainstream'],
        Software: ['PC software', 'Android APKs'],
      },
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
  max-width: 50%;
  .categories {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
  }
}
</style>
