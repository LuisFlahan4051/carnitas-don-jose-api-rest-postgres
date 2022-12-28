CREATE TABLE sales(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,


    entry_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    exit_date TIMESTAMP,
    table_number INT,
    packed BOOLEAN DEFAULT FALSE,


    user_id INT NOT NULL,
    branch_id INT NOT NULL,
    turn_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (turn_id) REFERENCES turns(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE sale_products(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,


    done BOOLEAN DEFAULT false,
    grammage_quantity REAL CHECK (grammage_quantity >= 0),
    kilogrammage_price REAL CHECK (kilogrammage_price >= 0),
    unit_quantity INT CHECK (unit_quantity >= 0),
    unit_price REAL CHECK (unit_price >= 0),
    discount REAL CHECK (discount >= 0),
    tax REAL CHECK (tax >= 0),

    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE sale_incomes(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    sale_id INT NOT NULL,
    income_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (income_id) REFERENCES incomes(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE sale_expenses(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    sale_id INT NOT NULL,
    expense_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (expense_id) REFERENCES expenses(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE sale_arguments(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    sale_id INT NOT NULL,
    argument_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (argument_id) REFERENCES arguments(id) ON DELETE CASCADE ON UPDATE CASCADE
);
