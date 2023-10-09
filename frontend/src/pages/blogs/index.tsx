import { Blog } from "@/features/blogs/types";
import { error } from "console";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/router";

type Props = {
  blogs: Blog[];
};

export async function getStaticProps() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/blogs`);
  const blogs = await res.json();
  return {
    props: {
      blogs,
    },
    revalidate: 60 * 60 * 24, // 24 hours
  };
}

export default function BlogIndex({ blogs }: Props) {
  return (
    <div className="p-20">
      <h1 className="text-3xl font-bold mb-4">Blogs</h1>
      <Link href="blogs/new">
        <button className="px-6 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80">
          create blog
        </button>
      </Link>
      {blogs.map((blog) => (
        <div
          key={blog.id}
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
        >
          <div className="w-full max-w-md px-8 py-4 mt-16 bg-white rounded-lg shadow-lg dark:bg-gray-800">
            <div className="flex justify-center -mt-16 md:justify-end"></div>

            <Link href={`blogs/${blog.id}`}>
              <h2 className="mt-2 text-xl font-semibold text-blue-800 dark:text-white md:mt-0">
                {blog.title}
              </h2>
            </Link>

            <p className="mt-2 text-sm text-gray-600 dark:text-gray-200">
              {blog.body}
            </p>

            <div className="flex justify-end mt-4">
              <a
                href="#"
                className="text-lg font-medium text-blue-600 dark:text-blue-300"
                role="link"
              >
                edit
              </a>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
