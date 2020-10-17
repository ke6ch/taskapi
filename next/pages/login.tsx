import React, { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import Layout from '../components/Layout'
// import FetchWrapper from '../utils/FetchWrapper'
import FetchWrapper from 'isomorphic-unfetch'

export default function Login() {
  const [email, setEmail] = useState<string>('')
  const [password, setPassWord] = useState<string>('')
  const router = useRouter()

  // useEffect(() => {
  //   const f = async () => {
  //     await FetchWrapper(process.env.BASE_URL + '/login', {
  //       method: 'POST',
  //       mode: 'cors',
  //       credentials: 'include',
  //       headers: {
  //         'Content-Type': 'application/json; charset=utf-8',
  //       },
  //     }).then((res) => {
  //       if (res.status === 200) {
  //         router.push('/home', { pathname: '/home' })
  //       } else {
  //         console.log('No COOKIEEEEEEEE')
  //       }
  //     })
  //   }
  //   f()
  // }, [])

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    await FetchWrapper(process.env.BASE_URL + '/session', {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
      body: JSON.stringify({ email, password }),
    }).then((res) => {
      if (res.status === 200) {
        router.push('/', { pathname: '/' })
      }
    })
  }

  return (
    <Layout>
      <div className="flex items-center justify-center mt-4">
        <form
          className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
          onSubmit={(e: React.FormEvent<HTMLFormElement>) => handleSubmit(e)}
        >
          <img
            className="object-contain h-12 mb-4 mx-auto"
            src="brand.webp"
            alt="logo"
          />
          <h1 className="text-center mb-3">Please sign in</h1>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Email Address
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="inputEmail"
              type="text"
              value={email}
              placeholder="Email Address"
              onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                setEmail(e.target.value)
              }
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Password
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
              id="password"
              type="password"
              value={password}
              placeholder="Password"
              onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                setPassWord(e.target.value)
              }
            />
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
            >
              Sign In
            </button>
            <a
              className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
              href="#"
            >
              Forgot Password?
            </a>
          </div>
        </form>
      </div>
    </Layout>
  )
}
