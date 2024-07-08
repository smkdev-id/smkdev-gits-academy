import React from "react";

const BlogCard = ({ judul, deskripsi, gambar }) => {
  return (
    <div className="w-full bg-gray-200 dark:bg-gray-800 pb-10 md:w-[30%] h-full rounded-2xl hover:shadow-2xl my-10">
      <div className="h-[70%]">
        <img src={gambar} alt="manblog" className="w-full h-full rounded-t-2xl" />
      </div>
      <div className="p-4 text-black dark:text-white">
        <h1 className="font-bold text-xl line-clamp-1">{judul}</h1>
        <p className="line-clamp-3 text-sm">{deskripsi}</p>
      </div>
      <div className="flex justify-center mt-5">
        <button className="py-3 px-10 border-2 border-solid border-blue-600 dark:border-blue-300 rounded-xl hover:bg-blue-600 dark:hover:bg-blue-300 hover:text-white dark:hover:text-black transition-all duration-300">
          <h1 className="font-bold text-xl">Read More</h1>
        </button>
      </div>
    </div>
  );
};

export default BlogCard;
