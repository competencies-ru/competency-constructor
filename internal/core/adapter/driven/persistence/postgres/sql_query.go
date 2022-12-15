package postgres

const (
	getAllUgsn = `select *from ugsn`

	createUgsn = `insert into ugsn(code, title) values ($1, $2)`

	existUgsn = `select exists(select 1 from ugsn where code=$1)`

	findUgns = `select distinct u.code, u.title,
                s.code, s.title, s.code_ugsn,
                p.id, p.title, p.specialty_code
				from ugsn u
         		left join specialty s on u.code = s.code_ugsn
         		left join program p on s.code = p.specialty_code
				where u.code=$1`
)
