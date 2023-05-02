import { AiOutlineFilter, AiOutlineMenu } from "react-icons/ai";

const Header = () => {
  return (
    <header className="p-4 pb-1 shadow-md">
      <div className="flex justify-between mb-5 items-center">
        <AiOutlineMenu />
        <span className="font-bold">NewsFeed</span>
        <AiOutlineFilter />
      </div>
      <nav className="flex gap-2 justify-center sm:gap-8 ">
        <span>Top</span>
        <span>World</span>
        <span>Sports</span>
        <span>Finance</span>
        <span>Politics</span>
      </nav>
    </header>
  );
};

export default Header;
