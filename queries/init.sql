CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY,
    login TEXT,
    password TEXT, 
    name TEXT, 
    age INTEGER);

CREATE TABLE IF NOT EXISTS phone(id INTEGER PRIMARY KEY, phone TEXT, description TEXT, is_fax TINYINT, FOREIGN KEY(user_id) REFERENCES user(id));