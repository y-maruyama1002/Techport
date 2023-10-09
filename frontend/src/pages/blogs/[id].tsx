import { Blog } from "@/features/blogs/types";
import { useRouter } from "next/router";

type Props = {
  blog: Blog;
};

export const getStaticPaths = async () => {
  const res = await fetch(`${process.env.BACKEND_URL}/blogs`);

  const blogs: Blog[] = await res.json();
  const paths = blogs.map((blog) => {
    return {
      params: {
        id: blog.id.toString(),
      },
    };
  });
  return {
    paths,
    fallback: true, // false or "blocking"
  };
};

export async function getStaticProps({ params }: { params: { id: string } }) {
  const res = await fetch(`${process.env.BACKEND_URL}/blogs/${params.id}`);
  const blog = await res.json();
  return { props: { blog }, revalidate: 60 };
}

const ShowBlog = ({ blog }: Props) => {
  const router = useRouter();
  if (router.isFallback) {
    return <div>Loading...</div>;
  }
  return (
    <div className="p-20">
      <div className="w-full max-w-sm px-4 py-3 bg-white rounded-md shadow-md dark:bg-gray-800">
        <div className="flex items-center justify-between">
          <span className="text-sm font-light text-gray-800 dark:text-gray-400">
            {blog.created_at}
          </span>
        </div>

        <div>
          <h1 className="mt-2 text-lg font-semibold text-gray-800 dark:text-white">
            {blog.title}
          </h1>
          <p className="mt-2 text-sm text-gray-600 dark:text-gray-300">
            {blog.body}
          </p>
        </div>

        <div>
          <div className="flex items-center mt-2 text-gray-700 dark:text-gray-200">
            <a
              className="mx-2 text-blue-600 cursor-pointer dark:text-blue-400 hover:underline"
              role="link"
            >
              edit
            </a>
            <a
              className="mx-2 text-red-600 cursor-pointer dark:text-blue-400 hover:underline"
              role="link"
            >
              delete
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ShowBlog;
