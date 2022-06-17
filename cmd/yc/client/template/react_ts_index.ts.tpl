import { {{.PackageAlias}} } from './api';

export let prefix = "/api";

export const updateApiPrefix = (apiPrefix: string) => {
  prefix = apiPrefix;
}

export {
  {{.PackageAlias}}
}
