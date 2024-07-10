import React from "react";
import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  Input,
  Textarea,
} from "@chakra-ui/react";
import { FieldError, UseFormRegister } from "react-hook-form";

interface FormFieldProps {
  id: string;
  label: string;
  placeholder: string;
  type?: string;
  error?: FieldError;
  register: UseFormRegister<any>;
  isTextArea?: boolean;
  validationRules: Record<string, any>;
  onBlur?: () => void;
}

const FormField: React.FC<FormFieldProps> = ({
  id,
  label,
  placeholder,
  type = "text",
  error,
  register,
  isTextArea = false,
  validationRules,
  onBlur,
}) => (
  <FormControl isInvalid={!!error} mb={4}>
    <FormLabel htmlFor={id}>{label}</FormLabel>
    {isTextArea ? (
      <Textarea
        id={id}
        placeholder={placeholder}
        {...register(id, { ...validationRules, onBlur })}
      />
    ) : (
      <Input
        id={id}
        type={type}
        placeholder={placeholder}
        {...register(id, { ...validationRules, onBlur })}
      />
    )}
    <FormErrorMessage>{error && error.message}</FormErrorMessage>
  </FormControl>
);

export default FormField;
