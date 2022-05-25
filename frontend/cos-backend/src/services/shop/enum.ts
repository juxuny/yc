// @ts-ignore
/* eslint-disable */

const reverseMap = (m: {[key: string]: number}) => {
  const ret: {[key: number]: string} = {};
  for(let k in m) {
    ret[m[k]] = k;
  }
  return ret;
}

export const StatusEnumMap: {[key: string]: API.EnableStatus} = {
  all: -1,
  enable: 1,
  disable: 0,
}

export const StatusEnumMapReverse: Record<number, string> = reverseMap(StatusEnumMap);
