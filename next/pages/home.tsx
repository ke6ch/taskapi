import React, { useEffect } from 'react'
import FetchWrapper from '../utils/FetchWrapper'
import fetch from 'isomorphic-unfetch'

interface User {
  username: string
  location: string
}

// export async function getServerSideProps() {
//   // const result = await FetchWrapper('http://localhost:3000/data.json')
//   const data = await FetchWrapper('https://api.github.com/repos/zeit/next.js')
//   console.log(data.stargazers_count)
//   // return { props: { data: { username: 'aaa' } } }
//   // return { props: { data } }
//   return {
//     props: {
//       stars: data.stargazers_count,
//     },
//   }
// }

const email = 'aaa'
const password = 'bbb'
export default function Home() {
  useEffect(() => {
    ;(async () => {
      const res = await fetch(process.env.BASE_URL + '/home', {
        method: 'GET',
        mode: 'cors',
        // credentials: 'include',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
        },
        // body: JSON.stringify({ email, password }),
      })
      console.log(res)
    })()
  }, [])

  return (
    <div>
      <div>Logined</div>
      <div>Logined</div>
    </div>
  )
}
