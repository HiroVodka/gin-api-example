-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :execresult
INSERT INTO users (
    name, email
) VALUES (
    ?, ?
);

-- name: UpdateUser :exec
UPDATE users
SET name = ?, email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetPet :one
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id
WHERE p.id = ? LIMIT 1;

-- name: ListPets :many
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id;

-- name: ListPetsByUser :many
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id
WHERE p.user_id = ?;

-- name: CreatePet :execresult
INSERT INTO pets (
    name, type_id, age, user_id
) VALUES (
    ?, ?, ?, ?
);

-- name: UpdatePet :exec
UPDATE pets
SET name = ?, type_id = ?, age = ?
WHERE id = ?;

-- name: DeletePet :exec
DELETE FROM pets
WHERE id = ?;

-- name: ListPetTypes :many
SELECT * FROM pet_types; 

-- name: ListSameAgePets :many
SELECT * FROM pets WHERE age = ?;