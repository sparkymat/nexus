import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

dayjs.extend(utc);

export class Object {
  public id: string;

  public createdAt: dayjs.Dayjs;

  public updatedAt: dayjs.Dayjs;

  public name: string;

  public isTemplate: boolean;

  // eslint-disable-next-line no-use-before-define
  public template?: Object;

  constructor(json: Object) {
    if (!json) {
      return;
    }

    this.id = json.id;
    this.createdAt = dayjs.utc(json.createdAt);
    this.updatedAt = dayjs.utc(json.updatedAt);
    this.name = json.name;
    this.isTemplate = json.isTemplate;

    if (json.template) {
      this.template = new Object(json.template);
    }
  }
}
