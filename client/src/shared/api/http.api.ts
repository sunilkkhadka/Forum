import axios from "axios";

import envConfigs from "../config/env";

export default axios.create({
  baseURL: envConfigs.BACKEND_URL,
  headers: {
    "Content-Type": "application/json",
  },
});
