-- to test after the migration that drop default of updated_at column
-- create a user
INSERT INTO users (email, password, last_name, first_name, is_active, role)
VALUES (
  'asdasdasd@gmail.com',
  '$2a$10$I9ZdFZ1OMx.LO3dnmv65DO344FPoaUj8LXXv01jzmIgIKFAqF5uia',
  'Dummy',
  'Account',
  true,
  'user'::user_role
  );

