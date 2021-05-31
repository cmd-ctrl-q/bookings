ALTER TABLE room_restrictions
  ADD CONSTRAINT room_id_fk
    FOREIGN KEY (room_id)
      REFERENCES rooms(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE;

ALTER TABLE room_restrictions
  ADD CONSTRAINT restriction_id_fk
    FOREIGN KEY (restriction_id)
      REFERENCES restrictions(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE;
