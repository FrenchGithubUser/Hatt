<template>
  <div class="category-selector shadow-3">
    <div class="categories">
      <Category
        v-for="(subcategories, category) in categories"
        :key="category"
        :category="category"
        :subcategories="subcategories"
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
        TV_Shows: ['Anime', 'Netflix_like'],
        Games: ['switch', 'Wii', 'Nintendo_ds', 'Nintendo_3ds', 'PC'],
        Movies: ['Cartoons', 'Rare', 'Others'],
        Software: ['PC software', 'Android APKs'],
      },
    }
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
