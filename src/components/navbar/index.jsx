import { useState } from "react";
import img from '../../assets/image'

const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);

  function handleClick() {
    setIsOpen((prevState) => !prevState);
  }

  return (
    <div className="justify-between w-full shadow-sm navbar bg-base-100 md:px-16 lg:px-16">
      <div className="navbar-start">
        <div className="dropdown">
          <div 
            tabIndex={0} 
            role="button" 
            className="btn btn-ghost lg:hidden"
            onClick={handleClick}
          >
            <label className="btn btn-circle btn-ghost swap swap-rotate">
              <input 
                type="checkbox" 
                checked={isOpen} 
                onChange={handleClick} 
              />
              <svg
                className={`fill-current ${isOpen ? "swap-off" : ""}`}
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 512 512"
              >
                <path d="M64,384H448V341.33H64Zm0-106.67H448V234.67H64ZM64,128v42.67H448V128Z" />
              </svg>
              <svg
                className={`fill-current ${isOpen ? "" : "swap-on"}`}
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 512 512"
              >
                <polygon points="400 145.49 366.51 112 256 222.51 145.49 112 112 145.49 222.51 256 112 366.51 145.49 400 256 289.49 366.51 400 400 366.51 289.49 256 400 145.49" />
              </svg>
            </label>
          </div>
          {isOpen && (
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
            >
              <li>
                <a>Item 1</a>
              </li>
              <li>
                <a>Parent</a>
                <ul className="p-2">
                  <li>
                    <a>Submenu 1</a>
                  </li>
                  <li>
                    <a>Submenu 2</a>
                  </li>
                </ul>
              </li>
              <li>
                <a>Item 3</a>
              </li>
            </ul>
          )}
        </div>
        <a className="text-xl btn btn-ghost">
          <img src={ img.logo } height={120} width={120} />
        </a>
      </div>
      <div className="hidden navbar-end lg:flex">
        <ul className="px-1 text-lg font-semibold menu menu-horizontal">
        <li>
            <details>
              <summary className="hover:bg-white">Learn</summary>
              <ul className="p-2 text-sm">
                <li className="rounded hover:bg-primary hover:text-white">
                  <a>Bootcamp</a>
                </li>
                <li className="rounded hover:bg-primary hover:text-white">
                  <a>Expert Class</a>
                </li>
                <li className="rounded hover:bg-primary hover:text-white">
                  <a>Challenges</a>
                </li>
              </ul>
            </details>
          </li>
          <li>
            <a className="hover:bg-white">Community</a>
          </li>
          <li>
            <a className="hover:bg-white">Blog</a>
          </li>
        </ul>
      </div>
      <a className="text-white transition-all px-7 btn btn-primary">Dashboard</a>
    </div>
  );
};

export default Navbar;
