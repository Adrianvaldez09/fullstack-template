import { use, useContext, useEffect, useState } from 'react'
import { AuthContext } from '../../AuthProvider'


function Plays() {
    const { user, setUser } = useContext(AuthContext)

    const [plays, setPlays] = useState([])

    useEffect(() => {
        console.log(user.currentUser.plays)
        setPlays(user.currentUser.plays || [])
    }, [user])

  return (
    <>
        {
            user && plays.length > 0 ? (
                <ul>
                    {plays.map((play) => (
                        <li key={play.ID}>
                            <h4>{play.name}</h4>
                            <p>{play.description}</p>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No plays available.</p>
            )
        }
    </>
  )
}

export default Plays