-- name: CreateExpense :one
INSERT INTO expense (description, amount, category)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expense
WHERE id = ?;

-- name: UpdateExpense :exec
UPDATE expense
SET description = ?,
amount = ?,
category = ?
WHERE id = ?;

-- name: DeleteExpense :exec
DELETE FROM expense
WHERE id = ?;

-- name: ListExpense :many
SELECT * FROM expense
ORDER BY id;

-- name: ListFilteredExpense :many
SELECT * FROM expense 
WHERE strftime('%Y', created_at) = CAST(? AS TEXT)
AND strftime('%m', created_at) = CAST(? AS TEXT);
