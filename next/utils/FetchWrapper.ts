import fetch from 'isomorphic-unfetch'

export default async function FetchWrapper(
  input: RequestInfo,
  init?: RequestInit
) {
  try {
    const data = await fetch(input, init).then((res) => {
      return res.json()
    })
    return data
  } catch (err) {
    throw new Error(err.message)
  }
}
