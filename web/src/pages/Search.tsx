import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
import { useSongByName } from "@/service/Song";
import { PlayIcon } from "lucide-react";
import { useState } from "react";

export const SearchPage = () => {
  const [name, setName] = useState("");

  return (
    <div className="flex flex-col gap-4">
      <Input
        placeholder="Search"
        onChange={(e) => {
          setName(e.target.value);
        }}
      ></Input>
      <SongSelect name={name} />
    </div>
  );
};

const SongSelect = ({ name }: { name: string }) => {
  const { data, isLoading } = useSongByName(name);

  if (isLoading || !data) {
    return <Skeleton className="w-80 h-48"></Skeleton>;
  }

  return (
    <Card className="p-4 flex flex-col gap-2 w-80 relative hover:bg-[#ffffff20]"  >
      <img src={data.cover_url} className="w-20" />
      <div className="text-2xl font-bold">{data.name}</div>
      <div className="text-sm flex gap-2">
        <div className="text-gray-500">Song</div>
        <div>{data.artist}</div>
      </div>
      <div className="absolute rounded-[100%] bg-[#1ed760] bottom-2 right-2 w-10 h-10 flex justify-center items-center"><PlayIcon fill="black" stroke="black" style={{
        transform: "translateX(2px)"
      }} /></div>
    </Card>
  );
};
