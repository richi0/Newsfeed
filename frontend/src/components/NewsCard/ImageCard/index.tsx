import Image from "next/image";
import { ImageCardProps } from "./types";
import { timeAgo } from "@/shared/time";

const ImageCard = ({
  title,
  enclosureUrl,
  createdAt,
  guidUrl,
}: ImageCardProps) => {
  return (
    <div className="m-3">
      <a href={guidUrl} target="_blank" className="flex gap-3">
        <Image
          src={enclosureUrl}
          alt="title image"
          className="rounded-md"
          width={100}
          height={100}
        />
        <span className="flex flex-col bottom-1 left-2 w-32">
          <span className="text-lg font-bold">{title}</span>
          <span className="text-gray-300 text-sm">{timeAgo(createdAt)}</span>
        </span>
      </a>
    </div>
  );
};

export default ImageCard;
