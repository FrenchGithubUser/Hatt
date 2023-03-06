<template>
  <div class="category shadow-2">
    <q-checkbox
      v-model="selectedCategory"
      class="checkbox"
      @click="handleCategorySelection"
    >
      <div class="stack">
        <img :src="'images/' + category + '.svg'" class="icon" />
        <div class="category-name">
          {{ $t('categories.' + category) }}
        </div>
      </div>
    </q-checkbox>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'CategorySelector',
  data() {
    return {
      selectedCategory: false,
    }
  },
  props: {
    category: { type: String },
  },
  methods: {
    handleCategorySelection() {
      // Object.entries(this.selectedSubcategories).forEach((object) => {})
      this.$emit('selection-updated')
    },
    // handleSubcategorySelection() {
    //   let allSelected = true
    //   Object.entries(this.selectedSubcategories).forEach((object) => {
    //     if (!this.selectedSubcategories[object[0]]) allSelected = false
    //   })
    //   this.selectedCategory = allSelected
    // },
  },
  computed: {
    isSelected() {
      return this.selectedCategory
    },
    // getSelectedSubcategories() {
    //   return this.selectedSubcategories
    // },
  },
  created() {
    // this.subcategories.forEach((c) => {
    //   this.selectedSubcategories[c] = false
    // })
  },
})
</script>
<style lang="scss" scoped>
.category {
  margin: 7px;
  display: flex;
  align-items: center;
  padding: 10px;
  padding-right: 14px;
  padding-left: 5px;
  border-radius: 7px;
  .stack {
    text-align: center;
    .icon {
      width: 30px;
    }
    .category-name {
      font-weight: bold;
      color: $primary;
      text-align: center;
      margin-top: -2px;
    }
  }
}
</style>
