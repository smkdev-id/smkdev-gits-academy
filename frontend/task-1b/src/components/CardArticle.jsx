import React from "react";

export const CardArticle = ({ title, desc, img }) => {
  return (
    <div className="w-full h-[70vh] bg-white rounded-[10px] cursor-pointer drop-shadow-[5px_5px_15px_rgba(39,93,132,0.05)]">
      <div className="w-full h-[35%] bg-gray-400 bg-cover rounded-t-[10px]">
        <img src={`${img}`} alt="" className="w-full h-full object-cover" />
      </div>
      <div className="h-[65%] p-[30px]">
        <p className="h-[20%] text-[20px] font-semibold overflow-hidden">{title}</p>
        <p className="h-[63%] sm:h-[60%] text-[#444444] my-[20px] overflow-hidden">{desc}</p>
        <p className="text-[14px] text-[#1C3965] font-semibold cursor-pointer">
          BACA LEBIH LANJUT »
        </p>
      </div>
    </div>
  );
};
