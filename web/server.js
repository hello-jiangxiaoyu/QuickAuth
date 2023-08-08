const express = require('express')
const next = require('next')
const createProxyMiddleware = require('http-proxy-middleware').createProxyMiddleware;

const hostname = 'localhost'
const port = 3000
const dev = process.env.NODE_ENV !== 'production' //只在开发环境使用
const app = next({dev, hostname, port})
const handle = app.getRequestHandler()

app.prepare().then(() => {
  const server = express()
  const devProxy = {
    ['/api/quick']: {
      target: 'http://localhost',
      changeOrigin: true,
      pathRewrite: {},
    },
  }
  if (dev && devProxy) {
    Object.keys(devProxy).forEach((context) => {
      server.use(createProxyMiddleware(context, devProxy[context]))
    })
  }

  server.all('*', (req, res) => handle(req, res))
  server.listen(port, err => {
    if (err) {
      throw err
    }
  })
}).catch(e => console.log("error: ", e))
