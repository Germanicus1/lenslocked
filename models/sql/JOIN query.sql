SELECT
	user_id,
	users.email,
	users.password_hash
FROM
	sessions
	JOIN users ON sessions.user_id = users.id
WHERE
	token_hash = $1;