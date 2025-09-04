import { useContext, useEffect, useState } from 'react'
import { AuthContext } from '../../AuthProvider'
import Plays from '../plays/Plays'

function Dashboard() {
    const { validateUser, user, setUser } = useContext(AuthContext)
    const [playName, setPlayName] = useState('')
    const [playDescription, setPlayDescription] = useState('')

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

    useEffect(() => {
        console.log('User in Dashboard:', user)
    }, [user])

    const handleCreatePlay = async (e) => {
        e.preventDefault()

        console.log(user.currentUser.ID)

        const newPlay = {
            name: playName,
            description: playDescription,
            creatorId: user.currentUser.ID
        }

        const res = await fetch('http://localhost:8080/plays', {
            method: 'POST',
            headers: {
            "Content-Type": "application/json"
            },
            body: JSON.stringify(newPlay)
        })

        if (res.ok) {
            const data = await res.json()
            console.log('Play created successfully:', data)
            setPlayName('')
            setPlayDescription('')
            validateUser() // Refresh user data to include new play
        }
    }

  return (
    <>
        <div>Welcome back, {user.currentUser.firstName}</div>
        <form onSubmit={handleCreatePlay}>
            <input 
                type="text" 
                placeholder='Play Name'
                value={playName}
                onChange={(e) => setPlayName(e.target.value)}
            />
            <input 
                type="text" 
                placeholder='Play Description'
                value={playDescription}
                onChange={(e) => setPlayDescription(e.target.value)}
            />
            <button type='submit'>
                Create Play
            </button>
        </form>
        <Plays />
        <button onClick={handleLogout}>
            Log out
        </button>
    </>
  )
}

export default Dashboard