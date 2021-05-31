ALTER TABLE reservations
  ADD CONSTRAINT room_id_fk
    FOREIGN KEY (room_id)
      REFERENCES rooms(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE;
