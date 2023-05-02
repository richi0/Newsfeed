export const minutesAgo = (date: string): string => {
  const now = Date.now();
  const then = new Date(date);
  const diff = now.valueOf() - then.valueOf();
  const minutes = Math.round(diff / (1000 * 60));
  return minutes.toString();
};
