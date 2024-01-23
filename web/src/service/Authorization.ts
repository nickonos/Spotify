export const Login = (code: string) => {
    fetch(process.env.GATEWAY_URL + "/api/login?code="+ code)
}