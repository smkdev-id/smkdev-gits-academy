/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "via.placeholder.com",
        port: "",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "smkdev.storage.googleapis.com",
        port: "",
        pathname: "/wp/**",
      },
    ],
  },
};

export default nextConfig;
