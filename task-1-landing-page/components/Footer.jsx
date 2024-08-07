// Footer.js
"use client";
import React from "react";
import Image from "next/image";

export function Footer() {
  return (
    <footer className="w-full  bg-slate-100 py-10">
      <div className="container mx-auto px-4">
        <div className="flex flex-col justify-between items-center">
          <Image src="/smkdev-logo.png" alt="Logo" width={100} height={10} />
          <p className="mt-2 text-md ">Creating High-Caliber Digital Talent</p>
          <p className="mt-1 text-md ">
            Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec.
            Gedebage, Kota Bandung, Jawa Barat 40295
          </p>
          <div className="flex mt-3 space-x-6">
            <a
              href="https://www.youtube.com/@smkdev"
              target="_blank"
              className="hover:underline"
            >
              Youtube
            </a>
            <a
              href="https://www.instagram.com/smkdev.official/"
              target="_blank"
              className="hover:underline"
            >
              Instagram
            </a>
            <a
              href="https://t.me/+wPXVCS29UKgwNmJl"
              target="_blank"
              className="hover:underline"
            >
              Telegram
            </a>
            <a
              href="https://www.linkedin.com/company/smkdev/"
              target="_blank"
              className="hover:underline"
            >
              Linkedin
            </a>
            <a
              href="https://www.tiktok.com/@smkdev.official"
              target="_blank"
              className="hover:underline"
            >
              Tiktok
            </a>
            <a
              href="https://chat.whatsapp.com/D5gVS4FbxSc2ij1zbZVxAV"
              target="_blank"
              className="hover:underline"
            >
              Whatsapp
            </a>
          </div>
        </div>
        <div className="mt-6 text-center">
          <p>&copy; 2024 Firyal. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
}
