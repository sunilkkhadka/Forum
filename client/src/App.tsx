import { Provider } from "react-redux";
import { ToastContainer } from "react-toastify";
import { RouterProvider } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import "react-toastify/dist/ReactToastify.css";

import router from "./routes/Router";
import { store } from "./store/store";

const queryClient = new QueryClient();

function App() {
  return (
    <main>
      <Provider store={store}>
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={router} />
          <ToastContainer
            position="bottom-right"
            limit={4}
            autoClose={5000}
            closeOnClick
            pauseOnFocusLoss
            draggable
            pauseOnHover
          />
        </QueryClientProvider>
      </Provider>
    </main>
  );
}

export default App;
