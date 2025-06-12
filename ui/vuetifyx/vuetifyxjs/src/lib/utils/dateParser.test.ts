import { EnhancedDateParser } from './dateParser'
import dayjs from 'dayjs'

/**
 * 测试增强日期解析器的功能
 */
export function testEnhancedDateParser() {
  console.log('=== 增强日期解析器测试 ===')

  const testCases = [
    // 标准格式
    '2023-12-25',
    '2023/12/25',
    '2023.12.25',

    // 中文格式
    '2023年12月25日',
    '2023年12月25日 14:30:00',
    '23年12月25日',

    // 中文数字
    '二零二三年十二月二十五日',
    '二零二三年一月一日',

    // 只有月日
    '12月25日',
    '12-25',
    '12/25',

    // 相对日期
    '今天',
    '明天',
    '昨天',
    '后天',
    '前天',
    '3天后',
    '2周前',
    '1个月后',

    // 错误格式（应该解析失败）
    '无效日期',
    '2023年13月40日'
  ]

  testCases.forEach((testCase) => {
    const parsed = EnhancedDateParser.parseDate(testCase)
    const isValid = EnhancedDateParser.isValidDate(testCase)
    const formatted = parsed
      ? EnhancedDateParser.formatDate(parsed, 'YYYY-MM-DD HH:mm:ss')
      : '解析失败'

    console.log(`输入: "${testCase}" | 有效: ${isValid} | 解析结果: ${formatted}`)
  })

  // 测试相对日期
  console.log('\n=== 相对日期测试 ===')
  const relativeCases = ['3天后', '2周前', '1个月后', '2年前']
  relativeCases.forEach((testCase) => {
    const parsed = EnhancedDateParser.parseRelativeDate(testCase)
    const formatted = parsed ? parsed.format('YYYY-MM-DD') : '解析失败'
    console.log(`相对日期: "${testCase}" | 结果: ${formatted}`)
  })

  // 测试支持的格式
  console.log('\n=== 支持的格式 ===')
  const formats = EnhancedDateParser.getSupportedFormats()
  console.log('支持的日期格式:', formats.slice(0, 10), '...等共', formats.length, '种格式')
}

// 在开发环境下自动运行测试
if (process.env.NODE_ENV === 'development') {
  // testEnhancedDateParser()
}
