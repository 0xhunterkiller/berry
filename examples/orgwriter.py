import yaml
import random
import string
import os

resc = 10 # (input("Enter the number of resources: ")) 
userc = 10 # (input("Enter the number of users: "))
rolec = 10 # (input("Enter the number of roles: "))
verbminc = 2 # (input("Enter the min number of verbs/resource: "))
verbmaxc = 10 # (input("Enter the max number of verbs/resource: "))
apiversion = "v1" # (input("Enter the API version: "))

orgname = f"org_{''.join(random.choices(string.ascii_letters,k=10))}"
os.makedirs(orgname, exist_ok=False)
os.makedirs(f"{orgname}/resources", exist_ok=False)
os.makedirs(f"{orgname}/roles", exist_ok=False)
os.makedirs(f"{orgname}/users", exist_ok=False)

# Generate a list of resources
org = {}

org["resources"] = []
org["users"] = []
org["roles"] = []

for i in range(resc):
    resname = "resource"+str(i+1)
    # Generate a list of for that verbs
    verbs = []
    for j in range(random.randint(verbminc, verbmaxc+1)):
        verbs.append(f"r{i+1}verb{j+1}")
    resource = {
        "name": resname,
        "description": f"resource description {i+1}",
        "verbs": verbs
    }
    org["resources"].append(resource) 

# Generate a list of roles
roles = []
for i in range(rolec):
    role = {}
    role["name"] = f"role{i+1}"
    role["description"] = f"role description {i+1}"
    assoc_res = random.choices(org["resources"], k=random.randint(1, resc))
    role["resources"] = []
    for res in assoc_res:
        name = res["name"]
        verbs = random.choices(res["verbs"], k=random.randint(1, len(res["verbs"])))
        role["resources"].append({
            "name": name,
            "verbs": verbs
        })

    roles.append(role)
org["roles"] = roles

# Generate a list of users
users = []
for i in range(userc):
    user = {}
    user["name"] = f"user{i+1}"
    user["description"] = f"user description {i+1}"
    user["roles"] = [x["name"] for x in random.choices(roles, k=random.randint(1, len(roles)))]
    if random.random() < 0.3:
        user["roles"].append("admin")
    users.append(user)

org["users"] = users

# Atleast 1 admin
if not any(["admin" in x["roles"] for x in org["users"]]):
    org["users"][0]["roles"].append("admin")

for resource in org["resources"]:
    # write resource to yaml file
    tmp = {}
    tmp["apiVersion"] = apiversion
    tmp["kind"] = "Resource"
    tmp["spec"] = resource
    with open(f"{orgname}/resources/{resource['name']}.yaml", "w") as f:
        yaml.dump(tmp, f, default_flow_style=False)

for role in org["roles"]:
    # write role to yaml file
    tmp = {}
    tmp["apiVersion"] = apiversion
    tmp["kind"] = "Role"
    tmp["spec"] = role
    with open(f"{orgname}/roles/{role['name']}.yaml", "w") as f:
        yaml.dump(tmp, f, default_flow_style=False)

for user in org["users"]:
    # write user to yaml file
    tmp = {}
    tmp["apiVersion"] = apiversion
    tmp["kind"] = "User"
    tmp["spec"] = user
    with open(f"{orgname}/users/{user['name']}.yaml", "w") as f:
        yaml.dump(tmp, f, default_flow_style=False)