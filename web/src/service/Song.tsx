import { useQuery } from "@tanstack/react-query"
const GATEWAY_URL = import.meta.env.VITE_GATEWAY_URL

type Song = {
    id: number
    name: string
    artist: string
    cover_url: string
}

export const useSongByName = (name: string) => {
    return useQuery<Song>({queryKey: ["song", name], queryFn: () => fetch(GATEWAY_URL + "/api/song?name=" + name).then((res)=> {
        if (!res.ok)
            return null

        return res.json().then(json => json.data) 
    })})
}