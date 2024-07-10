import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faQuoteLeft, faQuoteRight } from '@fortawesome/free-solid-svg-icons';

const Testimonials = () => {
  return (
    <section className="bg-gray-100 py-20">
      <div className="container mx-auto text-center">
        <h2 className="text-4xl font-bold mb-10 text-gray-800">What Our Users Say</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-10 max-w-4xl mx-auto">
          <div className="p-5 bg-white shadow-lg rounded-lg">
            <FontAwesomeIcon icon={faQuoteLeft} className="text-gray-400 text-3xl mb-3" />
            <blockquote className="italic mb-4">"SMK Dev has transformed my career. The courses are top-notch!"</blockquote>
            <FontAwesomeIcon icon={faQuoteRight} className="text-gray-400 text-3xl mb-3 float-right" />
            <div className="flex items-center mt-5">
              <div>
                <cite className="block font-bold">Jane Doe</cite>
                <span className="text-gray-600">Software Engineer</span>
              </div>
            </div>
          </div>
          <div className="p-5 bg-white shadow-lg rounded-lg">
            <FontAwesomeIcon icon={faQuoteLeft} className="text-gray-400 text-3xl mb-3" />
            <blockquote className="italic mb-4">"The community support is amazing. I learned so much from peers."</blockquote>
            <FontAwesomeIcon icon={faQuoteRight} className="text-gray-400 text-3xl mb-3 float-right" />
            <div className="flex items-center mt-5">
              <div>
                <cite className="block font-bold">John Smith</cite>
                <span className="text-gray-600">Web Developer</span>
              </div>
            </div>
          </div>
          <div className="p-5 bg-white shadow-lg rounded-lg">
            <FontAwesomeIcon icon={faQuoteLeft} className="text-gray-400 text-3xl mb-3" />
            <blockquote className="italic mb-4">"The hands-on projects provided real-world experience."</blockquote>
            <FontAwesomeIcon icon={faQuoteRight} className="text-gray-400 text-3xl mb-3 float-right" />
            <div className="flex items-center mt-5">
              <div>
                <cite className="block font-bold">Alice Johnson</cite>
                <span className="text-gray-600">Data Scientist</span>
              </div>
            </div>
          </div>
          <div className="p-5 bg-white shadow-lg rounded-lg">
            <FontAwesomeIcon icon={faQuoteLeft} className="text-gray-400 text-3xl mb-3" />
            <blockquote className="italic mb-4">"I improved my skills significantly, thanks to the expert instructors."</blockquote>
            <FontAwesomeIcon icon={faQuoteRight} className="text-gray-400 text-3xl mb-3 float-right" />
            <div className="flex items-center mt-5">
              <div>
                <cite className="block font-bold">Michael Brown</cite>
                <span className="text-gray-600">UI/UX Designer</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Testimonials;
