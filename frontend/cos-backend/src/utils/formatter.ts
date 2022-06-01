
import moment from 'moment'

export class Formatter {
  static convertTimestampFromMillionSeconds = (timestampInMillionSeconds: number): string => {
    return moment(timestampInMillionSeconds).format('YYYY-MM-DD hh:mm:ss');
  }
}
