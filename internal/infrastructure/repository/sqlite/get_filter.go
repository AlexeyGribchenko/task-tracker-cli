package sqlite

const queryGetFiltered = `
	SELECT *
	FROM tasks
	WHERE status = $1
`
