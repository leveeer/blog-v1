const formatDate = function(value) {
  console.log(value);
  const date = new Date(value);
  const y = date.getFullYear();
  let MM = date.getMonth() + 1;
  MM = MM < 10 ? ("0" + MM) : MM;
  let d = date.getDate();
  d = d < 10 ? ("0" + d) : d;
  return y + "-" + MM + "-" + d;
};
export default formatDate;