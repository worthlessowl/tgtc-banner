package service

import (
	"errors"

	"github.com/worthlessowl/tgtc-banner/backend/database"
	"github.com/worthlessowl/tgtc-banner/backend/dictionary"
)

func GetUser(paramID int) (*dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT user_id, user_name, user_email, user_balance, user_tokopoint, user_tier, user_location, user_bannerlist
	FROM user
	WHERE user_id = $1
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	// read query result, and assign to variable(s)
	var result dictionary.User
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Balance,
		&result.Tokopoint,
		&result.Tier,
		&result.Location,
		&result.BannerList
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetUsers() ([]dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT user_id, user_name, user_email, user_balance, user_tokopoint, user_tier, user_location, user_bannerlist
	FROM user
	`
	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.User
	for rows.Next() {
		var data dictionary.User
		rows.Scan(
			&result.ID,
			&result.Name,
			&result.Email,
			&result.Balance,
			&result.Tokopoint,
			&result.Tier,
			&result.Location,
			&result.BannerList
		)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func CreateUser(data dictionary.User) (*dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	INSERT INTO user
		(
			user_name,
			user_email,
			user_balance,
			user_tokopoint,
			user_tier,
			user_location
			user_bannerlist
		)
	VALUE
		(
			$2, $3, $4, $5, $6, $7 
		)
	`
	// actual query process
	result, err := db.Exec(query,
		data.Name,
		data.Email,
		data.Balance,
		data.Tokopoint,
		data.Tier,
		data.Location,
		data.BannerList
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	return &data, nil
}

func UpdateUser(data dictionary.User) (*dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE 
		user
	SET 
		user_name = $2,
		user_email = $3,
		user_balance = $4,
		user_tokopoint = $5
		user_tier = $6
		user_location = $7
		user_bannerlist = $8
	WHERE
		product_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.Email,
		data.Balance,
		data.Tokopoint,
		data.Tier,
		data.Location,
		data.BannerList
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}

func DeleteUser(paramID int) (*dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	DELETE
	FROM user
	WHERE user_id = $1
	`
	// actual query process
	result, err := db.Exec(query, paramID)

	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}



func GetBanner(paramID int) (*dictionary.Banner, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT banner_id, banner_name, banner_url, banner_imgsrc, banner_startdate, banner_enddate, banner_mintier, banner_maxtier, banner_minbalance, banner_maxbalance, banner_mintokopoint, banner_maxtokopoint, banner_isactive
	FROM banner
	WHERE banner_id = $1
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	// read query result, and assign to variable(s)
	var result dictionary.Banner
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Url,
		&result.ImgSrc,
		&result.StartDate,
		&result.EndDate,
		&result.MinTier,
		&result.MaxTier,
		&result.MinBalance,
		&result.MaxBalance,
		&result.MinTokopoint,
		&result.MaxTokopoint,
		&result.isActive,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetBanners() ([]dictionary.User, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT banner_id, banner_name, banner_url, banner_imgsrc, banner_startdate, banner_enddate, banner_mintier, banner_maxtier, banner_minbalance, banner_maxbalance, banner_mintokopoint, banner_maxtokopoint, banner_isactive
	FROM banner
	`
	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.Banner
	for rows.Next() {
		var data dictionary.Banner
		rows.Scan(
			&result.ID,
			&result.Name,
			&result.Url,
			&result.ImgSrc,
			&result.StartDate,
			&result.EndDate,
			&result.MinTier,
			&result.MaxTier,
			&result.MinBalance,
			&result.MaxBalance,
			&result.MinTokopoint,
			&result.MaxTokopoint,
			&result.isActive,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func CreateBanner(data dictionary.Banner) (*dictionary.Banner, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	INSERT INTO banner
		(
			banner_name, banner_url, banner_imgsrc, banner_startdate, banner_enddate, banner_mintier, banner_maxtier, banner_minbalance, banner_maxbalance, banner_mintokopoint, banner_maxtokopoint, banner_isactive
		)
	VALUE
		(
			$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
	`
	// actual query process
	result, err := db.Exec(query,
		data.Name,
		data.Url,
		data.ImgSrc,
		data.StartDate,
		data.EndDate,
		data.MinTier,
		data.MaxTier,
		data.MinBalance,
		data.MaxBalance,
		data.MinTokopoint,
		data.MaxTokopoint,
		data.isActive,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	return &data, nil
}

func UpdateBanner(data dictionary.User) (*dictionary.Banner, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE 
		banner
	SET 
		banner_name = $2,
		banner_url = $3,
		banner_imgsrc = $4,
		banner_startdate = $5
		banner_enddate = $6
		banner_mintier = $7
		banner_maxtier = $8
		banner_minbalance = $9
		banner_maxbalance = $10
		banner_mintokopoint = $11
		banner_maxtokopoint = $12
		banner_isactive = $13
	WHERE
		product_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.id
		data.Name,
		data.Url,
		data.ImgSrc,
		data.StartDate,
		data.EndDate,
		data.MinTier,
		data.MaxTier,
		data.MinBalance,
		data.MaxBalance,
		data.MinTokopoint,
		data.MaxTokopoint,
		data.isActive,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}

func DeleteBanner(paramID int) (*dictionary.Banner, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	DELETE
	FROM banner
	WHERE banner_id = $1
	`
	// actual query process
	result, err := db.Exec(query, paramID)

	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}