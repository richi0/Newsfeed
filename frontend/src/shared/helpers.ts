export const getHostName = (source: string): string => {
  const url = new URL(source);
  return url.hostname;
};
