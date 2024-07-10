import React from 'react';
import { FaRegNewspaper, FaCode, FaPaintBrush } from 'react-icons/fa';

const Articles = () => {
  return (
    <section className="py-20 bg-white">
      <div className="container mx-auto text-center">
        <h2 className="text-4xl font-bold mb-10 text-gray-800">Latest Articles</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-10">
          <div className="p-5 bg-gray-100 shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <FaRegNewspaper className="text-blue-500 text-4xl mb-4" />
            <h3 className="text-2xl font-bold mb-3 text-gray-700">The Future of Web Development</h3>
            <p className="text-gray-600 mb-4">Exploring the latest trends and technologies in web development.</p>
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Read More
            </button>
          </div>
          <div className="p-5 bg-gray-100 shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <FaPaintBrush className="text-blue-500 text-4xl mb-4" />
            <h3 className="text-2xl font-bold mb-3 text-gray-700">Designing for User Experience</h3>
            <p className="text-gray-600 mb-4">How to create user-centric designs that engage and delight.</p>
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Read More
            </button>
          </div>
          <div className="p-5 bg-gray-100 shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <FaCode className="text-blue-500 text-4xl mb-4" />
            <h3 className="text-2xl font-bold mb-3 text-gray-700">Mastering JavaScript Frameworks</h3>
            <p className="text-gray-600 mb-4">A comprehensive guide to the most popular JavaScript frameworks.</p>
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
              Read More
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Articles;
