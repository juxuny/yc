// Code generated by yc@{{.CommandLineVersion}}. DO NOT EDIT.
import { {{.ClassName}} } from './api';

export let prefix = "/api";

export const updateApiPrefix = (apiPrefix: string) => {
  prefix = apiPrefix;
}

export {
  {{.ClassName}}
}
