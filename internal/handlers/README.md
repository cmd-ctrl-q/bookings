

```go
// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// get the remote ip address of the person visiting the site and store in session.
	remoteIP := r.RemoteAddr                              // get ip (version 4 or 6)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // add ip to session
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
```



// extras
```go
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	// pull data out of session
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "can't parse form")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// sd := r.Form.Get("start_date")
	// ed := r.Form.Get("end_date")

	// // 2020-01-01 -- 01/02 03:04:05PM '06 -0700
	// layout := "2006-01-02"
	// startDate, err := time.Parse(layout, sd)
	// if err != nil {
	// 	m.App.Session.Put(r.Context(), "error", "can't parse start date")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }
	// endDate, err := time.Parse(layout, ed)
	// if err != nil {
	// 	m.App.Session.Put(r.Context(), "error", "can't parse end date")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// // get room id
	// roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	// if err != nil {
	// 	m.App.Session.Put(r.Context(), "error", "can't get room id from session")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// update reservation
	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Phone = r.Form.Get("phone")
	reservation.Email = r.Form.Get("email")
	// reservation.StartDate = startDate
	// reservation.EndDate = endDate
	// reservation.RoomID = roomID

	form := forms.New(r.PostForm)

	// Validation Rules
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.MinLength("last_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		// render form back to user if there were invalid fields
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// write reservation data into db
	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert reservation into database!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert room restriction!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// add reservation back into session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// do redirect because we dont want users to accidentally submit the form twice
	// anytime you receive a post request, you should direct users to an http redirect.
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}
```