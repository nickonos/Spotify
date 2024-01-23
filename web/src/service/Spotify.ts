const SCOPE = "user-read-private user-read-email"
const CLIENT_ID = import.meta.env.VITE_SPOTIFY_CLIENT_ID
const BASE_URL = import.meta.env.VITE_BASE_URL

export const GetSpotifyOAuthURL = () => {
    return encodeURI(`https://accounts.spotify.com/authorize?response_type=code&client_id=${CLIENT_ID}&scope=${SCOPE}&redirect_uri=${BASE_URL + "/auth/callback"}&state=test`)
}
