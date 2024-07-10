"use client";
import Image from "next/image";

import { FC, useEffect, useState } from "react";
import ActionMenu from "./DropdownMenu";
import ButtonLink from "./Button";
import MediaQuery from "./MediaQuery";

interface Props {}

const Topnav: FC<Props> = ({}) => {
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 30) {
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
  const menuItems = [
    { title: "Bootcamp", href: "#" },
    { title: "Expert Class", href: "#" },
    { title: "Challange", href: "#" },
  ];

  return (
    <header
      className={`${
        isScrolled ? "border-b border-gray-200 py-3" : "py-5"
      } sticky top-0 z-50 flex w-full items-center justify-center bg-white transition-all ease-in-out`}
    >
      <MediaQuery className="flex items-center justify-between">
        <div className="w-[130px]">
          <Image
            src="/assets/Logo-Smkdev.png"
            alt="logo"
            width="0"
            height="0"
            sizes="100vw"
            className="h-auto w-full"
          />
        </div>
        <div className="flex items-center gap-x-2">
          <div className="flex items-center gap-x-1">
            <ActionMenu color={"primary"} title="Learn" items={menuItems} />
            <ActionMenu color={"primary"} title="Community" href="/tgg" />
            <ActionMenu color={"primary"} title="Blog" href="/blog" />
          </div>

          <ButtonLink
            bg={"primary"}
            href="lms.smk.dev/login"
            color={"white"}
            _hover={{}}
          >
            Dashboard
          </ButtonLink>
        </div>
      </MediaQuery>
    </header>
  );
};

export default Topnav;