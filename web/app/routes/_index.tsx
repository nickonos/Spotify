import { Button } from "@/components/ui/button";
import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "Spotify" },
    { name: "description", content: "Welcome to Spotify!" },
  ];
};

export default function Index() {
  return (
    <div style={{ fontFamily: "system-ui, sans-serif" }}>
      <Button className="">Hello</Button>
    </div>
  );
}
