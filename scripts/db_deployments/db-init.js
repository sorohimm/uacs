db.createUser(
        {
            user: "dev",
            pwd: "dev",
            roles: [
                {
                    role: "readWrite",
                    db: "solnyshko"
                }
            ]
        }
);

//db.createCollection("competitions");