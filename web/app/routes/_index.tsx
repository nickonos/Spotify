import { Button } from "@/components/ui/button";
import type { MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { GetSpotifyOAuthURL } from "~/api/spotify.server";

export const meta: MetaFunction = () => {
  return [
    { title: "Spotify" },
    { name: "description", content: "Welcome to Spotify!" },
  ];
};

export async function loader() {
  return GetSpotifyOAuthURL()
}

export default function Index() {
  const url = useLoaderData<typeof loader>()

  return (
    <div style={{ fontFamily: "system-ui, sans-serif" }}>
      <Button onClick={() => {
        window.location.href = url
      }}>Hello</Button>
    </div>
  );
}
