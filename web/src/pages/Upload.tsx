import { useSessionProvider } from "@/service/SessionProvider"
import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

export const UploadPage = () => {
    const {role} = useSessionProvider()
    const navigate = useNavigate()

    useEffect(() => {
        if (role !== "admin"){
            navigate("/")
        }
    }, [role])
    
    return <>upload</>
}