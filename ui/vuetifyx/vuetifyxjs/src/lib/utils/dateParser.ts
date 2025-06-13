import dayjs from 'dayjs'
import customParseFormat from 'dayjs/plugin/customParseFormat'

// Enable custom format parsing plugin
dayjs.extend(customParseFormat)

/**
 * Enhanced date parser, supports more date formats, especially Chinese year-month-day formats
 */
export class EnhancedDateParser {
  // List of supported date formats
  private static readonly DATE_FORMATS = [
    // Standard formats
    'YYYY-MM-DD',
    'YYYY/MM/DD',
    'YYYY.MM.DD',
    'YYYY-M-D',
    'YYYY/M/D',
    'YYYY.M.D',

    // Formats with time
    'YYYY-MM-DD HH:mm:ss',
    'YYYY-MM-DD HH:mm',
    'YYYY/MM/DD HH:mm:ss',
    'YYYY/MM/DD HH:mm',
    'YYYY.MM.DD HH:mm:ss',
    'YYYY.MM.DD HH:mm',

    // Chinese formats
    'YYYY年MM月DD日',
    'YYYY年M月D日',
    'YYYY年MM月DD日 HH:mm:ss',
    'YYYY年MM月DD日 HH:mm',
    'YYYY年M月D日 HH:mm:ss',
    'YYYY年M月D日 HH:mm',

    // Short formats
    'MM-DD',
    'MM/DD',
    'MM.DD',
    'M-D',
    'M/D',
    'M.D',
    'MM月DD日',
    'M月D日',

    // Other common formats
    'DD-MM-YYYY',
    'DD/MM/YYYY',
    'DD.MM.YYYY',
    'DD-M-YYYY',
    'DD/M/YYYY',
    'DD.M.YYYY'
  ]

  // Chinese number mapping
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
   * Preprocess date string, handle Chinese numbers and special formats
   */
  private static preprocessDateString(dateStr: string): string {
    if (!dateStr || typeof dateStr !== 'string') {
      return dateStr
    }

    let processed = dateStr.trim()

    // Replace Chinese numbers
    Object.entries(this.CHINESE_NUMBERS).forEach(([chinese, arabic]) => {
      processed = processed.replace(new RegExp(chinese, 'g'), arabic)
    })

    // Handle special Chinese formats
    // Handle "today", "tomorrow", "yesterday"
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

    // Handle "the day after tomorrow", "the day before yesterday"
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

    // Handle cases with only month and day, automatically fill in the current year
    const monthDayRegex = /^(\d{1,2})[月\/\-\.](\d{1,2})[日]?$/
    const monthDayMatch = processed.match(monthDayRegex)
    if (monthDayMatch) {
      const currentYear = today.year()
      const month = monthDayMatch[1].padStart(2, '0')
      const day = monthDayMatch[2].padStart(2, '0')
      processed = `${currentYear}-${month}-${day}`
    }

    // Handle short year (e.g. 23年12月25日 -> 2023年12月25日)
    const shortYearRegex = /^(\d{2})年/
    const shortYearMatch = processed.match(shortYearRegex)
    if (shortYearMatch) {
      const shortYear = parseInt(shortYearMatch[1])
      const currentYear = today.year()
      const currentCentury = Math.floor(currentYear / 100) * 100
      // If the short year is less than the last two digits of the current year, it is considered this century, otherwise last century
      const fullYear =
        shortYear <= currentYear % 100
          ? currentCentury + shortYear
          : currentCentury - 100 + shortYear
      processed = processed.replace(shortYearRegex, `${fullYear}年`)
    }

    return processed
  }

  /**
   * Try to parse date string, support multiple formats
   */
  static parseDate(dateInput: string | number | Date): dayjs.Dayjs | null {
    if (!dateInput) {
      return null
    }

    // If already a valid Date object or timestamp, use dayjs to parse directly
    if (typeof dateInput === 'number' || dateInput instanceof Date) {
      const parsed = dayjs(dateInput)
      return parsed.isValid() ? parsed : null
    }

    // Preprocess string
    const processedStr = this.preprocessDateString(dateInput)

    // First try dayjs default parsing
    let parsed = dayjs(processedStr)
    if (parsed.isValid()) {
      return parsed
    }

    // Try various formats
    for (const format of this.DATE_FORMATS) {
      parsed = dayjs(processedStr, format, true) // strict mode
      if (parsed.isValid()) {
        return parsed
      }
    }

    // Try more lenient parsing
    try {
      // Use regex to extract numbers
      const numberRegex = /\d+/g
      const numbers = processedStr.match(numberRegex)

      if (numbers && numbers.length >= 2) {
        const [year, month, day, hour, minute, second] = numbers.map((num) => parseInt(num))

        // Smartly determine year
        let actualYear = year
        if (year < 100) {
          const currentYear = dayjs().year()
          const currentCentury = Math.floor(currentYear / 100) * 100
          actualYear =
            year <= currentYear % 100 ? currentCentury + year : currentCentury - 100 + year
        }

        // Build date
        if (numbers.length >= 3) {
          // Has year, month, day
          const dateStr = `${actualYear}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`
          if (numbers.length >= 6) {
            // Has hour, minute, second
            const timeStr = `${hour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')}:${second.toString().padStart(2, '0')}`
            parsed = dayjs(`${dateStr} ${timeStr}`)
          } else if (numbers.length >= 5) {
            // Has hour, minute
            const timeStr = `${hour.toString().padStart(2, '0')}:${minute.toString().padStart(2, '0')}:00`
            parsed = dayjs(`${dateStr} ${timeStr}`)
          } else {
            parsed = dayjs(dateStr)
          }
        } else {
          // Only month and day, use current year
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
   * Format date
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
   * Check if date is valid
   */
  static isValidDate(dateInput: string | number | Date): boolean {
    const parsed = this.parseDate(dateInput)
    return parsed !== null && parsed.isValid()
  }

  /**
   * Get list of supported date formats
   */
  static getSupportedFormats(): string[] {
    return [...this.DATE_FORMATS]
  }

  /**
   * Intelligently parse relative dates (e.g. "3 days later", "2 weeks ago", etc.)
   */
  static parseRelativeDate(dateStr: string): dayjs.Dayjs | null {
    if (!dateStr || typeof dateStr !== 'string') {
      return null
    }

    const str = dateStr.trim()
    const today = dayjs()

    // Match relative date patterns
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
