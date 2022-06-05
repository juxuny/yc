export const timestampInMinute = (): string => {
  return (new Date().getTime() / 1000 / 60).toFixed(0);
};

export const convertToQueryParams = (obj = {}): string => {
  let result = "";
  for (const k in obj) {
    if (result !== "") result += '&';
    result += k + '=' + obj[k];
  }
  return result;
}
