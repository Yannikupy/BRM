import {ContactData} from "./ContactData";

export interface ContactResponse {
  data: ContactData[],
  error: string[]
}
