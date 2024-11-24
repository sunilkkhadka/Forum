import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

import { login } from "./api/auth.service";
import { IBasicFormFieldProps } from "./auth.type";
import axios from "axios";

interface ErrorResponse {
  code: number;
  message: string;
}

export interface AuthState {
  user: {
    email: string;
  };
  isLoggedIn: boolean;
  status: "idle" | "loading" | "succeeded" | "failed";
  error: ErrorResponse | null;
}

const initialState: AuthState = {
  user: {
    email: "",
  },
  isLoggedIn: false,
  status: "idle",
  error: null,
};

export const loginUser = createAsyncThunk(
  "auth/loginUser",
  async (user: IBasicFormFieldProps, { rejectWithValue }) => {
    try {
      const response = await login(user);
      console.log("response", response);
      return response.data.response;
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        const apiError = error.response?.data.response;
        if (apiError) {
          return rejectWithValue(apiError as ErrorResponse);
        }
      }
      return rejectWithValue({
        code: 500,
        message: "An unexpected error occured",
      } as ErrorResponse);
    }
  }
);

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(loginUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.status = "succeeded";
        state.user.email = action.payload.data.email;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.status = "failed";
        state.error = (action.payload as ErrorResponse) || null;
      });
  },
});

// export const {} = authSlice.actions
export default authSlice.reducer;
