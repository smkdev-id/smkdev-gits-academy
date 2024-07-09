import { useFormik } from "formik";
import { validationSchema } from "./lib/validationSchema.js";

/**
 * Komponen formulir menggunakan Formik untuk mengelola state formulir dan validasi.
 * @component
 */
const ContactForm = () => {
  /**
   * Hook Formik untuk mengelola state formulir dan validasi.
   * @type {import('formik').FormikHelpers}
   */
  const formik = useFormik({
    /**
     * Nilai awal untuk bidang formulir.
     * @type {Object}
     * @property {string} name - Nama pengguna.
     * @property {string} email - Email pengguna.
     * @property {string} whatsapp - Nomor WhatsApp pengguna.
     * @property {string} institution - Institusi pengguna.
     * @property {string} message - Pesan pengguna.
     */
    initialValues: {
      name: "",
      email: "",
      whatsapp: "",
      institution: "",
      message: "",
    },

    /**
     * Skema validasi untuk bidang formulir menggunakan Yup.
     * @type {import('yup').ObjectSchema<any>}
     */
    validationSchema: validationSchema,

     /**
     * Fungsi untuk menangani pengiriman formulir.
     * @param {Object} values - Nilai formulir yang dikirim.
     * @param {string} values.name - Nama pengguna.
     * @param {string} values.email - Email pengguna.
     * @param {string} values.whatsapp - Nomor WhatsApp pengguna.
     * @param {string} values.institution - Institusi pengguna.
     * @param {string} values.message - Pesan pengguna.
     */
    onSubmit: (values) => {
      console.log(JSON.stringify(values, null, 2));
    },
  });

  return (
    <form onSubmit={formik.handleSubmit} className="w-full  ">
      <div className="mb-4">
        <label
          htmlFor="name"
          className="block text-gray-500 font-poppins text-sm font-medium mb-1"
        >
          Nama Anda
        </label>
        <input
          id="name"
          name="name"
          type="text"
          className={`bg-gray-100 border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            formik.touched.name && formik.errors.name
              ? "border-red-500 border-l-4"
              : ""
          }`}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.name}
        />
        {formik.touched.name && formik.errors.name ? (
          <p className="text-red-500 mt-2 p-2 rounded bg-custom-color-red text-xs font-poppins">
            {formik.errors.name}
          </p>
        ) : null}
      </div>

      <div className="mb-4">
        <label
          htmlFor="email"
          className="block text-gray-500 font-poppins text-sm font-medium mb-1"
        >
          Alamat Email
        </label>
        <input
          id="email"
          name="email"
          type="email"
          className={`bg-gray-100 border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            formik.touched.email && formik.errors.email
              ? "border-red-500 border-l-4"
              : ""
          }`}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.email}
        />
        {formik.touched.email && formik.errors.email ? (
          <p className="text-red-500 mt-2 p-2 rounded bg-custom-color-red text-xs font-poppins">
            {formik.errors.email}
          </p>
        ) : null}
      </div>

      <div className="mb-4">
        <label
          htmlFor="whatsapp"
          className="block text-gray-500 font-poppins text-sm font-medium mb-1"
        >
          Nomor WhatsApp
        </label>
        <input
          id="whatsapp"
          name="whatsapp"
          type="text"
          className={`bg-gray-100 border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            formik.touched.whatsapp && formik.errors.whatsapp
              ? "border-red-500 border-l-4"
              : ""
          }`}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.whatsapp}
        />
        {formik.touched.whatsapp && formik.errors.whatsapp ? (
          <p className="text-red-500 mt-2 p-2 rounded bg-custom-color-red text-xs font-poppins">
            {formik.errors.whatsapp}
          </p>
        ) : null}
      </div>

      <div className="mb-4">
        <label
          htmlFor="institution"
          className="block text-gray-500 font-poppins text-sm font-medium mb-1"
        >
          Nama Satuan Pendidikan/Instansi
        </label>
        <input
          id="institution"
          name="institution"
          type="text"
          className={`bg-gray-100 border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            formik.touched.institution && formik.errors.institution
              ? "border-red-500 border-l-4"
              : ""
          }`}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.institution}
        />
        {formik.touched.institution && formik.errors.institution ? (
          <p className="text-red-500 mt-2 p-2 rounded bg-custom-color-red text-xs font-poppins">
            {formik.errors.institution}
          </p>
        ) : null}
      </div>

      <div className="mb-4">
        <label
          htmlFor="message"
          className="block text-gray-500 font-poppins text-sm font-medium mb-1"
        >
          Pesan
        </label>
        <textarea
          id="message"
          name="message"
          className={`bg-gray-100 border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            formik.touched.message && formik.errors.message
              ? "border-red-500 border-l-4"
              : ""
          }`}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          value={formik.values.message}
        />
        {formik.touched.message && formik.errors.message ? (
          <p className="text-red-500 mt-2 p-2 rounded bg-custom-color-red text-xs font-poppins">
            {formik.errors.message}
          </p>
        ) : null}
      </div>

      <div className="flex items-center justify-between">
        <button
          type="submit"
          className="bg-custom-color-btn-send  text-white font-normal py-2 hover:shadow-xl px-10 rounded focus:outline-none focus:shadow-outline"
        >
          Send Message
        </button>
      </div>
    </form>
  );
};

export default ContactForm;
