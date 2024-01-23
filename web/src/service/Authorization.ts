import { useQuery } from "@tanstack/react-query"
const GATEWAY_URL = import.meta.env.VITE_GATEWAY_URL

export const useLogin = (code: string) => {
    return useQuery({
        queryKey: ["login"], queryFn: () => fetch(GATEWAY_URL + "/api/login?code=" + code).then((res) => {
            if (!res.ok) {
                throw new Error("Could not send request")
            }
            return res.json().then((res) => res.data)
        }), retry: false
    })
}