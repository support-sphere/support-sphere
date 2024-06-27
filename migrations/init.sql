-- Define ENUM types
CREATE TYPE user_role AS ENUM ('admin', 'agent', 'customer');
CREATE TYPE ticket_status AS ENUM ('open', 'in_progress', 'closed', 'on_hold');
CREATE TYPE ticket_priority AS ENUM ('low', 'medium', 'high', 'urgent');
CREATE TYPE field_type AS ENUM ('text', 'number', 'date', 'boolean', 'select');
CREATE TYPE entity_type AS ENUM ('ticket', 'user', 'department', 'sub_ticket');

-- Create tables
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tickets (
    ticket_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status ticket_status NOT NULL,
    priority ticket_priority NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(user_id),
    assigned_to UUID REFERENCES users(user_id)
);

CREATE TABLE departments (
    department_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ticket_interactions (
    interaction_id UUID PRIMARY KEY,
    ticket_id UUID REFERENCES tickets(ticket_id),
    user_id UUID REFERENCES users(user_id),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ticket_tags (
    ticket_id UUID REFERENCES tickets(ticket_id),
    tag VARCHAR(255) NOT NULL,
    PRIMARY KEY (ticket_id, tag)
);

CREATE TABLE attachments (
    attachment_id UUID PRIMARY KEY,
    ticket_id UUID REFERENCES tickets(ticket_id),
    file_path VARCHAR(255) NOT NULL,
    uploaded_by UUID REFERENCES users(user_id),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ticket_history (
    history_id UUID PRIMARY KEY,
    ticket_id UUID REFERENCES tickets(ticket_id),
    status ticket_status NOT NULL,
    changed_by UUID REFERENCES users(user_id),
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sub_tickets (
    sub_ticket_id UUID PRIMARY KEY,
    parent_ticket_id UUID REFERENCES tickets(ticket_id),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status ticket_status NOT NULL,
    priority ticket_priority NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID REFERENCES users(user_id),
    assigned_to UUID REFERENCES users(user_id)
);

CREATE TABLE user_groups (
    group_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE user_group_members (
    group_id UUID REFERENCES user_groups(group_id),
    user_id UUID REFERENCES users(user_id),
    PRIMARY KEY (group_id, user_id)
);

CREATE TABLE custom_fields (
    field_id UUID PRIMARY KEY,
    entity_type entity_type NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    field_type field_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE custom_field_values (
    value_id UUID PRIMARY KEY,
    field_id UUID REFERENCES custom_fields(field_id),
    entity_id UUID NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT entity_custom_field UNIQUE (field_id, entity_id)
);
