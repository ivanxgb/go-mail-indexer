import type {
  MailData,
  SearchResponseModel
} from "@/model/SearchResponseModel";
import type { SummaryModel } from "@/model/SummaryModel";

export async function fetchSearch(search: string): Promise<MailData[]> {
  const body = JSON.stringify({ search });

  const req = await baseRequest<SearchResponseModel>("search", "POST", body);
  return req.mails;
}

export async function fetchSummary(content: string): Promise<SummaryModel> {
  const body = JSON.stringify({ content });

  return await baseRequest<SummaryModel>("summary", "POST", body);
}

async function baseRequest<T>(
  path: string,
  method: string,
  body: string
): Promise<T> {
  const req = await fetch(`/api/${path}`, {
    method,
    headers: {
      "Content-Type": "application/json"
    },
    body
  });

  return (await req.json()) as T;
}
