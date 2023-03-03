<template>
  <div class="category">
    <div class="category-name">
      <q-checkbox
        v-model="selectedCategory"
        class="checkbox"
        @click="handleCategorySelection"
      />
      {{ $t('categories.' + category) }}
    </div>
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
  border: 2px solid $primary;
  padding: 4px;
  padding-left: 0px;
  border-radius: 15px;
  margin: 5px;
  .checkbox {
    margin-right: -7px;
  }
  .category-name {
    font-weight: bold;
    color: $primary;
    text-align: center;
  }
  .subcategories {
    display: flex;
  }
}
</style>
