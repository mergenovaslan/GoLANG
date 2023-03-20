package data

type Ingredient struct {
	ID        int
	Name      string
	ReceiptID int
	Amount    int
	Unit      string
}

func (ingredient *Ingredient) Create() (err error) {
	st, err := DB.Prepare("INSERT INTO INGREDIENTS(NAME, RECEIPT_ID, AMOUNT, UNIT) VALUES ($1, $2, $3, $4) RETURNING ID")
	if err != nil {
		return
	}
	defer st.Close()
	err = st.QueryRow(ingredient.Name, ingredient.ReceiptID, ingredient.Amount, ingredient.Unit).Scan(&ingredient.ID)
	return
}

func (ingredient *Ingredient) Delete() (err error) {
	stmt, err := DB.Prepare("DELETE FROM INGREDIENTS WHERE ID=$1")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(ingredient.ID)
	return
}

func (ingredient *Ingredient) Update(i Ingredient) (err error) {
	stmt, err := DB.Prepare("UPDATE INGREDIENTS SET NAME=$1, AMOUNT=$2, UNIT=$3 WHERE ID=$4")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(ingredient.Name, ingredient.Amount, ingredient.Unit, ingredient.ID)
	return
}

func DeleteIngredientByID(id int) (err error) {
	stmt, err := DB.Prepare("DELETE FROM INGREDIENTS WHERE ID=$1")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return
}

func IngredientsByReceiptID(id int) (ingredients []Ingredient, err error) {
	rows, err := DB.Query("SELECT ID, NAME, RECEIPT_ID, AMOUNT, UNIT FROM INGREDIENTS WHERE RECEIPT_ID=$1", id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ingredient Ingredient
		err = rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.ReceiptID, &ingredient.Amount, &ingredient.Unit)
		ingredients = append(ingredients, ingredient)
	}
	return
}
