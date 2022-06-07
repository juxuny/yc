import dayjs from 'dayjs'

export class Formatter {
  static convertTimestampFromMillionSeconds = (timestampInMillionSeconds: number): string => {
    return timestampInMillionSeconds ? dayjs(timestampInMillionSeconds).format('YYYY-MM-DD hh:mm:ss') : '-';
  };
}
