import React from 'react';
import { Link } from 'react-router-dom';
import StyleNavbar from './Navbar.module.css';
import Logo from '../assets/logosmkdev.png';
export default function Navbar() {
  return (
    <>
      <nav className="navbar sticky-top bg-light border-bottom border-warning border-2 shadow">
        <div className="container">
          <Link to="/" className="navbar-logo" href="#">
            <img src={Logo} width="140" height="30" alt="" />
          </Link>

          <ul className="nav">
            <li className="nav-item">
              <Link to="/" className="nav-link link-dark fw-bold">
                Home
              </Link>
            </li>
            <li className="nav-item">
              <Link to="/list-menu" className="nav-link link-dark fw-bold">
                List Menu
              </Link>
            </li>
            <li className="nav-item">
              <a href="#" className="nav-link link-dark fw-bold">
                Contact
              </a>
            </li>
            <button className={`${StyleNavbar.dashboard}`}>Dashboard</button>
          </ul>
        </div>
      </nav>
    </>
  );
}
