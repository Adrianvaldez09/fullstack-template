import { useState, useContext } from 'react'
import SignUp from './components/forms/SignUp'
import Login from './components/forms/Login'
import Dashboard from './components/dashboard/Dashboard'
import { AuthContext } from './AuthProvider'
import './App.css'

function App() {
  const { user } = useContext(AuthContext)
  return (
    <>
      {user ? <Dashboard /> : 
        <div>
          <Login />
          <SignUp />
        </div>
      }
    </>
  )
}

export default App
