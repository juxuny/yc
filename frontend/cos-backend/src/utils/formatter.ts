import dayjs from 'dayjs'

export class Formatter {
  static convertTimestampFromMillionSeconds = (timestampInMillionSeconds: number): string => {
    return dayjs(timestampInMillionSeconds).format('YYYY-MM-DD hh:mm:ss');
  };
}
