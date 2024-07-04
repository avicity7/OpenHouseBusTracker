package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateChatRoom(user1 string, user2 string) (string, error) {
	room_id := uuid.NewString()

	query := `INSERT INTO chat_room (room_id, user1, user2) VALUES (@RoomID, @User1, @User2)`

	args := pgx.NamedArgs{
		"RoomID": room_id,
		"User1":  user1,
		"User2":  user2,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return room_id, nil
}

func GetChatRooms(email string) ([]structs.ChatRoom, error) {
	var chat_rooms []structs.ChatRoom

	query := ` 
	SELECT room_id, ut1.name, ut2.name FROM chat_room
	JOIN user_table ut1 ON chat_room.user1 = ut1.email
	JOIN user_table ut2 ON chat_room.user2 = ut2.email
	WHERE user1 = @Email  OR user2 = @Email
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var chat_room structs.ChatRoom
		if err := rows.Scan(&chat_room.RoomId, &chat_room.User1, &chat_room.User2); err != nil {
			return nil, err
		}
		chat_rooms = append(chat_rooms, chat_room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chat_rooms, nil
}

func CreateMessage(email string, room_id string, body string) error {
	query := `INSERT INTO chat_message ("from", room_id, body) VALUES (@Email, @RoomId, @Body)`

	args := pgx.NamedArgs{
		"Email":  email,
		"RoomId": room_id,
		"Body":   body,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.Melody.Broadcast([]byte(room_id))

	return nil
}

func GetMessages(room_id string) ([]structs.Message, error) {
	var messages []structs.Message
	query := ` 
		SELECT * FROM chat_message
		WHERE room_id = @RoomId
		ORDER BY "timestamp" DESC
	`

	args := pgx.NamedArgs{
		"RoomId": room_id,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var message structs.Message
		if err := rows.Scan(&message.Timestamp, &message.From, &message.RoomId, &message.Body); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
