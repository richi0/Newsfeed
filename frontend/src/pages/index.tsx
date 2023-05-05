import ImageCard from "@/components/NewsCard/ImageCard";
import Layout from "@/components/Layout";
import { News } from "@/shared/types";
import { getHostName } from "../shared/helpers";

type HomeProps = {
  content: News[];
};

export async function getServerSideProps() {
  const res = await fetch("http://localhost:4000/news");
  const content = (await res.json()) as News[];
  return {
    props: { content },
  };
}

export default function Home({ content }: HomeProps) {
  return (
    <Layout>
      <div className="flex flex-wrap justify-center">
        {content.map((news) => {
          return (
            <ImageCard
              key={news.ID}
              title={news.Title}
              link={news.Link}
              enclosureUrl={news.EnclosureUrl}
              pubDate={news.PubDate}
              source={getHostName(news.Link)}
            />
          );
        })}
      </div>
    </Layout>
  );
}
