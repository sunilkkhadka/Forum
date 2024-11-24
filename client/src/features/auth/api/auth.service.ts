import axios from "axios";

import { IBasicFormFieldProps } from "../auth.type";

import envConfigs from "../../../shared/config/env";
import { SuccessResponse } from "../../../shared/types/api-response.type";

export const register = (user: IBasicFormFieldProps) => {
  const response = axios.post<SuccessResponse>(
    envConfigs.BACKEND_URL + "/auth/register",
    user,
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  return response;
};

export const login = (user: IBasicFormFieldProps) => {
  const response = axios.post(envConfigs.BACKEND_URL + "/auth/login", user, {
    headers: {
      "Content-Type": "application/json",
    },
    withCredentials: true,
  });

  return response;
};
