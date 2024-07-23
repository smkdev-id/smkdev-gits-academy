"use client";
import { FC, useRef } from "react";
import {
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Button,
  useDisclosure,
  ButtonProps,
} from "@chakra-ui/react";
import Link from "next/link";
import ButtonLink from "./Button";

interface ActionMenuProps extends ButtonProps {
  title: string;
  href?: string;
  items?:
    | string
    | {
        title: string;
        href: string;
      }[];
}

const ActionMenu: FC<ActionMenuProps> = ({ items, title, href, ...props }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const menuRef = useRef<HTMLDivElement>(null);

  const isSingleItem = typeof items === "string" || items === undefined;
  const hoverMode = !isSingleItem;

  const renderMenuItem = (item: string, href?: string) => {
    return (
      <MenuItem key={item}>
        <Link href={`/${href?.toLowerCase()}`}>{item}</Link>
      </MenuItem>
    );
  };

  const renderMenuItems = () => {
    if (isSingleItem) {
      return (
        <ButtonLink
          bg={"transparent"}
          href={href}
          color={"primary"}
          _hover={{
            bg: "rgba(28, 57, 101, 0.2)",
          }}
        >
          {title}
        </ButtonLink>
      );
    }

    return (
      <Menu isOpen={hoverMode ? isOpen : undefined}>
        <div
          className="remove-border"
          ref={menuRef}
          onMouseEnter={hoverMode ? onOpen : undefined}
          onMouseLeave={hoverMode ? onClose : undefined}
        >
          <MenuButton
            className="remove-border"
            {...props}
            as={Button}
            onClick={!hoverMode ? (isOpen ? onClose : onOpen) : undefined}
            bg={"transparent"}
            _hover={{
              bg: "rgba(28, 57, 101, 0.2)",
            }}
          >
            {title}
          </MenuButton>
          <MenuList>
            {items.map((item, index) => renderMenuItem(item.title, item.href))}
          </MenuList>
        </div>
      </Menu>
    );
  };

  return renderMenuItems();
};

export default ActionMenu;
