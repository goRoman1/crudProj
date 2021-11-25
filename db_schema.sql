DROP TABLE IF EXISTS Scooters CASCADE;
CREATE TABLE Scooters(
    id serial PRIMARY KEY,
    model varchar NOT NULL,
    brand varchar NOT NULL,
    capacity int,
    max_weight int,
    max_distance int,
    serial int
);


DROP TABLE IF EXISTS Test CASCADE;
CREATE TABLE Test(
     id serial PRIMARY KEY,
     model varchar NOT NULL,
     brand varchar NOT NULL
);

