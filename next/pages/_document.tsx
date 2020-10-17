import React from 'react'
import Document, { Head, Main, NextScript } from 'next/document'

export default class MyDocument extends Document {
  render() {
    return (
      <html>
        <Head>
          <script> </script>
        </Head>
        <body>
          <Main />
          <NextScript />
          {/* https://jakearchibald.com/2016/link-in-body/ */}
          <script> </script>
        </body>
      </html>
    )
  }
}
