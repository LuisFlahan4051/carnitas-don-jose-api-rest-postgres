CREATE TABLE turns (
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP,
    active BOOLEAN DEFAULT true,

    incomes_counter REAL NOT NULL CHECK (incomes_counter >= 0) DEFAULT 0,
	netincomes_counter REAL NOT NULL DEFAULT 0,
	expenses_counter REAL NOT NULL CHECK (expenses_counter >= 0) DEFAULT 0,

    user_id INT NOT NULL,
    branch_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE turn_user_roles (
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    login_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    logout_date TIMESTAMP DEFAULT NULL,

    user_id INT NOT NULL,
    turn_id INT NOT NULL,
    role_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (turn_id) REFERENCES turns(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE turn_safebox(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    turn_id INT NOT NULL,
    safebox_id INT NOT NULL UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (safebox_id) REFERENCES safeboxes(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (turn_id) REFERENCES turns(id) ON DELETE CASCADE ON UPDATE CASCADE
);
