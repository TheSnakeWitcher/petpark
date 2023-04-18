-- name: GetPet :one
select * from pets where id = $1 ;

-- name: ListPets :many
select * from pets order by picked ;

-- name: AddPet :execresult
insert into pets (id , picked , address , contact, details)
values(DEFAULT, $1 , $2 , $3 , $4) ;

-- name: DelPet :exec
delete from pets where id = $1 ;
