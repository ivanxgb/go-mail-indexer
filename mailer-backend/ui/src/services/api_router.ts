import type {
  MailData,
  SearchResponseModel
} from "@/model/SearchResponseModel";

export async function search(search: string): Promise<MailData[]> {
  const body = JSON.stringify({ search });

  const req = await baseRequest<SearchResponseModel>("search", "POST", body);
  return req.mails;
}

async function baseRequest<T>(
  path: string,
  method: string,
  body: string
): Promise<T> {
  console.log("baseRequest", path, method, body);
  const req = await fetch(`/api/${path}`, {
    method,
    headers: {
      "Content-Type": "application/json"
    },
    body
  });

  return (await req.json()) as T;
}
