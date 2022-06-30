mongo

use solnyshko

db.createUser(
        {
            user: "root",
            pwd: "root",
            roles: [
                {
                    role: "readWrite",
                    db: "solnyshko"
                }
            ]
        }
);