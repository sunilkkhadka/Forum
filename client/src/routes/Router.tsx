import { Suspense } from "react";
import { createBrowserRouter } from "react-router-dom";

import Root from "./Root";
import { Login, Post, PostForm, Register } from "./LazyRoutes";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
  },
  {
    path: "/register",
    element: (
      <Suspense fallback={<h1>Loading...</h1>}>
        <Register />
      </Suspense>
    ),
  },
  {
    path: "/login",
    element: (
      <Suspense fallback={<h1>Loading...</h1>}>
        <Login />
      </Suspense>
    ),
  },
  {
    path: "/post",
    element: (
      <Suspense fallback={<h1>Loading...</h1>}>
        <Post />
      </Suspense>
    ),
  },
  {
    path: "/post/create",
    element: (
      <Suspense fallback={<h1>Post Form Loading...</h1>}>
        <PostForm />
      </Suspense>
    ),
  },
]);

export default router;
