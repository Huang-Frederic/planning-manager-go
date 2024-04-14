CREATE TABLE Room (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    capacity INT NOT NULL
);

CREATE TABLE Reservation (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES Room(id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);
