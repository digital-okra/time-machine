-- Creating the USER table, setting S for standard user permissions and A for admin permissions

CREATE TABLE 'user' (
  'user' CHAR(22) PRIMARY KEY NOT NULL,
  username TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  type CHAR(1) CHECK( type IN ('S', 'A') ) NOT NULL,
  platoon INT DEFAULT NULL,
  section INT DEFAULT NULL,
  man INT DEFAULT NULL,
  name TEXT NOT NULL
)

CREATE TABLE 'task' (
  'task' CHAR(22) PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  assigned_to CHAR(22) NOT NULL,
  FOREIGN KEY(assigned_to) REFERENCES 'user'('user'),
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  verified BOOLEAN NOT NULL DEFAULT FALSE
)

