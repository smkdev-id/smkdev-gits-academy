import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBook, faUsers, faProjectDiagram } from '@fortawesome/free-solid-svg-icons';

const Features = () => {
  return (
    <section className="py-20 bg-gray-100">
      <div className="container mx-auto text-center">
        <h2 className="text-4xl font-bold mb-10 text-gray-800">Our Features</h2>
        <div className="flex flex-wrap justify-center gap-8">
          <div className="max-w-sm p-5 bg-white shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <h3 className="text-2xl font-bold mb-3 text-gray-700">High Quality Content</h3>
            <FontAwesomeIcon icon={faBook} className="text-blue-500 text-3xl mb-3" />
            <p className="text-gray-600">Access a wide range of tutorials and courses designed by experts.</p>
          </div>
          <div className="max-w-sm p-5 bg-white shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <h3 className="text-2xl font-bold mb-3 text-gray-700">Interactive Community</h3>
            <FontAwesomeIcon icon={faUsers} className="text-green-500 text-3xl mb-3" />
            <p className="text-gray-600">Engage with a community of learners and professionals.</p>
          </div>
          <div className="max-w-sm p-5 bg-white shadow-lg rounded-lg transform transition duration-500 hover:scale-105">
            <h3 className="text-2xl font-bold mb-3 text-gray-700">Real World Projects</h3>
            <FontAwesomeIcon icon={faProjectDiagram} className="text-red-500 text-3xl mb-3" />
            <p className="text-gray-600">Work on projects that simulate real-world scenarios.</p>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Features;
