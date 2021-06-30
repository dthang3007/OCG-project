export const currency = (value) => {
  return value.toLocaleString("en-US", { style: "currency", currency: "VND" });
};
