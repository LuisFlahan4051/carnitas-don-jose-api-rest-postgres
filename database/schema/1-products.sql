CREATE TABLE food_types (
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE food_meats(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE foods(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    description TEXT,
    photo TEXT,


    food_type_id INT NOT NULL,
    food_meat_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (food_type_id) REFERENCES food_types(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (food_meat_id) REFERENCES food_meats(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE drink_sizes(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE drink_flavors(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE drinks(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    description TEXT,
    photo TEXT,

    drink_size_id INT NOT NULL,
    drink_flavor_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (drink_size_id) REFERENCES drink_sizes(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (drink_flavor_id) REFERENCES drink_flavors(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE products(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(50) NOT NULL,
    description TEXT,
    price REAL NOT NULL CHECK (price >= 0),
    photo TEXT,

    branch_id INT NOT NULL
);

CREATE TABLE food_products(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,

    food_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (food_id) REFERENCES foods(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE drink_products(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,

    drink_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (drink_id) REFERENCES drinks(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
);
