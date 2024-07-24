import React from 'react';
import logo from '../assets/logo-smkdev.png';

const Header = () => {
  return (
    <header className="bg-gray-800 text-white p-5 flex justify-between items-center">
      <div className="flex items-center">
        <img src={logo} alt="SMK Dev Logo" className="h-10 mr-3" />
        
      </div>
      <nav>
        <ul className="flex space-x-4">
          <li><a href="#learn" className="hover:underline">Learn</a></li>
          <li><a href="#community" className="hover:underline">Community</a></li>
          <li><a href="#blog" className="hover:underline">Blog</a></li>
          <li><a href="#dashboard" className="hover:underline">Dashboard</a></li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;

