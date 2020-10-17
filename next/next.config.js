require('dotenv').config()

module.exports = {
  env: {
    BASE_URL:
      process.env.BASE_URL === undefined
        ? 'http://localhost:1323'
        : process.env.BASE_URL,
  },
}
