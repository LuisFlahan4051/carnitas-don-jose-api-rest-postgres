CREATE TABLE inventory_types(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    inventory_type TEXT NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE inventories(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    acepted BOOLEAN,

    inventory_type_id INT NOT NULL, /* INIT/FINAL/LOSSES */
    turn_id INT NOT NULL,
    branch_id INT NOT NULL, /*  << Opcional, se puede referir indirectamente por lógica pero lo dejo por facilidad */
    PRIMARY KEY (id),
    FOREIGN KEY (inventory_type_id) REFERENCES inventory_types(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (turn_id) REFERENCES turns(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE inventory_products_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,

    inventory_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE inventory_supplies_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,

    inventory_id INT NOT NULL,
    supply_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (supply_id) REFERENCES supplies(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE inventory_articles_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0),
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,

    inventory_id INT NOT NULL,
    article_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE inventory_safebox(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    inventory_id INT NOT NULL UNIQUE,
    safebox_id INT NOT NULL UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (safebox_id) REFERENCES safeboxes(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (inventory_id) REFERENCES inventories(id) ON DELETE CASCADE ON UPDATE CASCADE
);
