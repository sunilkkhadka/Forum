import { Suspense } from "react";
import { createBrowserRouter } from "react-router-dom";

import Root from "./Root";
import { Register } from "./LazyRoutes";

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
]);

export default router;
