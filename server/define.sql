-- Creating the USER table, setting S for standard user permissions and A for admin permissions

CREATE TABLE 'user' (
  'user' CHAR(22) PRIMARY KEY NOT NULL,
  username TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  type TEXT CHECK( type IN ('normal', 'admin') ) NOT NULL,
  amb INT NOT NULL DEFAULT -1,
  depot INT NOT NULL DEFAULT -1,
  platoon INT NOT NULL DEFAULT -1,
  section INT NOT NULL DEFAULT -1,
  man INT NOT NULL DEFAULT -1,
  name TEXT NOT NULL
);

CREATE TABLE 'task' (
  'task' CHAR(22) PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  assigned_to CHAR(22) NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  verified BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY(assigned_to) REFERENCES 'user'('user')
);

