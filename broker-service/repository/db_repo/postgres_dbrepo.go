package dbrepo

import (
	pb "broker/proto/generated"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllChatMessages() ([]*pb.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select
			id, content
		from
			chat_messages
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var chatMessages []*pb.ChatMessage

	for rows.Next() {
		var chatMessage pb.ChatMessage
		err := rows.Scan(
			&chatMessage.Id,
			&chatMessage.Content,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		chatMessages = append(chatMessages, &chatMessage)
	}
	return chatMessages, nil
}

func (m *PostgresDBRepo) CreateMessage(msg *pb.ChatMessage) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
			insert into chat_messages
				(content, username, created_at)
			values	
				($1,$2, $3) 
			returning id
	`
	var NewID int
	err := m.DB.QueryRowContext(ctx, stmt,
		msg.Content,
		msg.Username,
		time.Now(),
	).Scan(&NewID)
	if err != nil {
		return 0, err
	}
	return NewID, nil

}
