import styles from "../styles";
const NewFeatures = ({ name, message }) => (
  <div className="flex-1 flex flex-col sm:max-w-[250px] min-w-[210px]">
    <h1 className="mt-[26px] font-bold text-[24px] leading-[30px] text-[#ebb000]">
      {name}
    </h1>
    <p className="flex-1 mt-[16px] font-normal text-[18px] text-[#00bd00] leading-[32px]">
      {message}
    </p>
  </div>
);

export default NewFeatures;
