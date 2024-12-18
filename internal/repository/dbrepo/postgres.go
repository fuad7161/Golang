package dbrepo

import (
	"context"
	"errors"
	"github.com/fuad7161/Golang/tree/Project/Bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newID int
	stmt := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) 
 		values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

// InsertRoomRestriction insert room restriction
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
			values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)
	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID, and false if no availability
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select count(id) from room_restrictions where room_id = $1 and $2 < end_date and $3 > start_date`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	var numRows int

	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any for given date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select r.id , r.room_name
				from rooms r 
				where r.id not in (select rr.room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date)`

	var rooms []models.Room
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetRoomID get a room by id
func (m *postgresDBRepo) GetRoomID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room

	query := `select id, room_name, created_at, updated_at from rooms where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)

	if err != nil {
		return room, err
	}
	return room, nil
}

// GetUserByID returns a user by ID
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, first_name, last_name, email, password,access_level, created_at, updated_at from users where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)
	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.AccessLevel, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser update a user in the database
func (m *postgresDBRepo) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5 where id = $6`
	_, err := m.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.AccessLevel, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Authenticate authenticate a user in db
func (m *postgresDBRepo) Authenticate(email, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string
	row := m.DB.QueryRowContext(ctx, "select id, password from users where email = $1", email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return 0, "", errors.New("invalid email or password")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPassword, nil
}

// AllReservations return the slice of all information
func (m *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var reservations []models.Reservation
	query := `SELECT     r.id AS reservation_id,     r.first_name,     r.last_name,     r.email,     r.phone,     
    r.start_date,     r.end_date,     r.room_id,     r.created_at,     r.updated_at, r.processed, rm.id AS room_id,  rm.room_name 
	FROM     reservations r LEFT JOIN     rooms rm ON     r.room_id = rm.id ORDER BY     r.start_date ASC;`
	rows, err := m.DB.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return reservations, err
	}
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}
	if err := rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

// AllReservations return the slice of all information
func (m *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var reservations []models.Reservation
	query := `SELECT     r.id AS reservation_id,     r.first_name,     r.last_name,     r.email,     r.phone,     
    r.start_date,     r.end_date,     r.room_id,     r.created_at,     r.updated_at, rm.id AS room_id,  rm.room_name 
	FROM     reservations r LEFT JOIN     rooms rm ON     r.room_id = rm.id where r.processed = 0 ORDER BY     r.start_date ASC;`
	rows, err := m.DB.QueryContext(ctx, query)
	//fmt.Println(rows)
	defer rows.Close()
	if err != nil {
		return reservations, err
	}
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Room.ID,
			&i.Room.RoomName,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}
	if err := rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

// GetReservationByID return single id info
func (m *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var reservation models.Reservation
	query := `SELECT     r.id AS reservation_id,     r.first_name,     r.last_name,     r.email,     r.phone,     
    r.start_date,     r.end_date,     r.room_id,     r.created_at,     r.updated_at, r.processed, rm.id AS room_id,  rm.room_name 
	FROM     reservations r LEFT JOIN     rooms rm ON     r.room_id = rm.id where r.id = $1;`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&reservation.ID,
		&reservation.FirstName,
		&reservation.LastName,
		&reservation.Email,
		&reservation.Phone,
		&reservation.StartDate,
		&reservation.EndDate,
		&reservation.RoomID,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
		&reservation.Processed,
		&reservation.Room.ID,
		&reservation.Room.RoomName,
	)
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

func (m *postgresDBRepo) UpdateReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
UPDATE reservations
SET first_name = $1,
    last_name = $2,
    email = $3,
    phone = $4
WHERE id = $5;

`
	_, err := m.DB.ExecContext(ctx, query, res.FirstName, res.LastName, res.Email, res.Phone, res.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReservation delete a reservation
func (m *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from reservations where id = $1;`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProcessedForReservation update the process for reservation
func (m *postgresDBRepo) UpdateProcessedForReservation(id, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update reservations set process = $1 where id = $2;`
	_, err := m.DB.ExecContext(ctx, query, processed, id)
	if err != nil {
		return err
	}
	return nil
}

// InsertUser to insert a user
func (m *postgresDBRepo) InsertUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		INSERT INTO users (first_name, last_name, email, password, access_level,created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);
`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = m.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, hashedPassword, user.AccessLevel, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) ShowAllUser() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select u.id, u.first_name, u.last_name, u.email, u.password, u.access_level, u.created_at, u.updated_at from users u`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.AccessLevel,
			&u.CreatedAt,
			&u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *postgresDBRepo) AllRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select r.id, r.room_name, r.created_at, r.updated_at from rooms r order by r.room_name`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var romms []models.Room
	for rows.Next() {
		var u models.Room
		err := rows.Scan(
			&u.ID,
			&u.RoomName,
			&u.CreatedAt,
			&u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		romms = append(romms, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return romms, nil
}

// GetRestrictionsForRoomByDate return restricon
func (m *postgresDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var restrictions []models.RoomRestriction
	query := `
	select rr.id, rr.reservation_id, rr.restriction_id, rr.room_id , rr.start_date , rr.end_date from room_restrictions rr 
where $1 < rr.end_date  and $2 >= rr.start_date   and room_id = $3
`
	rows, err := m.DB.QueryContext(ctx, query, start, end, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var restriction models.RoomRestriction
		err := rows.Scan(
			&restriction.ID,
			&restriction.ReservationID,
			&restriction.RestrictionID,
			&restriction.RoomID,
			&restriction.StartDate,
			&restriction.EndDate,
		)
		if err != nil {
			return nil, err
		}
		restrictions = append(restrictions, restriction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return restrictions, nil
}

//
//func (m *postgresDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	query := `insert into room_restrictions (start_date, end_date, room_id) values ($1, $2, $3)`
//}
