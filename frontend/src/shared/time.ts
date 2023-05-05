export const timeAgo = (date: string): string => {
  const now = Date.now();
  const then = new Date(date);
  const diff = now.valueOf() - then.valueOf();
  const minutes = Math.round(diff / (1000 * 60));
  if (minutes < 1 || minutes === 1) {
    return "One minute ago";
  }
  if (minutes > 1 && minutes < 60) {
    return `${minutes} minutes ago`;
  }
  if (minutes > 60 * 24) {
    const days = Math.round(minutes / (60 * 24));
    if (days < 1 || days === 1) {
      return "One day ago";
    }
    return `${days} days ago`;
  }
  if (minutes > 60) {
    const hours = Math.round(minutes / 60);
    if (hours < 1 || hours === 1) {
      return "One hour ago";
    }
    return `${hours} hours ago`;
  }
  return `${minutes} minutes ago`;
};
