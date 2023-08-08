const express = require('express')
const next = require('next')
const { createProxyMiddleware } = require("http-proxy-middleware")

const port = process.env.PORT || 3000
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()

const proxyTable = {
  '/api': {
    target: 'http://127.0.0.1',
    pathRewrite: {'^/api': '/api'},
    changeOrigin: true
  }
}

app.prepare().then(() => {
  const server = express()
  if (dev) {
    server.use('/api', createProxyMiddleware(proxyTable['/api']));
  }

  // 托管所有请求
  server.all('*', (req, res) => {
    return handle(req, res)
  })

  server.listen(port, (err) => {
    if (err) throw err
    console.log(`> Ready on http://localhost:${port}`)
  })
}).catch(err => console.log('Error:::::', err))
