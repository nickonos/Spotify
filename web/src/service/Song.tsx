import { useMutation, useQuery } from "@tanstack/react-query"
import { useSessionProvider } from "./SessionProvider"
import { json } from "stream/consumers"
const GATEWAY_URL = import.meta.env.VITE_GATEWAY_URL

type Song = {
    id: number
    name: string
    artist: string
    cover_url: string
}

export const useSongByName = (name: string) => {
    const {jwt} = useSessionProvider()

    return useQuery<Song>({queryKey: ["song", name], queryFn: () => fetch(GATEWAY_URL + "/api/song?name=" + name, {
        headers: {
            Authorization: jwt
        }
    }).then((res)=> {
        if (!res.ok)
            return null

        return res.json().then(json => json.data) 
    })})
}

export const useCreateSong = () => {
    const {jwt} = useSessionProvider()

    return useMutation({mutationKey: ["song"], mutationFn: (song: Omit<Song, "id">) => fetch(GATEWAY_URL + "/api/song", {
        method: "POST",
        body: JSON.stringify(song),
        headers: {
            Authorization: jwt
        }
    })})
}