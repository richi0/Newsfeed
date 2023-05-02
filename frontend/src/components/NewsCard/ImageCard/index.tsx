import { minutesAgo } from "@/shared/time";
import { ImageCardProps } from "./types";

const ImageCard = ({
  title,
  enclosureUrl,
  createdAt,
  guidUrl,
}: ImageCardProps) => {
  return (
    <div className="w-44 h-60 relative m-3">
      <a href={guidUrl} target="_blank">
        <img
          src={enclosureUrl}
          alt="title image"
          className="w-full h-full object-cover rounded-md"
        />
        <span className="flex flex-col absolute bottom-1 left-2">
          <span className="text-white text-lg font-bold">{title}</span>
          <span className="text-gray-300 text-sm">
            {minutesAgo(createdAt) + " mins ago"}
          </span>
        </span>
      </a>
    </div>
  );
};

export default ImageCard;
