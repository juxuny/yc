export const timestampInMinute = (): string => {
  return (new Date().getTime() / 1000 / 60).toFixed(0);
};
