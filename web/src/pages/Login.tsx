import { Button } from "@/components/ui/button";
import { GetSpotifyOAuthURL } from "@/service/Spotify";
import { useEffect } from "react";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";

export const LoginPage = () => {
  const url = GetSpotifyOAuthURL()
  const [cookies, __, ___] = useCookies(['jwt']);
  const navigate = useNavigate()

  useEffect(() => {
    if (cookies.jwt){
      navigate("/")
    }
  }, [cookies])

  return (
    <div className="w-full h-full flex justify-center items-center">
      <Button onClick={() => {
         window.location.href = url
      }}>Login</Button>
    </div>
  );
};
