import axios from "axios";

import { IBasicFormFieldProps } from "../auth.type";

import envConfigs from "../../../shared/config/env";

export const register = (user: IBasicFormFieldProps) => {
  const response = axios.post(
    envConfigs.BACKEND_URL + "/api/v1/auth/register",
    user,
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  return response;
};
