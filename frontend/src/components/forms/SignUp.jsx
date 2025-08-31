import { useState } from 'react'

function SignUp() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [username, setUsername] = useState('')
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')

  const handleSignUp = async (e) => {
    e.preventDefault()

    const newUser = {
        username,
        firstName,
        lastName,
        email,
        password
    }
    console.log(newUser)

    const res = await fetch('http://localhost:8080/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newUser)
    })

    if (res.ok) {
      const data = await res.json()
      console.log('User created successfully:', data)
      setFirstName('')
      setLastName('')
      setUsername('')
      setEmail('')
      setPassword('')
    }
  }

  return (
    <>
      <h3>Sign Up</h3>
      <form onSubmit={handleSignUp}>
        <input
          type="text"
          placeholder="First Name"
          value={firstName}
          onChange={(e) => setFirstName(e.target.value)} 
        />
        <input
          type="text"
          placeholder="Last Name"
          value={lastName}
          onChange={(e) => setLastName(e.target.value)} 
        />
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)} 
        />
        <input
          type="text"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)} 
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)} 
        />
        <button type='submit'>Create account</button>
      </form>
    </>
  )
}

export default SignUp
