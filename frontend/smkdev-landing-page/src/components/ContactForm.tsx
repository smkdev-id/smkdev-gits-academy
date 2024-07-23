"use client";
import React from "react";
import { useForm } from "react-hook-form";
import ButtonLink from "./Button";
import FormField from "./FormField";

interface IFormInput {
  name: string;
  email: string;
  phone: string;
  institution: string;
  message: string;
}

const ContactForm: React.FC = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
    trigger,
  } = useForm<IFormInput>({
    mode: "onBlur",
  });

  const onSubmit = (values: IFormInput) => {
    console.log(values);
  };

  const handleBlur = async (field: keyof IFormInput) => {
    await trigger(field);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="w-full">
      <FormField
        id="name"
        label="Nama Anda"
        placeholder="Nama Anda"
        error={errors.name}
        register={register}
        validationRules={{ required: "Nama is required" }}
        onBlur={() => handleBlur("name")}
      />

      <FormField
        id="email"
        label="Alamat Email"
        placeholder="Alamat Email"
        type="email"
        error={errors.email}
        register={register}
        validationRules={{ required: "Email is required" }}
        onBlur={() => handleBlur("email")}
      />

      <FormField
        id="phone"
        label="Nomor WhatsApp"
        placeholder="Nomor WhatsApp"
        error={errors.phone}
        register={register}
        validationRules={{ required: "Phone number is required" }}
        onBlur={() => handleBlur("phone")}
      />

      <FormField
        id="institution"
        label="Nama Satuan Pendidikan/Instansi"
        placeholder="Nama Satuan Pendidikan/Instansi"
        error={errors.institution}
        register={register}
        validationRules={{ required: "Institution name is required" }}
        onBlur={() => handleBlur("institution")}
      />

      <FormField
        id="message"
        label="Pesan"
        placeholder="Pesan"
        error={errors.message}
        register={register}
        validationRules={{ required: "Message is required" }}
        isTextArea={true}
        onBlur={() => handleBlur("message")}
      />

      <ButtonLink
        mt={5}
        bg={"primary"}
        color={"white"}
        type="submit"
        _hover={{}}
      >
        Send Message
      </ButtonLink>
    </form>
  );
};

export default ContactForm;
