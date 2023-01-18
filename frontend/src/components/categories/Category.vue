<template>
  <div class="category">
    <div class="category-name">
      <q-checkbox
        v-model="selectedCategory"
        class="checkbox"
        @click="handleCategorySelection"
      />
      {{ category }}
    </div>
    <div class="subcategories">
      <div
        class="subcategory"
        v-for="subcategory in subcategories"
        :key="subcategory"
      >
        <q-checkbox
          v-model="selectedSubcategories[subcategory]"
          class="checkbox"
          @click="handleSubcategorySelection"
        />
        {{ subcategory }}
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'CategorySelector',
  data() {
    return {
      selectedSubcategories: {},
      selectedCategory: false,
    }
  },
  props: {
    subcategories: { type: Array },
    category: { type: String },
  },
  methods: {
    handleCategorySelection() {
      console.log(this.selectedSubcategories)
      Object.entries(this.selectedSubcategories).forEach((object) => {
        this.selectedSubcategories[object[0]] = true
      })
    },
    handleSubcategorySelection() {
      let allSelected = true
      Object.entries(this.selectedSubcategories).forEach((object) => {
        if (!this.selectedSubcategories[object[0]]) allSelected = false
      })
      this.selectedCategory = allSelected
    },
  },
  created() {
    this.subcategories.forEach((c) => {
      this.selectedSubcategories[c] = false
    })
  },
})
</script>
<style lang="scss" scoped>
.category {
  border: 2px solid black;
  padding: 7px;
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
