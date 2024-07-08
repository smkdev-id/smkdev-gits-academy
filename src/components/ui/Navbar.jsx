import { useState } from "react";
import img from "../../assets/image";

const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);

  const handleClick = () => {
    setIsOpen((prevState) => !prevState);
  };
  
  const navbarItems = [
    {
      name: "Learn",
      submenu: [
        "Bootcamp", 
        "Expert Class", 
        "Challenges"
      ],
    },
    { name: "Community", submenu: [] },
    { name: "Blog", submenu: [] },
  ];

  return (
    <div className="fixed z-10 justify-between w-full mb-20 bg-white shadow-sm navbar md:px-16 lg:px-16">
      <div className="navbar-start">
        <div className="dropdown">
          <div
            tabIndex={0}
            role="button"
            className="btn btn-ghost lg:hidden"
            onClick={handleClick}
          >
            <label className="btn btn-circle btn-ghost swap swap-rotate">
              <input type="checkbox" checked={isOpen} onChange={handleClick} />
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
              {navbarItems.map((item, i) => (
                <li key={i}>
                  <a>{item?.name}</a>
                  {item?.submenu.length > 0 && (
                    <ul className="p-2">
                      {item?.submenu?.map((subItem, subIndex) => (
                        <li key={subIndex}>
                          <a>{subItem}</a>
                        </li>
                      ))}
                    </ul>
                  )}
                </li>
              ))}
            </ul>
          )}
        </div>
        <a className="text-xl cursor-pointer" href="/">
          <img src={img.logo} height={120} width={120} alt="logo img" />
        </a>
      </div>
      <div className="hidden navbar-end lg:flex">
        <ul className="px-1 text-lg font-semibold menu menu-horizontal">
          {navbarItems.map((item, i) => (
            <li key={i}>
              {item?.submenu?.length > 0 ? (
                <details>
                  <summary className="hover:bg-white">{item?.name}</summary>
                  <ul className="p-2 text-sm">
                    {item?.submenu?.map((subItem, subIndex) => (
                      <li
                        key={subIndex}
                        className="rounded hover:bg-primary hover:text-white"
                      >
                        <a>{subItem}</a>
                      </li>
                    ))}
                  </ul>
                </details>
              ) : (
                <a className="hover:bg-white">{item?.name}</a>
              )}
            </li>
          ))}
        </ul>
      </div>
      <a className="text-white transition-all px-7 btn btn-primary">
        Dashboard
      </a>
    </div>
  );
};

export default Navbar;
