SELECT
  user_id
FROM
  sessions
WHERE
  token_hash = $1;