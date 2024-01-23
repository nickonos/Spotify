import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useSessionProvider } from "@/service/SessionProvider";
import { useCreateSong } from "@/service/Song";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const UploadPage = () => {
  const { role } = useSessionProvider();
  const navigate = useNavigate();
  const {mutate} = useCreateSong()

  useEffect(() => {
    if (role !== "admin") {
      navigate("/");
    }
  }, [role]);

  return (
    <div >
      <form className="flex flex-col gap-4" onSubmit={(e) => {
        e.preventDefault()
        
        const song = {
            name: (e.currentTarget.children.namedItem("name") as HTMLInputElement).value,
            artist: (e.currentTarget.children.namedItem("artist") as HTMLInputElement).value,
            cover_url: (e.currentTarget.children.namedItem("cover_image") as HTMLInputElement).value
        }
        
        mutate(song)
      }}>
        <div className="text-xl">Create a new song</div>
        <Input placeholder="name" id="name" required></Input>
        <Input placeholder="artist" id="artist" required></Input>
        <Input placeholder="cover_image" id="cover_image" required></Input>
        <Button type="submit">Sumbit</Button>
      </form>
    </div>
  );
};
