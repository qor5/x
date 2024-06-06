<script setup lang="ts">
import { Ref, ref } from 'vue'
import LinkageSelect from '@/lib/LinkageSelect.vue'

const value = ref(['2', '3', '6'])
const items: Ref<any> = ref([
  [
    {
      ID: '1',
      Name: '浙江',
      ChildrenIDs: ['1', '2']
    },
    {
      ID: '2',
      Name: '江苏',
      ChildrenIDs: ['3', '4']
    }
  ],
  [
    { ID: '1', Name: '杭州', ChildrenIDs: ['1', '2'] },
    { ID: '2', Name: '宁波', ChildrenIDs: ['3', '4'] },
    { ID: '3', Name: '南京', ChildrenIDs: ['5', '6'] },
    { ID: '4', Name: '苏州', ChildrenIDs: ['7', '8'] }
  ],
  [
    { ID: '1', Name: '拱墅区' },
    { ID: '2', Name: '西湖区' },
    { ID: '3', Name: '镇海区' },
    { ID: '4', Name: '鄞州区' },
    { ID: '5', Name: '鼓楼区' },
    { ID: '6', Name: '玄武区' },
    { ID: '7', Name: '常熟区' },
    { ID: '8', Name: '吴江区' }
  ]
])
const labels = ref(['Province', 'City', 'District'])
</script>

<template>
  <p>{{ value }}</p>
  <linkage-select v-model="value" :items="items" :labels="labels"></linkage-select>
</template>

<style scoped></style>
