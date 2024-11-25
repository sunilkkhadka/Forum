import axios from "axios";
import { toast } from "react-toastify";

import envConfigs from "../config/env";

export default axios.create({
  baseURL: envConfigs.BACKEND_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const client = axios.create({
  baseURL: envConfigs.BACKEND_URL,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

client.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error?.config;
    if (error && error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      if (originalRequest.url === "/auth/refresh") {
        setTimeout(() => {
          window.location.href = "/login";
        }, 2000);
        showToast("Session Expired! Redirecting to Login Page", "error");
        return Promise.reject(error);
      }

      try {
        await client.post("/auth/refresh");
        return client(originalRequest);
      } catch (refreshError) {
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

const showToast = (
  message: string,
  type: "success" | "error" | "info" = "info"
) => {
  toast[type](message, {
    position: "top-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
  });
};
