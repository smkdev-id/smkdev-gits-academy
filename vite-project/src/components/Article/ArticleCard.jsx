import PropTypes from "prop-types";

const ArticleCard = ({ image, title, description, link }) => {
  return (
    <div className="max-w-sm rounded-lg overflow-hidden shadow-lg m-4 ">
      {image ? (
        <img className="w-full h-40 object-cover" src={image} alt={title} />
      ) : (
        <div className="w-full h-40 bg-gray-200 flex items-center justify-center">
          <span className="text-gray-500">No Image</span>
        </div>
      )}
      <div className="px-6 py-4">
        <div className="font-semibold font-poppins text-md mb-2">{title}</div>
        <p className="text-gray-700 font-poppins text-sm truncate">{description}</p>
      </div>
      <div className="px-6 pt-4 pb-5">
        <a href={link} className="text-custom-blue-secondary font-poppins hover:text-custom-color-font">
          BACA LEBIH LANJUT »
        </a>
      </div>
    </div>
  );
};

ArticleCard.propTypes = {
  image: PropTypes.string,
  category: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  link: PropTypes.string.isRequired,
};
export default ArticleCard;
