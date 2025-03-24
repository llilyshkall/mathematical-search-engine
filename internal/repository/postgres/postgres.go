package postgres

import (
	"database/sql"
)

type expressionRepository struct {
	db *sql.DB
}

type Expression struct {
	Latex string
}

type ExpressionRepository interface {
	Insert(expr Expression) error
	GetAll() ([]Expression, error)
	Delete(latex string) error
}

func NewExpressionRepository(db *sql.DB) ExpressionRepository {
	return &expressionRepository{db: db}
}

func (r *expressionRepository) Insert(expr Expression) error {
	query := `INSERT INTO expressions (latex) VALUES ($1)`
	_, err := r.db.Exec(query, expr.Latex)
	return err
}

func (r *expressionRepository) GetAll() ([]Expression, error) {
	query := `SELECT latex FROM expressions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expressions []Expression
	for rows.Next() {
		var expr Expression
		if err := rows.Scan(&expr.Latex); err != nil {
			return nil, err
		}
		expressions = append(expressions, expr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expressions, nil
}

func (r *expressionRepository) Delete(latex string) error {
	query := `DELETE FROM expressions WHERE latex = $1`
	_, err := r.db.Exec(query, latex)
	return err
}
