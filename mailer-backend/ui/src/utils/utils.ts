export const getInitialsName = (name?: string) => {
  const words = name?.split(' ')
  return `${words?.[0]?.[0].toUpperCase() ?? 'N'}${words?.[1]?.[0].toUpperCase() ?? 'N'}`
}

export const debounce = (func: Function, delay: number = 400) => {
  let timeout: NodeJS.Timeout;
  return (...args: any) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => func(...args), delay);
  };
};