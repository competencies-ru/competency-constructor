db.createUser(
    {
        user: "admin",
        pwd: "qwerty",
        roles: [
            {
                role: "readWrite",
                db: "competency-constructor"
            }
        ]
    }
)