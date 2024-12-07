// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createPet = `-- name: CreatePet :execresult
INSERT INTO pets (
    name, type_id, age, user_id
) VALUES (
    ?, ?, ?, ?
)
`

type CreatePetParams struct {
	Name   string `json:"name"`
	TypeID int32  `json:"type_id"`
	Age    int32  `json:"age"`
	UserID int32  `json:"user_id"`
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPet,
		arg.Name,
		arg.TypeID,
		arg.Age,
		arg.UserID,
	)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    name, email
) VALUES (
    ?, ?
)
`

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Name, arg.Email)
}

const deletePet = `-- name: DeletePet :exec
DELETE FROM pets
WHERE id = ?
`

func (q *Queries) DeletePet(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePet, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getPet = `-- name: GetPet :one
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id
WHERE p.id = ? LIMIT 1
`

type GetPetRow struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	TypeName string `json:"type_name"`
	Age      int32  `json:"age"`
	UserID   int32  `json:"user_id"`
}

func (q *Queries) GetPet(ctx context.Context, id int32) (GetPetRow, error) {
	row := q.db.QueryRowContext(ctx, getPet, id)
	var i GetPetRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TypeName,
		&i.Age,
		&i.UserID,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const listPetTypes = `-- name: ListPetTypes :many
SELECT id, name FROM pet_types
`

func (q *Queries) ListPetTypes(ctx context.Context) ([]PetType, error) {
	rows, err := q.db.QueryContext(ctx, listPetTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PetType
	for rows.Next() {
		var i PetType
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPets = `-- name: ListPets :many
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id
`

type ListPetsRow struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	TypeName string `json:"type_name"`
	Age      int32  `json:"age"`
	UserID   int32  `json:"user_id"`
}

func (q *Queries) ListPets(ctx context.Context) ([]ListPetsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPetsRow
	for rows.Next() {
		var i ListPetsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.TypeName,
			&i.Age,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPetsByUser = `-- name: ListPetsByUser :many
SELECT 
    p.id,
    p.name,
    pt.name as type_name,
    p.age,
    p.user_id
FROM pets p
JOIN pet_types pt ON p.type_id = pt.id
WHERE p.user_id = ?
`

type ListPetsByUserRow struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	TypeName string `json:"type_name"`
	Age      int32  `json:"age"`
	UserID   int32  `json:"user_id"`
}

func (q *Queries) ListPetsByUser(ctx context.Context, userID int32) ([]ListPetsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listPetsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPetsByUserRow
	for rows.Next() {
		var i ListPetsByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.TypeName,
			&i.Age,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSameAgePets = `-- name: ListSameAgePets :many
SELECT id, name, type_id, age, user_id FROM pets WHERE age = ?
`

func (q *Queries) ListSameAgePets(ctx context.Context, age int32) ([]Pet, error) {
	rows, err := q.db.QueryContext(ctx, listSameAgePets, age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pet
	for rows.Next() {
		var i Pet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.TypeID,
			&i.Age,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePet = `-- name: UpdatePet :exec
UPDATE pets
SET name = ?, type_id = ?, age = ?
WHERE id = ?
`

type UpdatePetParams struct {
	Name   string `json:"name"`
	TypeID int32  `json:"type_id"`
	Age    int32  `json:"age"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdatePet(ctx context.Context, arg UpdatePetParams) error {
	_, err := q.db.ExecContext(ctx, updatePet,
		arg.Name,
		arg.TypeID,
		arg.Age,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = ?, email = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Email, arg.ID)
	return err
}