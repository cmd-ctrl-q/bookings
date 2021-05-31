-- drop foreign keys
ALTER TABLE room_restrictions
  DROP CONSTRAINT reservation_id_fk;

-- drop indices
DROP INDEX reservations_email_idx;
DROP INDEX reservations_last_name_idx;