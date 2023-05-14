export interface SearchResponseModel {
  total: Number;
  mails: MailData[];
}

export interface MailData {
  id: string;
  mail: Mail;
}

export interface Mail {
  "@timestamp": string;
  bcc: string[];
  cc: string[];
  content: string;
  date: string;
  from: string;
  "message-id": string;
  subject: string;
  to: string[];
  "x-bcc": string[];
  "x-cc": string[];
  "x-filename": string;
  "x-from": string;
  "x-to": string;
}
