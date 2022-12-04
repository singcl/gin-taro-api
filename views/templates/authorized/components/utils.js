export function getTagType(method) {
  let tagType = 'default';
  switch (method) {
    case 'GET':
      tagType = 'primary';
      break;
    case 'POST':
      tagType = 'success';
      break;
    case 'DELETE':
      tagType = 'error';
      break;
    case 'PUT':
      tagType = 'warning';
      break;
    case 'PATCH':
      tagType = 'info';
      break;
    default:
      break;
  }

  return tagType;
}
