import { Outlet } from "react-router";
import {
  ResizablePanelGroup,
  ResizablePanel,
  ResizableHandle,
} from "./resizable";
import { HomeIcon, SearchIcon } from "../icons";
import { useCookies } from "react-cookie";
import { useEffect, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { LogOutIcon, UploadIcon } from "lucide-react";
import { jwtDecode } from "jwt-decode";
import { SessionProvider } from "@/service/SessionProvider";

export const Layout = () => {
  const [cookies, __, removeCookies] = useCookies(["jwt"]);
  const navigate = useNavigate();

  useEffect(() => {
    if (!cookies.jwt) {
      navigate("/auth/login");
    }
  }, [cookies]);

  const role = useMemo(() => {
    if (!cookies.jwt) return undefined;

    return (jwtDecode(cookies.jwt) as any).role;
  }, [cookies]);

  return (
    <SessionProvider>
      <ResizablePanelGroup
        direction={"horizontal"}
        className="flex h-screen w-screen gap-1 p-2"
      >
        <ResizablePanel
          defaultSize={20}
          maxSize={40}
          className="h-full bg-[#0C0A09] pt-0 p-1"
        >
          <ResizablePanelGroup direction="vertical" className="flex gap-1">
            <ResizablePanel defaultSize={role === "admin" ? 15 : 10}>
              <div
                className="flex"
                onClick={() => {
                  navigate("/");
                }}
              >
                <HomeIcon />
                <div>Home</div>
              </div>
              <div
                className="flex"
                onClick={() => {
                  navigate("/search");
                }}
              >
                <SearchIcon />
                <div>Search</div>
              </div>
              {role === "admin" && (
                <div className="flex gap-4" onClick={() => {
                  navigate("/upload")
                }}>
                  <UploadIcon />
                  <div>Upload</div>
                </div>
              )}
            </ResizablePanel>
            <ResizableHandle hidden />
            <ResizablePanel className="relative">
              <div
                className="flex gap-4 absolute bottom-4"
                onClick={() => {
                  removeCookies("jwt");
                }}
              >
                <LogOutIcon /> <div>Log out</div>
              </div>
            </ResizablePanel>
          </ResizablePanelGroup>
        </ResizablePanel>
        <ResizableHandle hidden />
        <ResizablePanel>
          <Outlet />
        </ResizablePanel>
      </ResizablePanelGroup>
    </SessionProvider>
  );
};
