const FormInput = ({ label, type, placeholder }) => {
  return (
    <label className="w-full max-w-xs form-control">
      <div className="label">
        <span className="label-text">{label}</span>
      </div>
      <input
        type={type}
        placeholder={placeholder}
        className="w-full input input-bordered"
      />
    </label>
  );
};

export default FormInput;
