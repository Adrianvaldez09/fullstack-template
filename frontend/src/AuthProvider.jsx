import { createContext, useState, useEffect, use } from "react";

export const AuthContext = createContext();

function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [isAuthenticaed, setIsAuthenticated] = useState(false);

    useEffect(() => {
        validateUser();
    }, [])

    const validateUser = async () => {
        const res = await fetch('http://localhost:8080/validate', {
            method: 'GET',
            credentials: 'include'
        })

        if (res.ok) {
            const data = await res.json()
            console.log('Validated user:', data)
            setUser(data)
        } else {
            setUser(null)
            console.log('No valid session')
        }
    }

    const context ={
        user,
        setUser,
        validateUser,
    }


    return (
        <AuthContext.Provider value={context}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
  return useContext(AuthContext)
}

export default AuthProvider;