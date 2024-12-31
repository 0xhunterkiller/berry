CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    userid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    hpassword TEXT NOT NULL,
    isactive BOOLEAN DEFAULT TRUE,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL,
    updatedat TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Roles table
CREATE TABLE roles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Permissions table
CREATE TABLE permissions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Resources table
CREATE TABLE resources (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Actions table
CREATE TABLE actions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL
);

-- UserRole table
CREATE TABLE users_roles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    role_id UUID NOT NULL,
    user_id UUID NOT NULL,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(userid) ON DELETE CASCADE,
    CONSTRAINT fk_role_id FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_role UNIQUE (user_id, role_id)
);

-- RolePermission table
CREATE TABLE roles_permissions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_permission_id FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_id FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT unique_role_permission UNIQUE (role_id, permission_id)
);

-- ResourceAction table
CREATE TABLE resource_actions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    resource_id UUID NOT NULL,
    action_id UUID NOT NULL,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_resource_id FOREIGN KEY(resource_id) REFERENCES resources(id) ON DELETE CASCADE,
    CONSTRAINT fk_action_id FOREIGN KEY(action_id) REFERENCES actions(id) ON DELETE CASCADE,
    CONSTRAINT unique_resource_action UNIQUE (resource_id, action_id)
);

-- PermissionResourceAction table
CREATE TABLE permissions_resource_actions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    resource_action_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    createdat TIMESTAMP DEFAULT NOW() NOT NULL,
    CONSTRAINT fk_resource_action_id FOREIGN KEY(resource_action_id) REFERENCES resource_actions(id) ON DELETE CASCADE,
    CONSTRAINT fk_permission_id FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    CONSTRAINT unique_permission_resource_action UNIQUE (permission_id, resource_action_id)
);
