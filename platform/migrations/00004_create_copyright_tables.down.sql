CREATE TABLE ip_copyrights (
    id UUID FOREIGN KEY REFERENCES intellectual_property(id),
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    name VARCHAR (255) NOT NULL,
    holder UUID FOREIGN KEY REFERENCES users(id),
    owner JSONB NOT NULL,
    approved_at TIMESTAMP,
    approved_by UUID FOREIGN KEY REFERENCES users(id),
    file VARCHAR (255) NOT NULL
);