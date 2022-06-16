import { cos } from './api';
import * as typing from './typing';

export let prefix = "/api";

export const updateApiPrefix = (apiPrefix: string) => {
  prefix = apiPrefix;
}


export {
  cos,
  typing
}
