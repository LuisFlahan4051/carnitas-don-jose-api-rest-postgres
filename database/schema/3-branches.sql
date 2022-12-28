CREATE TABLE branches(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    address TEXT,

    PRIMARY KEY (id)
);

CREATE TABLE branch_safeboxes (
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    name VARCHAR(40) NOT NULL,
    content INT DEFAULT 0,

    branch_id INT NOT NULL,
    safebox_id INT NOT NULL UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (safebox_id) REFERENCES safeboxes(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE branch_products_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,


    branch_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE branch_supplies_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0) DEFAULT 0,
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,

    branch_id INT NOT NULL,
    supply_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (supply_id) REFERENCES supplies(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE branch_articles_stock(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    unit_quantity INT NOT NULL CHECK (unit_quantity >= 0),
    grammage_quantity INT NOT NULL CHECK (grammage_quantity >= 0) DEFAULT 0,
    in_use BOOLEAN,

    branch_id INT NOT NULL,
    article_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE ON UPDATE CASCADE
);
