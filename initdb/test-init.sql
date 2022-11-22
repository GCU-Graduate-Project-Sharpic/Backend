CREATE TABLE user_list (
    username VARCHAR(30) PRIMARY KEY,
    password VARCHAR(200) NOT NULL,
    email VARCHAR(40) NOT NULL
);

CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    album VARCHAR(30) DEFAULT 'default',
    image_name VARCHAR(100) NOT NULL,
    image_file bytea NOT NULL,
    size int NOT NULL,

    /*
    Specify the image upgrade method by integer
    0: normal (non sr)
    1: image SR
    2: image filter
    ... 
    */
    up int NOT NULL
);
