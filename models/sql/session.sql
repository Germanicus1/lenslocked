INSERT INTO sessions (user_id, token_hash)
values($1, $2)
returning id;