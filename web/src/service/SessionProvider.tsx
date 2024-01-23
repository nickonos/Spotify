import React, { ReactNode, useContext, useEffect } from "react"
import { Outlet, useNavigate } from "react-router-dom"
import { useAuthenticationProvider } from "./AuthenticationProvider"

type Session = {
    jwt: string
    email: string
    role: "admin" | "user"
}

const AuthContext = React.createContext<Session | null>(null)


export const SessionProvider = ({children}: {children: ReactNode}) => {
    const navigate = useNavigate()
    const {Session} = useAuthenticationProvider()

    useEffect(() => {
        if (!Session){
            navigate("/auth/login")
        }
    }, [Session])

    if (!Session)
        return <></>
    
    return <AuthContext.Provider value={Session}>{children}</AuthContext.Provider>
}

export const useSessionProvider = () => {
    const context = useContext(AuthContext)
    if (!context){
        throw new Error("you cannot use useSessionProvider outside a session provider")
    }
    return context
}
