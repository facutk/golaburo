import getPagination from './util/getPagination';

const parseJson = (response) => response.json().then((json) => {
  const pagination = getPagination(response);

  return pagination ? { pagination, items: json } : json;
});

const request = (url, options) => fetch(url, options).then(parseJson);

export default request;
