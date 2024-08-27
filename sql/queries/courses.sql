-- name: CreateCourse :exec
insert into courses (id, name, description, category_id, price)
values ($1, $2, $3, $4, $5);

-- name: ListCourses :many
select c.*, ca.name as category_name 
from courses c join categories ca on c.category_id = ca.id;