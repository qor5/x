# VXSegmentForm

这是一个用于构建分段表单的组件，可以用于创建复杂的条件筛选器。

## API

### Props

| 名称       | 介绍                   | 类型     | 默认值 |
| ---------- | ---------------------- | -------- | ------ |
| modelValue | 表单的值，用于双向绑定 | `Object` | `{}`   |
| options    | 分段表单的选项配置     | `Array`  | `[]`   |
| readonly   | 是否为只读模式         | `Boolean`| `false`|

### Methods

| 名称            | 返回类型  | 介绍                               |
| --------------- | --------- | ---------------------------------- |
| validate        | `Boolean` | 验证表单项是否有效，并显示错误信息 |
| resetValidation | `void`    | 重置验证状态，清除错误信息         |
| isValid         | `Boolean` | 验证表单是否有效，返回验证结果     |

> 目前只校验 第一项 select，其余后端校验

### Events

| 名称              | 载荷     | 介绍               |
| ----------------- | -------- | ------------------ |
| update:modelValue | `Object` | 当表单值更新时触发 |

## 基本用法

VXSegmentForm 组件允许用户通过添加规则来构建复杂的条件筛选逻辑。

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'

const modelValue = ref({})
const formRef = ref(null)
const options = ref([
  {
    id: 'demographics',
    name: 'Demographics',
    description: 'Demographic filters',
    builders: [
      {
        id: 'user_gender',
        name: 'User Gender',
        description: 'Filter users by gender',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: 'EQ',
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'equals',
                  value: 'EQ'
                },
                {
                  label: 'not equals',
                  value: 'NE'
                },
                {
                  label: 'in',
                  value: 'IN'
                },
                {
                  label: 'not in',
                  value: 'NOT_IN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              multiple: false,
              options: [
                {
                  label: 'Male',
                  value: 'MALE'
                },
                {
                  label: 'Female',
                  value: 'FEMALE'
                },
                {
                  label: 'Other',
                  value: 'OTHER'
                }
              ],
              required: true,
              skipIf: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'values',
              multiple: true,
              options: [
                {
                  label: 'Male',
                  value: 'MALE'
                },
                {
                  label: 'Female',
                  value: 'FEMALE'
                },
                {
                  label: 'Other',
                  value: 'OTHER'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'user_age',
        name: 'User Age',
        description: 'Filter users by age range',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'Equals',
                  value: 'EQ'
                },
                {
                  label: 'Not Equals',
                  value: 'NE'
                },
                {
                  label: 'Less Than',
                  value: 'LT'
                },
                {
                  label: 'Less Than or Equals',
                  value: 'LTE'
                },
                {
                  label: 'Greater Than',
                  value: 'GT'
                },
                {
                  label: 'Greater Than or Equals',
                  value: 'GTE'
                },
                {
                  label: 'Between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              max: 120,
              min: 0,
              required: true,
              skipIf: {
                operator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'min',
              max: 120,
              min: 0,
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'max',
              max: 120,
              min: 0,
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'user_city',
        name: 'User City',
        description: 'Filter users by city',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: 'EQ',
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'equals',
                  value: 'EQ'
                },
                {
                  label: 'not equals',
                  value: 'NE'
                },
                {
                  label: 'in',
                  value: 'IN'
                },
                {
                  label: 'not in',
                  value: 'NOT_IN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              multiple: false,
              options: [
                {
                  label: 'Tokyo',
                  value: 'TOKYO'
                },
                {
                  label: 'Osaka',
                  value: 'OSAKA'
                },
                {
                  label: 'Kyoto',
                  value: 'KYOTO'
                },
                {
                  label: 'Sapporo',
                  value: 'SAPPORO'
                },
                {
                  label: 'Yokohama',
                  value: 'YOKOHAMA'
                },
                {
                  label: 'Nagoya',
                  value: 'NAGOYA'
                },
                {
                  label: 'Fukuoka',
                  value: 'FUKUOKA'
                },
                {
                  label: 'Hiroshima',
                  value: 'HIROSHIMA'
                }
              ],
              required: true,
              skipIf: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'values',
              multiple: true,
              options: [
                {
                  label: 'Tokyo',
                  value: 'TOKYO'
                },
                {
                  label: 'Osaka',
                  value: 'OSAKA'
                },
                {
                  label: 'Kyoto',
                  value: 'KYOTO'
                },
                {
                  label: 'Sapporo',
                  value: 'SAPPORO'
                },
                {
                  label: 'Yokohama',
                  value: 'YOKOHAMA'
                },
                {
                  label: 'Nagoya',
                  value: 'NAGOYA'
                },
                {
                  label: 'Fukuoka',
                  value: 'FUKUOKA'
                },
                {
                  label: 'Hiroshima',
                  value: 'HIROSHIMA'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'user_signup_source',
        name: 'User Signup Source',
        description: 'Filter users by signup source',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: 'EQ',
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'equals',
                  value: 'EQ'
                },
                {
                  label: 'not equals',
                  value: 'NE'
                },
                {
                  label: 'in',
                  value: 'IN'
                },
                {
                  label: 'not in',
                  value: 'NOT_IN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              multiple: false,
              options: [
                {
                  label: 'Website',
                  value: 'WEBSITE'
                },
                {
                  label: 'Mobile App',
                  value: 'MOBILE_APP'
                },
                {
                  label: 'Referral',
                  value: 'REFERRAL'
                },
                {
                  label: 'Advertisement',
                  value: 'ADVERTISEMENT'
                }
              ],
              required: true,
              skipIf: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'values',
              multiple: true,
              options: [
                {
                  label: 'Website',
                  value: 'WEBSITE'
                },
                {
                  label: 'Mobile App',
                  value: 'MOBILE_APP'
                },
                {
                  label: 'Referral',
                  value: 'REFERRAL'
                },
                {
                  label: 'Advertisement',
                  value: 'ADVERTISEMENT'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              type: 'SELECT',
              validation: null
            }
          ]
        }
      }
    ]
  },
  {
    id: 'activities',
    name: 'Activities',
    description: 'User activity filters',
    builders: [
      {
        id: 'user_last_active',
        name: 'User Last Active',
        description: 'Filter users by last active time range',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Between',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: null,
              includeTime: true,
              key: 'start',
              max: '0001-01-01T00:00:00Z',
              min: '0001-01-01T00:00:00Z',
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'DATE_PICKER',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'and',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: null,
              includeTime: true,
              key: 'end',
              max: '0001-01-01T00:00:00Z',
              min: '0001-01-01T00:00:00Z',
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'DATE_PICKER',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_login',
        name: 'logged in Events',
        description: 'Filter users by logged in events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'logged in',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_view_pdp',
        name: 'viewed products Events',
        description: 'Filter users by viewed products events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'viewed products',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_add_to_cart',
        name: 'added to cart Events',
        description: 'Filter users by added to cart events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'added to cart',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_begin_checkout',
        name: 'began checkout Events',
        description: 'Filter users by began checkout events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'began checkout',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_confirm',
        name: 'confirmed orders Events',
        description: 'Filter users by confirmed orders events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'confirmed orders',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'event_purchase',
        name: 'made purchases Events',
        description: 'Filter users by made purchases events in a time period',
        categoryID: 'activities',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'Users who',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'COUNT',
              key: 'accumulation',
              multiple: false,
              options: [
                {
                  label: 'total occurrences',
                  value: 'COUNT'
                },
                {
                  label: 'unique days',
                  value: 'DAYS'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'made purchases',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: 'GTE',
              key: 'countOperator',
              multiple: false,
              options: [
                {
                  label: 'exactly',
                  value: 'EQ'
                },
                {
                  label: 'not exactly',
                  value: 'NE'
                },
                {
                  label: 'less than',
                  value: 'LT'
                },
                {
                  label: 'at most',
                  value: 'LTE'
                },
                {
                  label: 'more than',
                  value: 'GT'
                },
                {
                  label: 'at least',
                  value: 'GTE'
                },
                {
                  label: 'between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: 1,
              key: 'countValue',
              max: 1000,
              min: 1,
              required: true,
              skipIf: {
                countOperator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMin',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'countMax',
              max: 1000,
              min: 1,
              required: true,
              skipIf: null,
              skipUnless: {
                $countOperator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: '',
              required: false,
              skipIf: null,
              skipUnless: null,
              text: 'times in the last',
              type: 'TEXT',
              validation: null
            },
            {
              defaultValue: '30D',
              key: 'timeRange',
              multiple: false,
              options: [
                {
                  label: 'past 7 days',
                  value: '7D'
                },
                {
                  label: 'past 10 days',
                  value: '10D'
                },
                {
                  label: 'past 30 days',
                  value: '30D'
                },
                {
                  label: 'past 90 days',
                  value: '90D'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            }
          ]
        }
      }
    ]
  }
])

const reset = () => {
  modelValue.value = []
  if (formRef.value) {
    formRef.value.resetValidation()
  }
}

const save = () => {
  if (formRef.value && formRef.value.validate()) {
    // 这里可以添加保存逻辑
    console.log('表单验证通过，可以保存数据', modelValue.value)
  } else {
    console.log('表单验证失败')
  }
}

const clearValidation = () => {
  if (formRef.value) {
    formRef.value.resetValidation()
  }
}
</script>

<template>
  <vx-segment-form ref="formRef" v-model="modelValue" :options="options" :readonly="readonly" />
  <div class="text-right mt-4">
    <div class="d-flex justify-end">
      <vx-btn class="mr-2" color="grey" @click="reset">Reset</vx-btn>
      <vx-btn class="mr-2" color="secondary" @click="clearValidation">Clear Validation</vx-btn>
      <vx-btn @click="save">Save</vx-btn>
    </div>
  </div>
  <VueJsonPretty :data="modelValue" />
  <br />
  <vx-dialog title="Create New Segment" width="840" okText="Save">
    <vx-field label="Title" v-model="segmentName" />

    <div style="min-height: 752px">
      <vx-label class="mb-4">Conditions</vx-label>
      <vx-segment-form ref="dialogFormRef" v-model="modelValue" :options="options" />
    </div>

    <template v-slot:activator="{ props: { activatorProps } }">
      <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
    </template>
  </vx-dialog>
</template>
```

<style scoped></style>

## 只读模式

VXSegmentForm 组件支持只读模式，适用于查看已创建的条件而不允许编辑。

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'

// 预设一个带有条件的表单值
const readonlyModelValue = ref({
  intersect: [
    {
      union: [
        {
          tag: {
            builderID: "user_gender",
            params: {
              operator: "EQ",
              value: "MALE"
            }
          }
        }
      ]
    },
    {
      union: [
        {
          tag: {
            builderID: "user_age",
            params: {
              operator: "BETWEEN",
              min: 18,
              max: 35
            }
          }
        }
      ]
    }
  ]
})

// 控制只读状态的变量
const isReadonly = ref(true)

// 使用与前面示例相同的选项
const options = ref([
  {
    id: 'demographics',
    name: 'Demographics',
    description: 'Demographic filters',
    builders: [
      {
        id: 'user_gender',
        name: 'User Gender',
        description: 'Filter users by gender',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: 'EQ',
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'equals',
                  value: 'EQ'
                },
                {
                  label: 'not equals',
                  value: 'NE'
                },
                {
                  label: 'in',
                  value: 'IN'
                },
                {
                  label: 'not in',
                  value: 'NOT_IN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              multiple: false,
              options: [
                {
                  label: 'Male',
                  value: 'MALE'
                },
                {
                  label: 'Female',
                  value: 'FEMALE'
                },
                {
                  label: 'Other',
                  value: 'OTHER'
                }
              ],
              required: true,
              skipIf: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'values',
              multiple: true,
              options: [
                {
                  label: 'Male',
                  value: 'MALE'
                },
                {
                  label: 'Female',
                  value: 'FEMALE'
                },
                {
                  label: 'Other',
                  value: 'OTHER'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['IN', 'NOT_IN']
                }
              },
              type: 'SELECT',
              validation: null
            }
          ]
        }
      },
      {
        id: 'user_age',
        name: 'User Age',
        description: 'Filter users by age range',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: null,
              key: 'operator',
              multiple: false,
              options: [
                {
                  label: 'Equals',
                  value: 'EQ'
                },
                {
                  label: 'Not Equals',
                  value: 'NE'
                },
                {
                  label: 'Less Than',
                  value: 'LT'
                },
                {
                  label: 'Less Than or Equals',
                  value: 'LTE'
                },
                {
                  label: 'Greater Than',
                  value: 'GT'
                },
                {
                  label: 'Greater Than or Equals',
                  value: 'GTE'
                },
                {
                  label: 'Between',
                  value: 'BETWEEN'
                }
              ],
              required: true,
              skipIf: null,
              skipUnless: null,
              type: 'SELECT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'value',
              max: 120,
              min: 0,
              required: true,
              skipIf: {
                operator: 'BETWEEN'
              },
              skipUnless: null,
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'min',
              max: 120,
              min: 0,
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            },
            {
              defaultValue: null,
              key: 'max',
              max: 120,
              min: 0,
              required: true,
              skipIf: null,
              skipUnless: {
                $operator: {
                  IN: ['BETWEEN']
                }
              },
              type: 'NUMBER_INPUT',
              validation: null
            }
          ]
        }
      }
    ]
  }
])

const toggleReadonlyMode = () => {
  isReadonly.value = !isReadonly.value;
}
</script>

<template>
  <h3>只读模式演示</h3>
  <vx-segment-form 
    v-model="readonlyModelValue" 
    :options="options" 
    :readonly="isReadonly" 
  />
  
  <div class="text-right mt-4">
    <div class="d-flex justify-end">
      <vx-btn @click="toggleReadonlyMode">
        {{ isReadonly ? '切换到编辑模式' : '切换到只读模式' }}
      </vx-btn>
    </div>
  </div>
  
  <div class="mt-4">
    <p>当前模式: <b>{{ isReadonly ? '只读' : '编辑' }}</b></p>
    <p>表单数据:</p>
    <VueJsonPretty :data="readonlyModelValue" />
  </div>
</template>
```

<style scoped></style>

:::

## API更新

### Props

| 名称       | 介绍                   | 类型      | 默认值 |
| ---------- | ---------------------- | --------- | ------ |
| modelValue | 表单的值，用于双向绑定 | `Object`  | `{}`   |
| options    | 分段表单的选项配置     | `Array`   | `[]`   |
| readonly   | 是否为只读模式         | `Boolean` | `false`|

## 在Go代码中使用

在Go代码中，您可以使用`Readonly`方法来设置表单的只读状态：

```go
VXSegmentForm("").
    Options(options).
    Readonly(true).  // 设置为只读状态
    Bind("model-value", "formData")
```
