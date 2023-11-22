BEGIN;
ALTER TABLE appeals
    RENAME COLUMN x TO latitude;
ALTER TABLE appeals
    RENAME COLUMN y TO longitude;
COMMIT;