// import { Configuration } from "./types/typescript-axios/configuration";
import { Configuration } from "./types/typescript-fetch";
// import { TodoApi } from "./types/typescript-axios/api";
import { TodoApi } from './types/typescript-fetch'

const REACT_APP_REST_URL = process.env.REACT_APP_REST_URL

const config = new Configuration({
  basePath: REACT_APP_REST_URL,
});

export const todoApi = new TodoApi(config);
