import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "pub-8227eb78753c488897042f87b7c55bca.r2.dev"
      }
    ]
  }
};

export default nextConfig;
