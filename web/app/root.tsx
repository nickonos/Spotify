import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";

import styles from "./globals.css";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { HomeIcon, SearchIcon } from "@/components/icons";

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: styles },
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
];

export default function App() {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="dark h-screen">
        <Layout />
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  );
}

const Layout = () => (
  <ResizablePanelGroup
    direction={"horizontal"}
    className="flex h-screen w-screen gap-1 p-2"
  >
    <ResizablePanel
      defaultSize={20}
      maxSize={40}
      className="h-full bg-black p-1"
    >
      <ResizablePanelGroup direction="vertical" className="flex gap-1">
        <ResizablePanel defaultSize={10}>
          <div className="flex">
            <HomeIcon />
            <div>Home</div>
          </div>
          <div className="flex">
            <SearchIcon />
            <div>Search</div>
          </div>
        </ResizablePanel>
        <ResizableHandle hidden />
        <ResizablePanel></ResizablePanel>
      </ResizablePanelGroup>
    </ResizablePanel>
    <ResizableHandle hidden />
    <ResizablePanel>
      <Outlet />
    </ResizablePanel>
  </ResizablePanelGroup>
);
