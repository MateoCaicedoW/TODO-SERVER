SELECT EXISTS(
    SELECT 1
    FROM users WHERE users.email = ?
) 
