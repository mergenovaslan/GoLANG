package data

type Session struct {
	ID         int
	UUID       string
	Email      string
	User_ID    int
	Created_at string
}

func (session *Session) Check() (valid bool, err error) {
	err = DB.QueryRow(
		"SELECT ID, UUID, EMAIL, USER_ID, CREATED_AT FROM SESSIONS WHERE UUID=$1", session.UUID,
	).Scan(&session.ID, &session.UUID, &session.Email, &session.User_ID, &session.Created_at)
	if err != nil {
		valid = false
		return
	}

	return session.ID != 0, nil
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid=$1"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.UUID)
	return
}

func (session *Session) User() (user User, err error) {
	user = User{}
	err = DB.QueryRow(
		"SELECT ID, UUID, NAME, EMAIL, CREATED_AT FROM USERS WHERE ID = ?", session.User_ID,
	).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func SessionDeleteAll() (err error) {
	_, err = DB.Exec("DELETE FROM SESSIONS")
	return
}
