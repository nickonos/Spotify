import { jwtDecode } from "jwt-decode"
import React, { useContext, useEffect } from "react"
import { useCookies } from "react-cookie"
import { Outlet, useLocation, useNavigate } from "react-router-dom"

type AuthContextType = {
    Session?: Session
    LogOut: () => void
}

type Session = {
    jwt: string
    email: string
    role: "admin" | "user"
}

const AuthContext = React.createContext<AuthContextType | null>(null)

const public_routes = ["/auth/login", "/auth/callback"]

export const AuthenticationProvider = () => {
    const [cookies, _, removeCookies] = useCookies(["jwt"])
    const location = useLocation()
    const navigate = useNavigate()

    const isPublicRoute = public_routes.some(publicRoute => location.pathname.startsWith(publicRoute))

    useEffect(() => {
        if (!cookies.jwt && !isPublicRoute){
            navigate("/auth/login")
        }
    }, [cookies, isPublicRoute])
    
    return <AuthContext.Provider value={{
        Session: cookies.jwt ? {
            jwt: cookies.jwt,
            role: (jwtDecode(cookies.jwt) as any).role,
            email: (jwtDecode(cookies.jwt) as any).email,
        } : undefined,
        LogOut: () => {
            removeCookies("jwt")
        }
    }}><Outlet/></AuthContext.Provider>
}

export const useAuthenticationProvider = () => {
    const context = useContext(AuthContext)
    if (!context){
        throw new Error("you cannot use useAuthenticationProvider outside a authentication provider")
    }
    return context
}
