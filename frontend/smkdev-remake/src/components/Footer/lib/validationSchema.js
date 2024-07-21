// src/validationSchema.js
import * as Yup from "yup";

/**
 * Skema validasi menggunakan Yup untuk formulir kontak.
 * @type {import('yup').ObjectSchema<any>}
 */
export const validationSchema = Yup.object({
  /**
   * Nama pengguna.
   * @type {Yup.StringSchema<string>}
   */
  name: Yup.string()
    .max(50, "Name must be less than 50 characters.")
    .matches(/^[a-zA-Z\s]+$/, "Name cannot contain numbers.")
    .required("This field is required. Please input your name."),
  /**
   * Email pengguna.
   * @type {Yup.StringSchema<string>}
   */
  email: Yup.string()
    .email("Invalid email address.")
    .required("This field is required. Please input a valid email."),
  /**
   * Nomor WhatsApp pengguna.
   * @type {Yup.StringSchema<string>}
   */
  whatsapp: Yup.string()
    .matches(/^[0-9]+$/, "WhatsApp number must contain only numbers.")
    .min(10, "WhatsApp number must be at least 10 digits.")
    .required("This field is required. Please input a phone number."),
  /**
   * Institusi pengguna.
   * @type {Yup.StringSchema<string>}
   */
  institution: Yup.string().required(
    "This field is required. Please input your institution name."
  ),
  /**
   * Pesan dari pengguna.
   * @type {Yup.StringSchema<string>}
   */
  message: Yup.string()
    .max(500, "Message must be less than 500 characters.")
    .required("This field is required. Please input a message."),
});
