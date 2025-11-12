package authentification

/*func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("method not allowed")
		services.ResponseFunc(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}

	var ul usersModels.UserLogin
	if err := services.DecodeRequest(w, r, &ul); err != nil {
		log.Printf("failed to decode request,/auth/login\nerror:%v", err)
		services.ResponseFunc(w, http.StatusBadRequest, "failed to decode request", nil)
		return
	}

	if err := encryption.CheckPasswordHash(ul); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Printf("invalid password,/auth/login error:%v", err)
			services.ResponseFunc(w, http.StatusUnauthorized, "invalid password or email", nil)
			return
		}
		log.Printf("failed to check match passwords,/auth/login error:%v", err)
		services.ResponseFunc(w, http.StatusUnauthorized, "failed to check match passwords", nil)
		return
	}

	uid, err := database.GetUserID(ul)
	if err != nil {
		log.Printf("failed to get UserID,error:%v", err)
		services.ResponseFunc(w, http.StatusInternalServerError, "failed to get UserID", nil)
		return
	}

	token, err := CreateJwt(uid)
	if err != nil {
		log.Printf("failed to create JWT /login, error:%v", err)
		services.ResponseFunc(w, http.StatusInternalServerError, "failed to create JWT", nil)
		return
	}
	services.ResponseFunc(w, http.StatusOK, "user login successfully", map[string]string{"token": token})
}*/
