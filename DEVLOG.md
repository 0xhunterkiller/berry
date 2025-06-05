# Development Logs for Berry

## May 11, 2025

- Berry is a long due personal project for me. I finally decided to get started today.
- Berry is a simple RBAC server (that's what is planned). Let's see how long I take to actually finish it.
- Finished doing a basic DB design.
- Going with an API-based model, with a kubectl-like control interface.
- Only the admin can grant or remove access for now.
- Added a Makefile and bootstrapped a simple project.
- Chose Postgres for the DB—no particular reason, might rewrite the storage layer later.
- Wrote the Docker Compose file and set up migrate (for applying schema to the Postgres DB).
- Decided to use `Gin` for the web framework.

## May 16, 2025

- Been busy with work the past few days, but found some time to continue.
- Designed the system: a decision tree-based RBAC system to reduce the time taken to resolve access questions like "Can X access Y?".
- Berry will only respond with a boolean to enhance security.
- Completed project setup.
- Created dummy API endpoints (they don't do anything yet).
- Considered examples to use for explaining and testing.

## May 18, 2025

- generated manifests and data models
- There will be a k8s style rbac, resources, verbs and users...that's all!
- wrote an organization generator (kind of a sidequest, but will prove crucial for testing system integrity in the future)
- setup complete cobra-cli (writing each function is remaining)

## May 20, 2025

- Wrote code for marshalling user, role and resource data models. 
- BerryCTL development is happening side by side, where I have finished writing a cobra cli, with create, delete, get and apply commands. There is also a `can` command that will help users find out who has access to what. It is also being written in golang.

# June 6, 2025

- Major break, I just got back at this. I have been super distracted due to some work and family stuff.
- 