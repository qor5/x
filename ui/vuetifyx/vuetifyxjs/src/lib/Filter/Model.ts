export interface SelectOption {
  text: string
  value: string
}

export interface IndependentTranslations {
  filterBy: string
}

export interface linkageSelectData {
  items: Array<any>
  labels: Array<string>
  selectOutOfOrder: Boolean
}

export interface FilterItem {
  key: string
  label: string
  folded: boolean
  itemType: string
  modifier: string
  valueIs: string
  valuesAre: string[]
  selected?: boolean
  valueFrom?: string
  valueTo?: string
  inTheLastValue?: string
  inTheLastUnit?: string
  options?: SelectOption[]
  translations?: IndependentTranslations
  linkageSelectData?: linkageSelectData
}
