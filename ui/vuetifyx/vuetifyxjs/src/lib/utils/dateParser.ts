import dayjs from 'dayjs'
import customParseFormat from 'dayjs/plugin/customParseFormat'

// 启用自定义格式解析插件
dayjs.extend(customParseFormat)

/**
 * 增强的日期解析器，支持更多日期格式，特别是中文年月日格式
 */
export class EnhancedDateParser {
  // 支持的日期格式列表
  private static readonly DATE_FORMATS = [
    // 标准格式
    'YYYY-MM-DD',
    'YYYY/MM/DD',
    'YYYY.MM.DD',
    'YYYY-M-D',
    'YYYY/M/D',
    'YYYY.M.D',

    // 带时间的格式
    'YYYY-MM-DD HH:mm:ss',
    'YYYY-MM-DD HH:mm',
    'YYYY/MM/DD HH:mm:ss',
    'YYYY/MM/DD HH:mm',
    'YYYY.MM.DD HH:mm:ss',
    'YYYY.MM.DD HH:mm',

    // 中文格式
    'YYYY年MM月DD日',
    'YYYY年M月D日',
    'YYYY年MM月DD日 HH:mm:ss',
    'YYYY年MM月DD日 HH:mm',
    'YYYY年M月D日 HH:mm:ss',
    'YYYY年M月D日 HH:mm',

    // 短格式
    'MM-DD',
    'MM/DD',
    'MM.DD',
    'M-D',
    'M/D',
    'M.D',
    'MM月DD日',
    'M月D日',

    // 其他常见格式
    'DD-MM-YYYY',
    'DD/MM/YYYY',
    'DD.MM.YYYY',
    'DD-M-YYYY',
    'DD/M/YYYY',
    'DD.M.YYYY'
  ]

  // 中文数字映射
  private static readonly CHINESE_NUMBERS = {
    一: '1',
    二: '2',
    三: '3',
    四: '4',
    五: '5',
    六: '6',
    七: '7',
    八: '8',
    九: '9',
    十: '10',
    十一: '11',
    十二: '12',
    十三: '13',
    十四: '14',
    十五: '15',
    十六: '16',
    十七: '17',
    十八: '18',
    十九: '19',
    二十: '20',
    二十一: '21',
    二十二: '22',
    二十三: '23',
    二十四: '24',
    二十五: '25',
    二十六: '26',
    二十七: '27',
    二十八: '28',
    二十九: '29',
    三十: '30',
    三十一: '31'
  }

  /**
   * 预处理日期字符串，处理中文数字和特殊格式
   */
  private static preprocessDateString(dateStr: string): string {
    if (!dateStr || typeof dateStr !== 'string') {
      return dateStr
    }

    let processed = dateStr.trim()

    // 替换中文数字
    Object.entries(this.CHINESE_NUMBERS).forEach(([chinese, arabic]) => {
      processed = processed.replace(new RegExp(chinese, 'g'), arabic)
    })

    // 处理特殊的中文格式
    // 处理"今天"、"明天"、"昨天"
    const today = dayjs()
    if (processed.includes('今天') || processed === '今日') {
      return today.format('YYYY-MM-DD')
    }
    if (processed.includes('明天') || processed === '明日') {
      return today.add(1, 'day').format('YYYY-MM-DD')
    }
    if (processed.includes('昨天') || processed === '昨日') {
      return today.subtract(1, 'day').format('YYYY-MM-DD')
    }

    // 处理"大后天"、"后天"、"前天"、"大前天"
    if (processed.includes('大后天')) {
      return today.add(3, 'day').format('YYYY-MM-DD')
    }
    if (processed.includes('后天')) {
      return today.add(2, 'day').format('YYYY-MM-DD')
    }
    if (processed.includes('前天')) {
      return today.subtract(2, 'day').format('YYYY-MM-DD')
    }
    if (processed.includes('大前天')) {
      return today.subtract(3, 'day').format('YYYY-MM-DD')
    }

    // 处理只有月日的情况，自动补充当前年份
    const monthDayRegex = /^(\d{1,2})[月\/\-\.](\d{1,2})[日]?$/
    const monthDayMatch = processed.match(monthDayRegex)
    if (monthDayMatch) {
      const currentYear = today.year()
      const month = monthDayMatch[1].padStart(2, '0')
      const day = monthDayMatch[2].padStart(2, '0')
      processed = `${currentYear}-${month}-${day}`
    }

    // 处理年份简写（如：23年12月25日 -> 2023年12月25日）
    const shortYearRegex = /^(\d{2})年/
    const shortYearMatch = processed.match(shortYearRegex)
    if (shortYearMatch) {
      const shortYear = parseInt(shortYearMatch[1])
      const currentYear = today.year()
      const currentCentury = Math.floor(currentYear / 100) * 100
      // 如果简写年份小于当前年份的后两位，认为是当前世纪，否则是上个世纪
      const fullYear =
        shortYear <= currentYear % 100
          ? currentCentury + shortYear
          : currentCentury - 100 + shortYear
      processed = processed.replace(shortYearRegex, `${fullYear}年`)
    }

    return processed
  }

  /**
   * 尝试解析日期字符串，支持多种格式
   */
  static parseDate(dateInput: string | number | Date): dayjs.Dayjs | null {
    if (!dateInput) {
      return null
    }

    // 如果已经是有效的Date对象或时间戳，直接使用dayjs解析
    if (typeof dateInput === 'number' || dateInput instanceof Date) {
      const parsed = dayjs(dateInput)
      return parsed.isValid() ? parsed : null
    }

    // 预处理字符串
    const processedStr = this.preprocessDateString(dateInput)

    // 首先尝试dayjs的默认解析
    let parsed = dayjs(processedStr)
    if (parsed.isValid()) {
      return parsed
    }

    // 尝试各种格式
    for (const format of this.DATE_FORMATS) {
      parsed = dayjs(processedStr, format, true) // strict mode
      if (parsed.isValid()) {
        return parsed
      }
    }

    // 尝试更宽松的解析
    try {
      // 使用正则表达式提取数字
      const numberRegex = /\d+/g
      const numbers = processedStr.match(numberRegex)

      if (numbers && numbers.length >= 2) {
        const [year, month, day, hour, minute, second] = numbers.map((num) => parseInt(num))

        // 智能判断年份
        let actualYear = year
        if (year < 100) {
          const currentYear = dayjs().year()
          const currentCentury = Math.floor(currentYear / 100) * 100
          actualYear =
            year <= currentYear % 100 ? currentCentury + year : currentCentury - 100 + year
        }

        // 构建日期
        if (numbers.length >= 3) {
          // 有年月日
          const dateStr = `${actualYear}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`
          if (numbers.length >= 6) {
            // 有时分秒
            const timeStr = `${hour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')}:${second.toString().padStart(2, '0')}`
            parsed = dayjs(`${dateStr} ${timeStr}`)
          } else if (numbers.length >= 5) {
            // 有时分
            const timeStr = `${hour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')}:00`
            parsed = dayjs(`${dateStr} ${timeStr}`)
          } else {
            parsed = dayjs(dateStr)
          }
        } else {
          // 只有月日，使用当前年份
          const currentYear = dayjs().year()
          const dateStr = `${currentYear}-${year.toString().padStart(2, '0')}-${month.toString().padStart(2, '0')}`
          parsed = dayjs(dateStr)
        }

        if (parsed.isValid()) {
          return parsed
        }
      }
    } catch (error) {
      console.warn('Enhanced date parsing failed:', error)
    }

    return null
  }

  /**
   * 格式化日期
   */
  static formatDate(
    date: dayjs.Dayjs | string | number | Date,
    format: string = 'YYYY-MM-DD'
  ): string {
    if (!date) {
      return ''
    }

    const dayjsDate = dayjs.isDayjs(date) ? date : dayjs(date)
    return dayjsDate.isValid() ? dayjsDate.format(format) : ''
  }

  /**
   * 检查日期是否有效
   */
  static isValidDate(dateInput: string | number | Date): boolean {
    const parsed = this.parseDate(dateInput)
    return parsed !== null && parsed.isValid()
  }

  /**
   * 获取支持的日期格式列表
   */
  static getSupportedFormats(): string[] {
    return [...this.DATE_FORMATS]
  }

  /**
   * 智能解析相对日期（如"3天后"、"2周前"等）
   */
  static parseRelativeDate(dateStr: string): dayjs.Dayjs | null {
    if (!dateStr || typeof dateStr !== 'string') {
      return null
    }

    const str = dateStr.trim()
    const today = dayjs()

    // 匹配相对日期模式
    const patterns = [
      { regex: /(\d+)天[后之]/, unit: 'day', direction: 1 },
      { regex: /(\d+)天[前之]/, unit: 'day', direction: -1 },
      { regex: /(\d+)[个]?月[后之]/, unit: 'month', direction: 1 },
      { regex: /(\d+)[个]?月[前之]/, unit: 'month', direction: -1 },
      { regex: /(\d+)[个]?[周星期][后之]/, unit: 'week', direction: 1 },
      { regex: /(\d+)[个]?[周星期][前之]/, unit: 'week', direction: -1 },
      { regex: /(\d+)[个]?年[后之]/, unit: 'year', direction: 1 },
      { regex: /(\d+)[个]?年[前之]/, unit: 'year', direction: -1 }
    ]

    for (const pattern of patterns) {
      const match = str.match(pattern.regex)
      if (match) {
        const amount = parseInt(match[1]) * pattern.direction
        return today.add(amount, pattern.unit as any)
      }
    }

    return null
  }
}
