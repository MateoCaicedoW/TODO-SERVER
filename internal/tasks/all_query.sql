SELECT 
tasks.id "id",
tasks.title "title",
tasks.description "description",
tasks.must "must",
tasks.status "status",
tasks.completeby "completed",
users.rol "user_role",
users.email "user_email",
users.id "user_id",
CONCAT(users.first_name , ' ', users.last_name) "user_fullname"

FROM tasks
JOIN users ON users.id = tasks.user_id