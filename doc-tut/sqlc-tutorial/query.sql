-- name: GetAuthor :one
select * from authors
where id = $1 limit 1;

-- name: ListAuthors :many
select * from authors
order by name;

-- name: CreateAuthor :one
insert into authors (
    name, bio
) values (
    $1, $2
)
RETURNING *;

-- name: UpdateAuthor :one
update authors
    set name = $2,
    bio = $3
where id = $1
RETURNING *;

-- name: DeleteAuthor :exec
delete from authors
where id = $1;
