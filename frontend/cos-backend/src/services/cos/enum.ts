// @ts-ignore
/* eslint-disable */

const reverseMap = (m: { [key: string]: number }) => {
  const ret: { [key: number]: string } = {};
  for (let k in m) {
    ret[m[k]] = k;
  }
  return ret;
};

export const StatusEnumMap: { [key: string]: API.EnableStatus } = {
  all: -1,
  enable: 1,
  disable: 0,
};

export const StatusEnumMapReverse: Record<number, string> = reverseMap(StatusEnumMap);

export const AccountTypeEnumMap: { [key: string]: API.AccountType } = {
  unknown: 0,
  password: 1,
};

export const AccountTypeEnumMapReserve: Record<number, string> = reverseMap(AccountTypeEnumMap);


export const ValueTypeEnumMap: { [key: string]: API.KeyValue.ValueType } = {
  number: 1,
  string: 2,
};

export const ValueTypeEnumMapReverse: Record<number, string> = reverseMap(ValueTypeEnumMap);
