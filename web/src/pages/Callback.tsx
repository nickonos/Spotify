import { useLogin } from "@/service/Authorization";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useCookies } from "react-cookie";

export const CallbackPage = () => {
  const [params, _] = useSearchParams();
  const [__, setCookie, ___] = useCookies(['jwt']);
  const code = params.get("code")
  const navigate = useNavigate();

  const { status, data } = useLogin(code ?? "");

  if (status === "pending") return <>loading...</>;

  if (status === "error") {
    navigate("/auth/login");
    return <></>;
  }

  if (status === "success"){
    setCookie("jwt", data, {
      path: "/"
    })
    navigate("/")
    return <>logged in {data}</>
  }
};
