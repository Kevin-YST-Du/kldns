package repository

import (
	"context"
)

// CreateUser 在数据库中创建新用户并返回生成的用户 ID
func (r *AdminRepository) CreateUser(ctx context.Context, input UserWrite) (int64, error) {
	res, err := r.DB.ExecContext(
		ctx,
		"INSERT INTO users (username, email, group_id, status, password_hash, sid, points) VALUES (?, ?, ?, ?, ?, ?, 0)",
		input.Username, input.Email, input.GroupID, input.Status, input.PasswordHash, input.SID,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}