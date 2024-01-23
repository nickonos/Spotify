import { ActionFunctionArgs, MetaFunction, redirect } from "@remix-run/node";
import {
  Form,
  useFetcher,
  useNavigate,
  useSearchParams,
  useSubmit,
} from "@remix-run/react";
import { useEffect, useRef } from "react";
import { Login } from "~/api/login.server";

export const meta: MetaFunction = () => {
  return [
    { title: "Spotify" },
    { name: "description", content: "Spotify login callback" },
  ];
};

export async function action({request}: ActionFunctionArgs) {
    const formdata = await request.formData()
    const code = formdata.get("code")
    if (!code)
      return null
    
    Login(code.toString())
    return redirect("/")
}

export default function Index() {
  const navigate = useNavigate();
  const [searchParams, _] = useSearchParams();
  const code = searchParams.get("code");
  const submit = useSubmit()
  const ref = useRef(null)

  useEffect(() => {
    if (!ref)
      return
    
    if (!code) {
      navigate("/");
      return
    }

    submit(ref.current)

  }, [code]);

  return <div>
    <Form method="POST" ref={ref} className="hidden">
      <input id="code" value={code ?? undefined} ></input>
    </Form>
  </div>;
}
