import React, { useState, useEffect } from "react";
import NavLink from "./NavLink";
import Dropdown from "./DropDown";
import ResponsiveNavLink from "./ResponsiveNavLink";
import ToggleMode from "./ToggleMode";
import Spacer from "./Spacer";

const NavigationBar = ({ toggleTheme }) => {
  const [showingNavigationDropdown, setShowingNavigationDropdown] =
    useState(false);
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 50) {
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <>
      <div className="font-poppins bg-white dark:bg-gray-900">
        <nav
          className={`border-b fixed top-0 w-full z-50 ${
            isScrolled ? "blur-bg" : ""
          } bg-gray-200 border-gray-200 dark:bg-black dark:border-white`}
        >
          <div className="max-w-7xl mx-auto px-[30px] md:px-[135px]">
            <div className="flex justify-between h-16">
              <div className="flex items-center">
                <div className="shrink-0 flex items-center w-24 h-24">
                  <img src="/src/assets/logosmk.png" alt="logosmk" />
                </div>
              </div>

              <div className="flex">
                <div className="hidden space-x-8 sm:-my-px sm:ms-10 sm:flex">
                  <NavLink>
                    <h1 className="pb-1 hover:border-b-2 text-black hover:border-black hover:opacity-50 dark:text-white dark:hover:border-white">
                      Learn
                    </h1>
                  </NavLink>
                </div>

                <div className="hidden space-x-8 sm:-my-px sm:ms-10 sm:flex">
                  <NavLink>
                    <h1 className="pb-1 hover:border-b-2 text-black hover:border-black hover:opacity-50 dark:text-white dark:hover:border-white">
                      Community
                    </h1>
                  </NavLink>
                </div>

                <div className="hidden space-x-8 sm:-my-px sm:ms-10 sm:flex">
                  <NavLink>
                    <h1 className="pb-1 hover:border-b-2 text-black hover:border-black hover:opacity-50 dark:text-white dark:hover:border-white">
                      Blog
                    </h1>
                  </NavLink>
                </div>
              </div>

              <div className="hidden sm:flex sm:items-center sm:ms-6">
                <div className="ms-3 relative">
                  <Dropdown>
                    <Dropdown.Trigger>
                      <span className="inline-flex rounded-md">
                        <button
                          type="button"
                          className="inline-flex items-center px-2 py-1 border border-transparent text-sm leading-4 font-medium rounded-md text-black hover:text-gray-500 focus:outline-none transition ease-in-out duration-150 dark:text-white dark:hover:text-gray-700"
                        >
                          <img
                            src="/src/assets/profile.png"
                            alt="Profile Icon"
                            className="w-10 h-10 rounded-full"
                          />
                        </button>
                        <ToggleMode toggleTheme={toggleTheme} />
                      </span>
                    </Dropdown.Trigger>

                    <Dropdown.Content>
                      <Dropdown.Link where={"/profile"}>
                        <h1 className="text-black dark:text-white">Profile</h1>
                      </Dropdown.Link>

                      <Dropdown.Link>
                        <h1 className="text-black dark:text-white">Log Out</h1>
                      </Dropdown.Link>
                    </Dropdown.Content>
                  </Dropdown>
                </div>
              </div>

              <div className="-me-2 flex items-center sm:hidden">
                <button
                  onClick={() =>
                    setShowingNavigationDropdown(
                      (previousState) => !previousState
                    )
                  }
                  className="inline-flex items-center justify-center p-2 rounded-md text-gray-600 hover:text-gray-800 hover:bg-gray-200 focus:outline-none transition duration-150 ease-in-out dark:text-gray-400 dark:hover:text-gray-500 dark:hover:bg-gray-100"
                >
                  <svg
                    className="h-6 w-6"
                    stroke="currentColor"
                    fill="none"
                    viewBox="0 0 24 24"
                  >
                    <path
                      className={
                        !showingNavigationDropdown ? "inline-flex" : "hidden"
                      }
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M4 6h16M4 12h16M4 18h16"
                    />
                    <path
                      className={
                        showingNavigationDropdown ? "inline-flex" : "hidden"
                      }
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M6 18L18 6M6 6l12 12"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <div
            className={
              (showingNavigationDropdown ? "block" : "hidden") + " sm:hidden"
            }
          >
            <div className="pt-2 pb-3 space-y-1">
              <ResponsiveNavLink>
                <h1 className="pb-1 text-black hover:border-black hover:border-b-2 dark:text-white dark:hover:border-white">
                  Learn
                </h1>
              </ResponsiveNavLink>

              <ResponsiveNavLink>
                <h1 className="pb-1 text-black hover:border-black hover:border-b-2 dark:text-white dark:hover:border-white">
                  Community
                </h1>
              </ResponsiveNavLink>

              <ResponsiveNavLink>
                <h1 className="pb-1 text-black hover:border-black hover:border-b-2 dark:text-white dark:hover:border-white">
                  Blog
                </h1>
              </ResponsiveNavLink>
            </div>

            <div className="pt-4 pb-1 border-t border-gray-200 dark:border-gray-700">
              <div className="px-4">
                <div className="font-medium text-base text-black dark:text-white">Name</div>
                <div className="font-medium text-sm text-gray-700 dark:text-gray-300">
                  raturu@gmail.com
                </div>
              </div>

              <div className="mt-3 space-y-1">
                <ResponsiveNavLink where={"/profile"}>
                  <h1 className="text-black dark:text-white">Profile</h1>
                </ResponsiveNavLink>

                <ResponsiveNavLink where={"/profile"}>
                  <h1 className="text-black dark:text-white">Log Out</h1>
                </ResponsiveNavLink>
              </div>
            </div>
          </div>
        </nav>
      </div>
      
    </>
  );
};

export default NavigationBar;
