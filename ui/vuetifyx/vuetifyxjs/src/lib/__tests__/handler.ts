import { http, HttpResponse } from 'msw'

export const autocompleteHandle = [
  http.get('/autocomplete', () => {
    return HttpResponse.json({
      data: [
        {
          id: 1,
          title: '老张头'
        },
        {
          id: 2,
          title: '老李头'
        },
        {
          id: 3,
          title: '老赵头'
        },
        {
          id: 4,
          title: '老钱头'
        },
        {
          id: 5,
          title: '老孙头'
        },
        {
          id: 6,
          title: '老周头'
        }
      ],
      total: 6,
      pages: 0,
      current: 0
    })
  })
]
