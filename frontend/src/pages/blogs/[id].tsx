import { Blog } from "@/features/blogs/types";
import { useRouter } from "next/router";
import Link from "next/link";
import axios from "axios";

type Props = {
  blog: Blog;
};

export const getStaticPaths = async () => {
  const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/blogs`);

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
    fallback: true,
  };
};

export async function getStaticProps({ params }: { params: { id: string } }) {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_BACKEND_URL}/blogs/${params.id}`
  );
  const blog = await res.json();
  return { props: { blog }, revalidate: 60 };
}

const ShowBlog = ({ blog }: Props) => {
  const router = useRouter();

  const handleDelete = async (blogId: number) => {
    try {
      await axios.delete(`http://localhost:3000/api/v1/blogs/${blog.id}`);
      router.reload();
    } catch (err) {
      alert(`failed to delete blog: ${err}`);
    }
  };

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
            <Link
              href={`edit/${blog.id}`}
              className="mx-2 text-blue-600 cursor-pointer dark:text-blue-400 hover:underline"
            >
              edit
            </Link>
            <a
              className="mx-2 text-red-600 cursor-pointer dark:text-blue-400 hover:underline"
              role="link"
              onClick={() => handleDelete(blog.id)}
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
