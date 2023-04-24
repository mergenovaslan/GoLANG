package data

import (
	"strings"
	"time"
)

type Receipt struct {
	ID          int
	User_ID     int
	Name        string
	Photo       string
	Duration    int
	Instruction string
	CreatedAt   string
	Ingredients []Ingredient
}

func (receipt *Receipt) Create() (err error) {
	st, err := DB.Prepare("INSERT INTO RECEIPTS(USER_ID, NAME, PHOTO, DURATION, INSTRUCTION, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID, CREATED_AT")
	if err != nil {
		//danger method
		return
	}
	defer st.Close()
	err = st.QueryRow(
		receipt.User_ID, receipt.Name, receipt.Photo, receipt.Duration, receipt.Instruction, time.Now(),
	).Scan(
		&receipt.ID, &receipt.CreatedAt,
	)
	for _, ingredient := range receipt.Ingredients {
		err = ingredient.Create()
		if err != nil {
			return
		}
	}
	return
}

func (receipt *Receipt) Delete() (err error) {
	st, err := DB.Prepare("DELETE FROM RECEIPT WHERE ID=$1")
	if err != nil {
		return
	}
	defer st.Close()
	_, err = st.Exec(receipt.ID)
	return
}
func (receipt *Receipt) Update() (err error) {
	if existsUsernameNotID(receipt.Name, receipt.ID) {
		//danger method
		return
	}
	st, err := DB.Prepare("UPDATE USERS SET INSTRUCTION = $1 WHERE ID = $2")
	if err != nil {
		return
	}
	defer st.Close()
	_, err = st.Exec(receipt.Instruction, receipt.ID)
	return
}

func ReceiptsByUserID(userID int) (receipt Receipt, err error) {
	err = DB.QueryRow("SELECT ID, USER_ID, NAME, PHOTO, DURATION, INSTRUCTION FROM RECEIPTS WHERE USER_ID=$1", userID).Scan(
		&receipt.ID,
		&receipt.User_ID,
		&receipt.Name,
		&receipt.Photo,
		&receipt.Duration,
		&receipt.Instruction,
	)
	return
}

func AllReceipts() (receipts []Receipt, err error) {
	rows, err := DB.Query("SELECT ID, USER_ID, NAME, PHOTO, DURATION, INSTRUCTION, CREATED_AT FROM RECEIPTS")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var receipt Receipt
		err = rows.Scan(
			&receipt.ID, &receipt.User_ID, &receipt.Name, &receipt.Photo, &receipt.Duration, &receipt.Instruction,
			&receipt.CreatedAt,
		)
		receipt.CreatedAt = strings.Split(receipt.CreatedAt, "T")[0]
		if err != nil {
			return
		}
		receipt.Ingredients, err = IngredientsByReceiptID(receipt.ID)
		if err != nil {
			return
		}
		receipts = append(receipts, receipt)
	}
	return
}
