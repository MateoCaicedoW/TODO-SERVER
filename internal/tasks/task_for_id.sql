SELECT tasks.id as task_id,
tasks.title,
tasks.description, 
tasks.status,
users.id as user_id,
CONCAT(users.first_name, ' ', users.last_name) as user_fullname,
users.email as user_email
FROM tasks
JOIN users ON users.id = tasks.user_id
WHERE tasks.id = ?