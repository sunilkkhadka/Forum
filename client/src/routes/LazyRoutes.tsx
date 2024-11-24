import React from "react";

export const Register = React.lazy(
  () => import("../features/auth/views/Register")
);

export const Login = React.lazy(() => import("../features/auth/views/Login"));
