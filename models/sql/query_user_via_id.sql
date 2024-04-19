SELECT
  email,
  password_hash
FROM
  users
WHERE
  id = $1;