import request from 'superagent'

const request2 = u => {
  return new Promise((resolve, reject) => {
    request
      .get(u)
      .end((err, res) => {
        if (err) reject(err)
        if (res) resolve(res)
      })
  })
}

export default { request2 }
