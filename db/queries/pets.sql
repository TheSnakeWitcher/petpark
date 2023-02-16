-- name: GetPet :one
select * from pets where id = $1 ;

-- name: ListPets :many
select * from pets order by name;

-- name: AddPet :execresult
insert into pets ( id , name, picked ,location,data) 
values( $1,$2,now(),$3,'{}') ;

-- name: DelPet :exec
delete from pets where id = $1 ;

-- name: UpdateState :exec
update pets 
set status = 'adopted'
where id = ?;
