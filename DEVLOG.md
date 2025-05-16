# Development Logs for Berry

## May 11, 2025

Berry is a long due personal project for me. I finally decided to get started today. Berry is a simple RBAC server (thats what is planned). Let's see how long I take to actually finish it.
I finished doing a basic db design. We are gonna go with an API based model, with a kubectl like control interface. Only the admin can grant or remove access for now. I added  Makefile and bootstrapped a simple project. I chose postgres for the db - coz why not, no thought went into this, might rewrite the storage layer later. I wrote the Docker Compose file, and setup migrate (for applying schema to the psql db). Decided to use `Gin` for the web framework.

## May 16, 2025
Been busy with work the past few days, found some time to continue. I designed the system today, a decision tree based RBAC system, the decision tree reduces the time taken to resolve access questions like..."Can X access Y?". Berry will only respond in boolean, to enhance security. Enough with the yapping, lets start writing some code.
 

