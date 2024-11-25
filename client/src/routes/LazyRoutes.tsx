import React from "react";

export const Register = React.lazy(
  () => import("../features/auth/views/Register")
);

export const Login = React.lazy(() => import("../features/auth/views/Login"));

export const Post = React.lazy(() => import("../features/posts/views/Post"));

export const PostForm = React.lazy(
  () => import("../features/posts/views/PostForm")
);
