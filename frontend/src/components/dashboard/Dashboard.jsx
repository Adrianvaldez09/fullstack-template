import { useContext } from 'react'
import { AuthContext } from '../../AuthProvider'

function Dashboard() {
    const { validateUser } = useContext(AuthContext)

    const handleLogout = async () => {
        // Log out logic here
        const res = await fetch('http://localhost:8080/logout', {
            method: 'POST',
            credentials: 'include'
        })

        if (res.ok) {
            validateUser()
            console.log('User logged out successfully')
        }
    }

  return (
    <>
        <div>Dashboard</div>
        <button onClick={handleLogout}>
            Log out
        </button>
    </>
  )
}

export default Dashboard