export const getInitialsName = (name?: string) => {
  const words = name?.split(" ");
  return `${words?.[0]?.[0].toUpperCase() ?? "N"}${
    words?.[1]?.[0].toUpperCase() ?? "N"
  }`;
};

export const cleanEnronFormat = (text?: string) => {
  if (text === undefined) return "";

  const identifierRegex = /<\/O=ENRON\/OU=NA\/CN=RECIPIENTS\/CN=[^>]+>/g;
  return text.replace(identifierRegex, "");
};

export const parseDate = (date?: string) => {
  const parsedDate = new Date(date ?? "15/05/2023");
  return parsedDate.toLocaleDateString("en-US", {
    day: "numeric",
    month: "short",
    year: "numeric"
  });
};

export const debounce = (func: Function, delay: number = 400) => {
  let timeout: NodeJS.Timeout;
  return (...args: any) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => func(...args), delay);
  };
};
