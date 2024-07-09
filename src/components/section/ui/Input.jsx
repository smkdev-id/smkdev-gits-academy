const Input = ({ label, type, placeholder }) => {
  return (
    <label className={`form-control w-full`}>
      <div className="label">
        <span className="text-xs font-semibold lg:text-sm md:text-sm label-text">{label}</span>
      </div>
      <input
        type={type}
        placeholder={placeholder}
        className={`input input-bordered w-full focus:input-primary`}
      />
    </label>
  );
};

export default Input;
