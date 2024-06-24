import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default createVuetify({
  components,
  directives,
  icons: {
    iconfont: 'mdi',
  },
  theme: {
    themes: {
      light: {
        primary: '#5cb85c',
        secondary: '#4cae4c',
        accent: '#d9534f',
        error: '#c9302c',
      },
    },
  },
})
