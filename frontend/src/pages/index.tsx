import Layout from "@/components/Layout";
import ImageCard from "@/components/NewsCard/ImageCard";
import { News } from "@/shared/types";

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
              title={news.Title}
              guidUrl={news.GuidUrl}
              enclosureUrl={news.EnclosureUrl}
              createdAt={news.CreatedAt}
            />
          );
        })}
      </div>
    </Layout>
  );
}
