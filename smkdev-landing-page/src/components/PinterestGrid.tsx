import React, { FC } from "react";

interface Props {
  dataTestimoni: {
    text: string;
    nama: string;
    asal: string;
  }[];
}

const PinterestGrid: FC<Props> = ({ dataTestimoni }) => {
  return (
    <div className="h-fit columns-2 gap-5 lg:columns-3">
      {dataTestimoni.map((i, index) => (
        <div
          key={index}
          className="mb-5 inline-block space-y-5 rounded-lg border-gray-400 p-7 shadow-xl"
        >
          <p
            className="text-base"
            dangerouslySetInnerHTML={{ __html: `&quot;${i.text}&quot;` }}
          ></p>

          <div>
            <p className="font-semibold">{i.nama}</p>
            <p className="italic">{i.asal}</p>
          </div>
        </div>
      ))}
    </div>
  );
};

export default PinterestGrid;
