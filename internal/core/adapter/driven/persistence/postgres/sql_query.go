package postgres

// level query.
const (
	getLevel    = `SELECT id, title FROM LEVEL WHERE id = $1`
	getLevels   = `SELECT id, title FROM LEVEL`
	updateLevel = `UPDATE LEVEL SET title = $2 WHERE id = $1`
	insertLevel = `INSERT INTO LEVEL (id, title) VALUES ($1, $2)`
	deleteLevel = `DELETE FROM LEVEL WHERE id = $1`
)

// ugsn query.
const (
	getUgsn           = `SELECT id, code, title, level_id FROM UGSN WHERE id = $1`
	getUgsnsByLevelID = `SELECT id, code, title, level_id FROM UGSN WHERE level_id = $1`
	updateUgsn        = `UPDATE UGSN SET code = $1, title = $2 WHERE id = $3`
	insertUgsn        = `INSERT INTO UGSN (id, code, title, level_id) VALUES ($1, $2, $3, $4)`
	deleteUgsn        = `DELETE FROM UGSN WHERE id = $1`
)

// specialty query.
const (
	getSpecialty         = `SELECT id, code, title, ugsn_id FROM SPECIALTY WHERE id = $1`
	deleteSpecialty      = `DELETE FROM SPECIALTY WHERE id = $1`
	getSpecialtyByUgsnID = `SELECT id, code, title, ugsn_id FROM SPECIALTY WHERE ugsn_id = $1`
	updateSpecialty      = `UPDATE SPECIALTY SET code = $1, title = $2 WHERE id = $4`
	insertSpecialty      = `INSERT INTO SPECIALTY (id, code, title, ugsn_id) VALUES ($1, $2, $3, $4)`
)

// program query.
const (
	insertProgram           = `INSERT INTO PROGRAM (id, code, title, specialties_id) VALUES ($1, $2, $3, $4)`
	deleteProgram           = `DELETE FROM PROGRAM WHERE id = $1`
	getProgram              = `SELECT id, code, title, specialties_id FROM PROGRAM WHERE id = $1`
	updateProgram           = `UPDATE PROGRAM SET code = $1, title = $2 WHERE id = $4`
	getProgramBySpecialtyID = `SELECT id, code, title, specialties_id FROM PROGRAM WHERE specialties_id = $1`
)

// subject query.
const (
	insertSubject       = `INSERT INTO SUBJECT (id, name, sname) VALUES ($1, $2, $3)`
	deleteSubject       = `DELETE FROM SUBJECT WHERE id = $1`
	updateSubject       = `UPDATE SUBJECT SET name = $2, sname = $3 WHERE id = $1`
	filterSubjectByName = `SELECT id, name, sname FROM SUBJECT WHERE name LIKE '%' || $1 || '%'`
)

// competency query.
const (
	insertCompetency = `INSERT INTO COMPETENCY (id, code, title, category, competency_type, 
                          level_id, ugsn_id, specialty_id, program_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	deleteCompetency = `DELETE FROM COMPETENCY WHERE id = $1`

	updateCompetency = `UPDATE COMPETENCY SET code = $2, title = $3, category = $4 WHERE id = $1`

	getCompetency = `SELECT ID, CODE, TITLE, CATEGORY, COMPETENCY_TYPE, 
       LEVEL_ID, UGSN_ID, SPECIALTY_ID, PROGRAM_ID FROM COMPETENCY WHERE ID = $1`

	filterCompetency = `SELECT * FROM COMPETENCY WHERE
            (UGSN_ID = $1 OR $1 IS NULL) OR
            (SPECIALTY_ID = $2 OR $2 IS NULL) OR
            (PROGRAM_ID = $3 OR $3 IS NULL) OR
            (LEVEL_ID = $4 OR $4 IS NULL)`
)

// indicator question.
const (
	insertIndicator = `INSERT INTO INDICATOR(id, title, code, subject_id, competency_id) VALUES($1, $2, $3, $4, $5)`

	getIndicator = ` SELECT i.id, i.title, i.code, i.competency_id FROM INDICATOR i WHERE i.id = $1`

	findAllIndicators = `SELECT i.id, i.title, i.code, i.competency_id, s.id, s.name, s.sname  FROM INDICATOR i
        LEFT JOIN SUBJECT s ON s.id = i.subject_id
        WHERE i.competency_id = $1`

	updateIndicator = `UPDATE INDICATOR SET title = $1, code = $2, subject_id = $3 WHERE id = $4`
)
