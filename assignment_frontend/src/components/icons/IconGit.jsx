import React from 'react';

const IconGit = ({ fill = "#000000", hoverFill = "#FF0000", className, ...props }) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      x="0px"
      y="0px"
      className={className} // Menggunakan Tailwind untuk ukuran responsif
      viewBox="0 0 64 64"
      style={{ transition: 'fill 0.3s ease' }} // Menambahkan transition langsung di elemen SVG
      {...props} // Spread props untuk memungkinkan tambahan properti dinamis
    >
      <style>
        {`
          .icon-path:hover {
            fill: ${hoverFill}; // Menggunakan hoverFill untuk warna saat hover
          }
        `}
      </style>
      <path
        fill={fill} // Menambahkan atribut fill untuk mengatur warna
        className="icon-path"
        d="M32 0C14.327 0 0 14.327 0 32c0 14.138 9.168 26.085 21.888 30.293c1.6 0.293 2.193-0.693 2.193-1.527c0-0.735-0.026-2.678-0.041-5.257c-8.905 1.934-10.787-4.293-10.787-4.293c-1.456-3.693-3.556-4.677-3.556-4.677c-2.907-1.989 0.221-1.949 0.221-1.949c3.218 0.221 4.91 3.301 4.91 3.301c2.858 4.897 7.491 3.481 9.314 2.662c0.287-2.073 1.117-3.482 2.031-4.281c-7.113-0.812-14.586-3.556-14.586-15.84c0-3.493 1.248-6.353 3.301-8.591c-0.331-0.811-1.433-4.078 0.289-8.497c0 0 2.688-0.857 8.8 3.277c2.554-0.709 5.291-1.063 8.012-1.073c2.718 0.01 5.457 0.364 8.015 1.073c6.108-4.134 8.793-3.277 8.793-3.277c1.728 4.42 0.625 7.686 0.305 8.497c2.057 2.238 3.297 5.098 3.297 8.591c0 12.312-7.489 15.019-14.618 15.821c1.146 0.989 2.168 2.948 2.168 5.949c0 4.293-0.038 7.749-0.038 8.795c0 0.84 0.58 1.834 2.211 1.527C54.84 58.085 64 46.138 64 32C64 14.327 49.673 0 32 0z"
      ></path>
    </svg>
  );
};

export default IconGit;
