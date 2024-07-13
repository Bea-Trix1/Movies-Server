package dto

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationRequestParms struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

func FromValuePaginationRequestParams(r *http.Request) (*PaginationRequestParms, error) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(r.FormValue("itemsPerPage"))

	paginationRequestParms := PaginationRequestParms{
		Search:       r.FormValue("search"),
		Descending:   strings.Split(r.FormValue("descending"), ","),
		Sort:         strings.Split(r.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}

	return &paginationRequestParms, nil
}
