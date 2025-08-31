import { useState, useContext } from 'react'
import { AuthContext } from '../../AuthProvider'

function Login() {
  const [identity, setIdentity] = useState('')
  const [password, setPassword] = useState('')

  const { validateUser } = useContext(AuthContext)

  const handleLogin = async (e) => {
    e.preventDefault()


    const newUser = {
      identity,
      password
    }

    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include',
      body: JSON.stringify(newUser)
    })

    if (res.ok) {
      const data = await res.json()
      console.log('User logged in successfully:', data)
      validateUser()
      setIdentity('')
      setPassword('')
    }
  }

  return (
    <>
      <h3>Log in</h3>
      <form onSubmit={handleLogin}>
        <input
          type="text"
          placeholder="Username or Email"
          value={identity}
          onChange={(e) => setIdentity(e.target.value)} 
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)} 
        />
        <button type='submit'>Log in</button>
      </form>
    </>
  )
}

export default Login
