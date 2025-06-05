# vx-filter 筛选

## 基本用法

:::demo

```vue
<template>
  <!--  <h5 class="text-h5">v-model</h5>-->
  <VueJsonPretty :data="value"></VueJsonPretty>
  <v-divider />
  <!--  <h5 class="text-h5">internalValue</h5>-->
  <!--  <VueJsonPretty :data="internalValue"></VueJsonPretty>-->
  <v-divider />
  <vx-filter v-model="value" :internal-value="internalValue" />
</template>

<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { ref } from 'vue'

const internalValue = ref([
  {
    key: 'TestDatetimeRange',
    label: 'TestDatetimeRange',
    itemType: 'DatetimeRangeItem',
    modifier: 'between',
    selected: false,
    folded: false,
    valueFrom: '',
    valueTo: ''
  },
  {
    key: 'TestDatetimeRangePicker',
    label: 'TestDatetimeRangePicker',
    itemType: 'DatetimeRangePickerItem',
    modifier: 'between',
    selected: false,
    folded: false,
    valueFrom: '',
    valueTo: ''
  },
  {
    key: 'TestDateRange',
    label: 'TestDateRange',
    itemType: 'DateRangeItem',
    modifier: 'between',
    selected: false,
    folded: false,
    valueFrom: '',
    valueTo: ''
  },
  {
    key: 'TestDateRangePicker',
    label: 'TestDateRangePicker',
    itemType: 'DateRangePickerItem',
    selected: false,
    folded: false,
    valueFrom: '',
    valueTo: ''
  },
  {
    key: 'TestDate',
    label: 'TestDate',
    itemType: 'DateItem',
    selected: false,
    folded: false,
    valueIs: ''
  },
  {
    key: 'TestDatePicker',
    label: 'TestDatePicker',
    itemType: 'DatePickerItem',
    selected: false,
    folded: false,
    valueIs: ''
  },
  {
    key: 'TestNumber',
    label: 'TestNumber',
    itemType: 'NumberItem',
    selected: false,
    folded: false,
    valueIs: ''
  },
  {
    key: 'TestStringItem',
    label: 'TestStringItem',
    itemType: 'StringItem',
    selected: false,
    folded: false,
    valueIs: ''
  },
  {
    key: 'TestStringItemNoChoose',
    label: 'TestStringItemNoChoose',
    itemType: 'StringItem',
    selected: false,
    folded: false,
    valueIs: '',
    disableChooseModifier: true
  },
  {
    key: 'TestLinkageSelectItem',
    label: 'TestLinkageSelectItem',
    itemType: 'LinkageSelectItem',
    selected: false,
    folded: false,
    valuesAre: [],
    linkageSelectData: {
      items: [
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
          {ID: '1', Name: '杭州', ChildrenIDs: ['1', '2']},
          {ID: '2', Name: '宁波', ChildrenIDs: ['3', '4']},
          {ID: '3', Name: '南京', ChildrenIDs: ['5', '6']},
          {ID: '4', Name: '苏州', ChildrenIDs: ['7', '8']}
        ],
        [
          {ID: '1', Name: '拱墅区'},
          {ID: '2', Name: '西湖区'},
          {ID: '3', Name: '镇海区'},
          {ID: '4', Name: '鄞州区'},
          {ID: '5', Name: '鼓楼区'},
          {ID: '6', Name: '玄武区'},
          {ID: '7', Name: '常熟区'},
          {ID: '8', Name: '吴江区'}
        ]
      ],
      labels: ['Province', 'City', 'District'],
      selectOutOfOrder: true
    }
  },
  {
    key: 'TestMultipleSelectItem',
    label: 'TestMultipleSelectItem',
    itemType: 'MultipleSelectItem',
    selected: false,
    folded: false,
    valuesAre: [],
    options: [
      {
        text: 'John',
        value: 'John'
      },
      {
        text: 'Jacob',
        value: 'Jacob'
      }
    ]
  },
  {
    key: 'f_company',
    label: 'Company',
    itemType: 'SelectItem',
    selected: true,
    valueIs: '1',
    options: [
      {text: '高节', value: '1'},
      {text: '地界', value: '3'}
    ],
    linkageSelectData: {},
    translations: {filterBy: 'Filter by Company'}
  },
  {
    key: 'f_company_remote',
    label: 'AutoCompleteItemRemote',
    itemType: 'AutoCompleteItem',
    valuesAre: null,
    linkageSelectData: {},
    translations: {filterBy: 'Filter by Company'},
    autocompleteDataSource: {
      remoteUrl: 'http://localhost:7800/examples/api/complete/auto-complete-posts',
      itemTitle: 'title',
      itemValue: 'id',
      totalField: 'total',
      itemsField: 'data',
      isPaging: true,
      pageSize: 2,
      separator: "__"
    }
  },
  {
    key: 'f_linkage_remote',
    label: 'LinkageSelectItemRemote',
    itemType: 'LinkageSelectItemRemote',
    linkageSelectData: {
      labels: ['Province', 'City', 'District'],
      linkageSelectRemoteOptions: {
        remoteUrl: 'http://localhost:7800/examples/api/linkage-select-server',
        itemTitle: 'Name',
        itemValue: 'ID',
        totalField: 'total',
        itemsField: 'data',
        isPaging: true,
        levelStart: 1,
        pageSize: 2,
        separator: "__",
        selectOutOfOrder: true
      }
    },
  }
])
// const internalValue = [{
//   key: 'TestStringItem',
//   label: 'TestStringItem',
//   itemType: 'StringItem',
//   selected: true,
//   folded: true,
//   valueIs: 'active'
// }]
const value = ref()
</script>
```
:::