-- name: CreateBudget :one
INSERT INTO budget (month_year, amount)
VALUES (?, ?)
RETURNING *;

-- name: GetBudget :one
SELECT * FROM budget
WHERE month_year = ?;

-- name: UpdateBudget :exec
UPDATE budget
SET amount = ?
WHERE month_year = ?;

-- name: DeleteBudget :exec
DELETE FROM budget
WHERE month_year = ?;

-- name: ListBudget :many
SELECT * FROM budget
ORDER BY month_year;
