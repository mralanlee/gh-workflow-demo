import express from 'express'

const app = express()
const PORT = 3000

app.get('/', (req:any, res:any) => {
  res.send(`
  <html>
    <h1>Hello World</h1>
  </html>`)
})

console.log('Running')

app.listen(PORT, () => console.log(`Listening on ${PORT}`))
