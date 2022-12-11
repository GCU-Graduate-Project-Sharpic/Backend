CREATE TABLE user_account (
    username VARCHAR(30) PRIMARY KEY,
    password VARCHAR(200) NOT NULL,
    email VARCHAR(40) NOT NULL
);

CREATE TABLE album (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) REFERENCES user_account,
    title VARCHAR(30) NOT NULL,
    CONSTRAINT no_duplicate UNIQUE (username, title)
);

CREATE TABLE image (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) REFERENCES user_account,
    image_name VARCHAR(100) NOT NULL,
    image_file bytea NOT NULL,
    size int NOT NULL,
    added_date timestamp DEFAULT Now(),

    /*
    Specify the image upgrade method by integer
    0: normal (non sr)
    1: image SR
    2: image recovery
    ... 
    */
    up int NOT NULL
);

CREATE TABLE processed_image (
    id int PRIMARY KEY REFERENCES image(id),
    username VARCHAR(30) REFERENCES user_account,
    image_name VARCHAR(100) NOT NULL,
    image_file bytea NOT NULL,
    size int NOT NULL,
    added_date timestamp DEFAULT Now(),

    /*
    Specify the image upgrade method by integer
    0: normal (non sr)
    1: image SR
    2: image recovery
    ... 
    */
    up int NOT NULL
);

CREATE TABLE album_image (
    album_id int NOT NULL REFERENCES album(id),
    image_id int NOT NULL REFERENCES image(id),
    PRIMARY KEY(album_id, image_id)
);
