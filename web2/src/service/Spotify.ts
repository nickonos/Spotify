const SCOPE = "user-read-private user-read-email"

export const GetSpotifyOAuthURL = () => {
    return encodeURI(`https://accounts.spotify.com/authorize?response_type=code&client_id=${process.env.SPOTIFY_CLIENT_ID}&scope=${SCOPE}&redirect_uri=${process.env.BASE_URL + "/auth/callback"}&state=test`)
}
