const Header = ({ title, description }) => {
  return (
    <div className="flex flex-col items-center w-full gap-4 pb-6">
      <h1 className="text-3xl font-bold">{title}</h1>
      {description ? <p className="text-lg italic text-center">{description}</p> : ""}
    </div>
  );
};

export default Header;
