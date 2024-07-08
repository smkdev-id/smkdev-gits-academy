import { FC, ReactNode } from "react";
import { Button, ButtonProps } from "@chakra-ui/react";
import Link from "next/link";

interface ButtonLinkProps extends ButtonProps {
  href?: string;
  children: ReactNode;
}

const ButtonLink: FC<ButtonLinkProps> = ({ href, children, ...props }) => {
  if (href) {
    return (
      <Link href={href} passHref>
        <Button {...props}>{children}</Button>
      </Link>
    );
  }

  return <Button {...props}>{children}</Button>;
};

export default ButtonLink;
