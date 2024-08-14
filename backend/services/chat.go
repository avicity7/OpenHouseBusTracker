package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateChatRoom(users []string, groupname string) (string, error) {
	room_id := uuid.NewString()

	q := `INSERT INTO room_name(room_id, name) VALUES(@RoomID, @Groupname)`

	args := pgx.NamedArgs{
		"Groupname": groupname,
		"RoomID":    room_id,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, user := range users {
		q := `INSERT INTO chat_room (room_id, email) VALUES (@RoomID, @Email)`

		args := pgx.NamedArgs{
			"RoomID": room_id,
			"Email":  user,
		}

		_, err := config.Dbpool.Exec(context.Background(), q, args)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return room_id, nil
}

func GetChatRooms(email string) ([]structs.ChatRoom, error) {
	var chat_rooms []structs.ChatRoom

	q1 := ` 
		SELECT cr.room_id, '' as blank, cr.email, ut.name, COALESCE(timestamp, timestamp '2000-01-01 00:00:00') AS timestamp, COALESCE("from", '') AS from, COALESCE(ms.name, '') as from_name, cr.room_id, COALESCE(body, '') AS body FROM chat_room cr 
		JOIN user_table ut ON cr.email = ut.email
		FULL JOIN (
			SELECT timestamp, "from", name, room_id, body, RANK() OVER ( PARTITION BY room_id ORDER BY timestamp DESC ) 
			FROM chat_message cm
			JOIN user_table ut ON cm.from = ut.email
		) ms ON cr.room_id = ms.room_id AND "rank" = 1
		WHERE NOT cr.email = @Email
		AND cr.room_id IN (SELECT room_id FROM chat_room cr WHERE cr.email = @Email)
		AND cr.room_id IN (SELECT room_id FROM room_name WHERE name = '')
	`
	q2 := ` 
		SELECT cr.room_id, rn.name, cr.email, ut.name, COALESCE(timestamp, timestamp '2000-01-01 00:00:00') AS timestamp, COALESCE("from", '') AS from, COALESCE(ms.name, '') as from_name, cr.room_id, COALESCE(body, '') AS body FROM chat_room cr 
		JOIN user_table ut ON cr.email = ut.email
		JOIN room_name rn ON cr.room_id = rn.room_id
		FULL JOIN (
			SELECT timestamp, "from", name, room_id, body, RANK() OVER ( PARTITION BY room_id ORDER BY timestamp DESC ) 
			FROM chat_message cm
			JOIN user_table ut ON cm.from = ut.email
		) ms ON cr.room_id = ms.room_id AND "rank" = 1
		WHERE cr.email = @Email
		AND cr.room_id IN (SELECT room_id FROM chat_room cr WHERE cr.email = @Email)
		AND cr.room_id IN (SELECT room_id FROM room_name WHERE name != '')
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	rows, err := config.Dbpool.Query(context.Background(), q1, args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var chat_room structs.ChatRoom
		if err := rows.Scan(&chat_room.RoomId, &chat_room.RoomName, &chat_room.Email, &chat_room.Name, &chat_room.LatestMessage.Timestamp, &chat_room.LatestMessage.From, &chat_room.LatestMessage.FromName, &chat_room.LatestMessage.RoomId, &chat_room.LatestMessage.Body); err != nil {
			fmt.Println(err)
			return nil, err
		}
		chat_rooms = append(chat_rooms, chat_room)
	}

	rows, err = config.Dbpool.Query(context.Background(), q2, args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var chat_room structs.ChatRoom
		if err := rows.Scan(&chat_room.RoomId, &chat_room.RoomName, &chat_room.Email, &chat_room.Name, &chat_room.LatestMessage.Timestamp, &chat_room.LatestMessage.From, &chat_room.LatestMessage.FromName, &chat_room.LatestMessage.RoomId, &chat_room.LatestMessage.Body); err != nil {
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
	query := `INSERT INTO chat_message (timestamp, "from", room_id, body) VALUES (NOW(), @Email, @RoomId, @Body)`

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
		SELECT timestamp, "from", ut.name, room_id, body FROM chat_message cm
		JOIN user_table ut ON ut.email = cm.from
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
		if err := rows.Scan(&message.Timestamp, &message.From, &message.FromName, &message.RoomId, &message.Body); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func DeleteMessage(message structs.Message) error {
	query := `
		DELETE FROM chat_message cm
		WHERE "timestamp" = @Timestamp AND "from" = @Email AND room_id = @RoomId
	`

	args := pgx.NamedArgs{
		"Timestamp": message.Timestamp,
		"Email":     message.From,
		"RoomId":    message.RoomId,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.Melody.Broadcast([]byte(message.RoomId))

	return nil
}

func DeleteRoom(room_id string) error {
	q1 := `
		DELETE FROM chat_message cm
		WHERE room_id = @RoomId;
	`

	q2 := `
		DELETE FROM chat_room
		WHERE room_id = @RoomId
	`

	args := pgx.NamedArgs{
		"RoomId": room_id,
	}

	_, err := config.Dbpool.Exec(context.Background(), q1, args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = config.Dbpool.Exec(context.Background(), q2, args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.Melody.Broadcast([]byte(room_id + "del"))

	return nil
}
