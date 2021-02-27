import React from 'react'
import App from 'next/app'
import '../css/tailwind.css'
import { Provider } from 'react-redux'
import withRedux from 'next-redux-wrapper'
import { store } from '../store'
// import store from '../stores'

class MyApp extends App {
  render() {
    const { Component, pageProps, store } = this.props
    return (
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    )
  }
}

export default withRedux(store)(MyApp)
