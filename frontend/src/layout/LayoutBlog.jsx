import React, { useState } from "react";
import BlogCard from "../components/Blogs/BlogCard";
import Spacer from "../components/Spacer";
import { Link } from "react-router-dom";
import blogData from "../data/blogData";

const LayoutBlog = () => {
  const [visibleBlogs, setVisibleBlogs] = useState(3); // State untuk mengelola jumlah blog post yang ditampilkan

  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  return (
    <>
      <section className={`${padding} font-poppins text-black dark:text-white ${bgLight} ${bgDark}`}>
        <h1 className="font-bold text-center text-4xl text-black dark:text-white">Blog</h1>

        <Spacer height={"h-10"} />

        <div className="md:flex md:flex-wrap md:justify-around">
          {blogData.slice(0, visibleBlogs).map((blog, index) => (
            <BlogCard
              key={index}
              judul={blog.judul}
              deskripsi={blog.deskripsi}
              gambar={blog.gambar}
            />
          ))}
        </div>

        {visibleBlogs < blogData.length && (
          <div className="text-center">
            <Link to="/blogs" className="font-bold text-black dark:text-white hover:text-gray-600 dark:hover:text-gray-400">
              <button className="py-2 px-6 border-2 border-solid border-gray-500 dark:border-gray-400 rounded-xl">
                Show More
              </button>
            </Link>
          </div>
        )}
      </section>
    </>
  );
};

export default LayoutBlog;
