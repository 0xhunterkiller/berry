# Berry

# RBAC System

A Role-Based Access Control (RBAC) system that manages user roles and permissions to secure resources and streamline access control within an application.

## Features:

- [x] **User Authentication**: Basic user login system with secure authentication.
- [ ] **Roles & Permissions**: Define roles (e.g., Admin, User, Manager) and assign permissions (e.g., create, read, update, delete).
- [ ] **Dynamic Role Assignment**: Assign roles to users dynamically, with the ability to change roles as needed.
- [ ] **Resource Enrollment**: Register resources (e.g., API endpoints, database tables) within the system and manage their access controls.

- [ ] **Role Hierarchy**: Support for role inheritance (e.g., Admin > Manager > User) to simplify permission management.
- [ ] **Permission Validation**: Check user permissions dynamically for each resource or action based on assigned roles.
- [ ] **User Groups**: Group users with common roles or permissions for simplified management.
- [ ] **Access Control Lists (ACLs)**: Define access control lists to specify what resources can be accessed by which users or roles.
- [ ] **Audit Logs**: Track role assignments, permission changes, and resource access for security and accountability.
- [ ] **Token Based Access Control**: Use JWT or other token mechanisms to validate user permissions in the authentication flow.
- [ ] **Resource Based Policies**: Create policies that govern access to specific resources or actions, beyond just roles.

# ER Diagram

```mermaid
erDiagram

    users ||--o{ users_roles : has
    roles ||--o{ users_roles : assigned_to
    roles ||--o{ roles_permissions : has
    permissions ||--o{ roles_permissions : assigned_to
    permissions ||--o{ permissions_resource_actions : has
    resource_actions ||--o{ permissions_resource_actions : has
    resources ||--o{ resource_actions : has
    actions ||--o{ resource_actions : has

    audit {
        UUID id PK
        UUID resource_action_id
        UUID user_id
        JSONB changes
        string message
        timestamp createdat
    }

    users {
        UUID id PK
        string name
        string email
        string password
        boolean is_active
        timestamp createdat
        timestamp updatedat
    }

    roles {
        UUID id PK
        string name
        string description
        timestamp createdat
    }

    users_roles {
        UUID id PK
        UUID role_id FK
        UUID user_id FK
        timestamp createdat
    }

    permissions {
        UUID id PK
        string name
        string description
        timestamp createdat
    }

    roles_permissions {
        UUID id PK
        UUID role_id FK
        UUID permission_id FK
        timestamp createdat
    }

    resources {
        UUID id PK
        string name
        string description
        timestamp createdat
    }

    actions {
        UUID id PK
        string name
        string description
        timestamp createdat
    }

    resource_actions {
        UUID id PK
        UUID resource_id FK
        UUID action_id FK
        timestamp createdat
    }

    permissions_resource_actions {
        UUID id PK
        UUID permission_id FK
        UUID resource_action_id FK
        timestamp createdat
    }
```