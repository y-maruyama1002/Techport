import { ChangeEvent, useState } from "react";
import axios from "axios";

const NewBlog = () => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const handleSubmit = async () => {
    console.log("title is");
    console.log(title);
    console.log("body is");
    console.log(body);
    console.log(`http://localhost:3000/api/v1/blogs`);
    try {
      await axios.post(`http://localhost:3000/api/v1/blogs`, {
        title,
        body,
      });
    } catch (err) {
      alert("faild to create blog");
    }
  };

  return (
    <section className="max-w-4xl p-20 mx-auto bg-white rounded-md shadow-md dark:bg-gray-800">
      <h2 className="text-lg font-semibold text-gray-700 capitalize dark:text-white">
        Create Blog
      </h2>

      <form>
        <div className="">
          <div>
            <label className="text-gray-700 dark:text-gray-200">Title</label>
            <input
              id="username"
              type="text"
              className="block w-full px-4 py-2 mt-2 mb-5 text-gray-700 bg-white border border-gray-200 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 dark:focus:border-blue-300 focus:outline-none focus:ring"
              onChange={(e: ChangeEvent<HTMLInputElement>) =>
                setTitle(e.target.value)
              }
            />
          </div>

          <div>
            <label className="text-gray-700 dark:text-gray-200">Body</label>
            <textarea
              id="message"
              className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Write your thoughts here..."
              onChange={(e: ChangeEvent<HTMLTextAreaElement>) =>
                setBody(e.target.value)
              }
            ></textarea>
          </div>
        </div>

        <div className="flex justify-end mt-6">
          <div
            className="px-8 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600 cursor-pointer"
            onClick={handleSubmit}
          >
            Save
          </div>
        </div>
      </form>
    </section>
  );
};

export default NewBlog;
