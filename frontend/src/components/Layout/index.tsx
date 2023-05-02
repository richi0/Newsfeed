import Footer from "../Footer";
import Header from "../Header";
import { LayoutProps } from "./types";

const Layout = ({ children }: LayoutProps) => {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="mb-auto">{children}</div>
      <Footer />
    </div>
  );
};

export default Layout;
