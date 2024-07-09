import { GoArrowRight } from "react-icons/go";

const Card = ({ title, body, img, category }) => {
  return (
    <div className="relative w-full bg-white border border-gray-200 rounded-lg shadow md:w-1/4 lg:w-1/4">
      <img
        className="object-cover w-full h-48 rounded-t-lg"
        width={100}
        height={100}
        src={img}
        alt=""
      />
      <div className="p-5">
        {category && (
          <span className="absolute py-3 text-white badge badge-primary right-5 top-4">
            {category}
          </span>
        )}
        <a href="#">
          <h5 className="mb-2 text-xl font-bold tracking-tight text-gray-900">
            {title}
          </h5>
        </a>
        <p className="mb-3 text-sm font-normal text-gray-500 min-h-24">
          {body}
        </p>
        <a href="#" className="inline-block">
          <button
            className="flex items-center gap-2 py-2 font-sans text-xs font-bold text-gray-900 uppercase align-middle rounded-lg hover:underline hover:underline-offset-1"
            type="button"
          >
            Read more
            <GoArrowRight className="w-4 h-4 text-gray-900" />
          </button>
        </a>
      </div>
    </div>
  );
};

export default Card;