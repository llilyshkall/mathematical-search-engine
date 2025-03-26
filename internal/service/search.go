package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// Добавляем структуру для хранения данных формулы
type FormulaResult struct {
	Formula string `json:"formula"`
	Mask    string `json:"mask"`
}

type Request struct {
	Input string `json:"input"`
}

func (s *Service) SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response, err := s.search(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

//func (s *Service) search(r Request) ([]FormulaResult, error) {
//	expressions, err := s.db.GetAll()
//	if err != nil {
//		return nil, err
//	}
//	ret := make([]FormulaResult, 0)
//	expr := math_expression_v2.ParseLaTeX(r.Input)
//	for _, row := range expressions {
//		e := math_expression_v2.ParseLaTeX(row.Latex)
//		if e.Compare(expr) != math_expression_v2.Different {
//			m, _ := e.MaskLaTeX(expr)
//			ret = append(ret, FormulaResult{Formula: row.Latex, Mask: m})
//		}
//	}
//	log.Println("response:", ret)
//	return ret, nil
//}

func (s *Service) search(r Request) ([]FormulaResult, error) {
	expressions, err := s.db.GetAll()
	if err != nil {
		return nil, err
	}
	ret := make([]FormulaResult, 0)
	expr := math_expression.ParseLaTeX(r.Input)
	for _, row := range expressions {
		e := math_expression.ParseLaTeX(row.Latex)
		if e.HasPrefix(expr) {
			ret = append(ret, FormulaResult{Formula: row.Latex, Mask: e.Mask(expr)})
		}
	}
	log.Println("response:", ret)
	return ret, nil
}
