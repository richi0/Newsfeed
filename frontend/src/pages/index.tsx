export async function getServerSideProps() {
  const res = await fetch("http://localhost:4000/");
  const content = await res.text();
  return {
    props: { content }, // will be passed to the page component as props
  };
}

export default function Home({ content }: any) {
  return <div>{content}</div>;
}
