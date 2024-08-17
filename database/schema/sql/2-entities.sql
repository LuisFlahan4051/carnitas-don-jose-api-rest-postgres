CREATE TABLE supplies(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    description TEXT,
    photo TEXT,

    PRIMARY KEY (id)
);

CREATE TABLE articles(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    description TEXT,
    photo TEXT,

    PRIMARY KEY (id)
);

CREATE TABLE safeboxes (
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    cents10 INT CHECK (cents10 >= 0),
    cents50 INT CHECK (cents50 >= 0),
    coins1 INT CHECK (coins1 >= 0),
    coins2 INT CHECK (coins2 >= 0),
    coins5 INT CHECK (coins5 >= 0),
    coins10 INT CHECK (coins10 >= 0),
    coins20 INT CHECK (coins20 >= 0),

    bills20 INT CHECK (bills20 >= 0),
    bills50 INT CHECK (bills50 >= 0),
    bills100 INT CHECK (bills100 >= 0),
    bills200 INT CHECK (bills200 >= 0),
    bills500 INT CHECK (bills500 >= 0),
    bills1000 INT CHECK (bills1000 >= 0),

    PRIMARY KEY (id)
);

CREATE TABLE incomes(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    reason TEXT NOT NULL,
    income REAL NOT NULL CHECK (income >= 0),


    PRIMARY KEY (id)
);


CREATE TABLE expenses(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    reason TEXT NOT NULL,
    expense REAL NOT NULL CHECK (expense >= 0),

    PRIMARY KEY (id)
);

CREATE TABLE arguments(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,


    complaint BOOLEAN NOT NULL DEFAULT false,
    score INT NOT NULL CHECK (score >= 0 AND score <= 5),
    argument TEXT NOT NULL,

    PRIMARY KEY (id)
);

