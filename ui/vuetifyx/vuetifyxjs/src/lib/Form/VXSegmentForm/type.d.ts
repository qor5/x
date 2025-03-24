export type ConditionType = 'intersect' | 'union'

export type TagType = {
  tag: {
    builderID: string
    params: Record<string, any>
    values: string[]
  }
}

export type ConditionItemType =
  | TagType
  | {
      [key in ConditionType]?: ConditionItemType[]
    }

export type SavedFormType = {
  [key in ConditionType]?: ConditionItemType[]
}

export type OptionsType = {
  id: string
  name: string
  description: string
  builders: BuilderType[]
}

export type BuilderType = {
  id: string
  name: string
  description: string
  categoryID: string
  view: ViewType
}

export type ViewType = {
  fragments: FragmentType[]
}

export type FragmentType = {
  defaultValue: string
  key: string
  multiple: boolean
  options: OptionType[]
  required: boolean
  skipIf: null | SkipType
  skipUnless: null | SkipType
  type: 'SELECT' | 'DATE_PICKER' | 'NUMBER' | 'NUMBER_INPUT' | 'TEXT'
  validation: null | string
}

export type SkipType<K extends keyof FragmentType = keyof FragmentType> = {
  [`$${string & K}`]: {
    [key: string]: string[]
  }
}

export type OptionType = {
  label: string
  value: string
}
