package models

import (
	"example.com/beers/db"
)

type Beer struct {
	ID            int64
	Name          string  `json:"name"`
	Price         string  `json:"price"`
	Image         string  `json:"image"`
	RatingAverage float64 `json:"rating_average"`
	RatingReviews int     `json:"rating_reviews"`
}

func (b *Beer) Save() error {
	query := `INSERT INTO beers(name, price, image, rating_average, rating_reviews)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(b.Name, b.Price, b.Image, b.RatingAverage, b.RatingReviews)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id

	return err
}

func GetAllBeers() ([]Beer, error) {
	query := "SELECT * FROM beers"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []Beer

	for rows.Next() {
		var beer Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Price, &beer.Image, &beer.RatingAverage, &beer.RatingReviews)
		if err != nil {
			return nil, err
		}

		beers = append(beers, beer)
	}

	return beers, nil
}

func GetBeerById(id int64) (*Beer, error) {
	query := "SELECT * FROM beers WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var beer Beer

	err := row.Scan(&beer.ID, &beer.Name, &beer.Price, &beer.Image, &beer.RatingAverage, &beer.RatingReviews)
	if err != nil {
		return nil, err
	}

	return &beer, nil
}

func (b *Beer) Delete() error {
	query := "DELETE FROM beers WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.ID)

	return err
}

func (b *Beer) Update() error {
	query := `
	UPDATE beers
	SET name = ?, price = ?, image = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Name, b.Price, b.Image, b.ID)

	return err
}
