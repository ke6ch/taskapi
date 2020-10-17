import fetch from 'isomorphic-unfetch'

export default async function FetchWrapper(
  input: RequestInfo,
  init?: RequestInit
) {
  try {
    console.log(input)
    const data = await fetch(input, init).then((res) => {
      console.log(res)
      return res.json()
    })
    return data
  } catch (err) {
    console.log('Error')
    console.log(err)
    throw new Error(err.message)
  }
}
