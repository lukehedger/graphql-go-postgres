package api

import (
	"fmt"
)

func (r *Resolver) Hello(args struct{ ID string }) (string, error) {
	var name string
	rows, err := r.DB.Query("SELECT name FROM people WHERE id = $1 LIMIT $2", args.ID, 1)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello, %v", name), nil
}
