// Authorize user
func authorizeUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if user exists
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		log.Printf("Error checking if user exists: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check if password is correct
	if user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User authorized successfully")
}

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	// Initialize router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", registerUser).Methods("POST")
	r.HandleFunc("/authorize", authorizeUser).Methods("POST")

	// Start server
	log.Println("Starting server on port 8080")
	log.Fatal(http



