ALTER TABLE room_restrictions
  ADD CONSTRAINT reservation_id_fk
    FOREIGN KEY (reservation_id)
      REFERENCES reservations(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE;

-- quick lookup by email
CREATE INDEX reservations_email_idx ON reservations (email);

-- or quick lookup by last name
CREATE INDEX reservations_last_name_idx ON reservations (last_name);