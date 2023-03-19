import { AppClient } from "./generated";

export const appClient = new AppClient({
  BASE: process.env.REACT_APP_REST_URL,
});

export default appClient
