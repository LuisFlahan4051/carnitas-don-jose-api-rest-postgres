CREATE TABLE safebox_actions(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    withdrawal BOOLEAN DEFAULT false,

    user_id INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE action_safebox(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    safebox_actions_id INT NOT NULL UNIQUE,
    safebox_id INT NOT NULL UNIQUE,
    PRIMARY KEY (id),
    FOREIGN KEY (safebox_actions_id) REFERENCES safebox_actions(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (safebox_id) REFERENCES safeboxes(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE notifications(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    type TEXT NOT NULL, /* INFO/WARNING/ERROR/DONE/URGENT */
    solved BOOLEAN DEFAULT false,
    description TEXT NOT NULL,

    branch_id INT,
    user_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (branch_id) REFERENCES branches(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE notification_images(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    image TEXT NOT NULL,

    notification_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (notification_id) REFERENCES notifications(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE server_logs(
    id SERIAL UNIQUE,
    created_at TIMESTAMP DEFAULT  CURRENT_TIMESTAMP,

    transaction TEXT NOT NULL,
    user_id INT NOT NULL,
    branch_id INT,
    root BOOLEAN DEFAULT false,

    PRIMARY KEY (id)
);
