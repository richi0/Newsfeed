import Image from "next/image";
import { ImageCardProps } from "./types";
import { timeAgo } from "@/shared/time";

const ImageCard = ({
  title,
  enclosureUrl,
  pubDate,
  link,
  source,
}: ImageCardProps) => {
  return (
    <div className="m-3">
      <a href={link} target="_blank" className="flex flex-col gap-3">
        {(enclosureUrl && (
          <Image
            src={enclosureUrl}
            alt="title image"
            className="rounded-md w-32 h-32 bg-gray-200"
            width={128}
            height={128}
          />
        )) || <div className="rounded-md w-32 h-32 bg-gray-200"></div>}
        <span className="flex flex-col bottom-1 left-2 w-32">
          <span className="text-md font-bold">{title}</span>
          <span className="text-sm">{`By ${source}`}</span>
          {(pubDate && (
            <span className="text-gray-300 text-sm">{timeAgo(pubDate)}</span>
          )) ||
            null}
        </span>
      </a>
    </div>
  );
};

export default ImageCard;
