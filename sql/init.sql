drop table if exists orders;
drop table if exists slots;

CREATE TABLE slots(
    id SERIAL PRIMARY KEY,
    venue VARCHAR(255) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    busy BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    slot_id INT NOT NULL, 
    customer_name VARCHAR(255) NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(slot_id) REFERENCES slots(id)
);

INSERT INTO slots(venue,start_at,end_at) VALUES
('Loft Hall','2026-03-29 10:00:00','2026-03-29 13:00:00'),
('Loft Hall','2026-03-29 14:00:00','2026-03-29 17:00:00'),
('White Box','2026-04-29 10:00:00','2026-04-29 12:00:00')