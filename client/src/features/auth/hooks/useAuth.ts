import { toast } from "react-toastify";
import { useMutation } from "@tanstack/react-query";

import { register } from "../api/auth.service";
import { IBasicFormFieldProps } from "../auth.type";

export const useRegisterUser = () => {
  return useMutation({
    mutationFn: (user: IBasicFormFieldProps) => register(user),
    onSuccess: ({ data, status }) => {
      if (status == 200) {
        return toast.success(data.response.message);
      }
    },
  });
};
