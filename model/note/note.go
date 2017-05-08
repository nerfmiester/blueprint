// Package note provides access to the note table in the MySQL database.
package note

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	// table is the table name.
	table = "note"
)

// Item defines the model.
type Item struct {
	ID                 uint32         `db:"id"`
	FirstName          string         `db:"firstname"`
	MiddleName         string         `db:"middlename"`
	LastName           string         `db:"lastname"`
	Manager            string         `db:"manager"`
	PersonalEmail      string         `db:"personalemail"`
	BusinessEmail      string         `db:"businessemail"`
	HomePhone          string         `db:"homephone"`
	MobilePhone        string         `db:"mobilephone"`
	WorkPhone          string         `db:"workphone"`
	AddressLine1       string         `db:"addressline1"`
	HouseNameNumber    string         `db:"housenamenumber"`
	AddressLine2       string         `db:"addressline2"`
	AppartmentSuite    string         `db:"apartmentsuite"`
	Town               string         `db:"town"`
	County             string         `db:"county"`
	PostCode           string         `db:"postcode"`
	Country            string         `db:"country"`
	NokTitle           string         `db:"noktitle"`
	NokFirstName       string         `db:"nokfirstname"`
	NokLastName        string         `db:"noklastname"`
	NokRelationship    string         `db:"nokrelationship"`
	NokAddressLine1    string         `db:"nokaddressline1"`
	NokHouseNameNumber string         `db:"nokhousenamenumber"`
	NokAddressLine2    string         `db:"nokaddressline2"`
	NokApartmentSuite  string         `db:"nokapartmentsuite"`
	NokTown            string         `db:"noktown"`
	NokCounty          string         `db:"nokcounty"`
	NokPostCode        string         `db:"nokpostcode"`
	NokCountry         string         `db:"nokcountry"`
	NokHomePhone       string         `db:"nokhomephone"`
	NokWorkPhone       string         `db:"nokworkphone"`
	Ec1Title           string         `db:"ec1title"`
	Ec1FirstName       string         `db:"ec1firstname"`
	Ec1LastName        string         `db:"ec1lastname"`
	Ec1Relationship    string         `db:"ec1relationship"`
	Ec1AddressLine1    string         `db:"ec1addressline1"`
	Ec1HouseNameNumber string         `db:"ec1housenamenumber"`
	Ec1AddressLine2    string         `db:"ec1addressline2"`
	Ec1ApartmentSuite  string         `db:"ec1apartmentsuite"`
	Ec1Town            string         `db:"ec1town"`
	Ec1County          string         `db:"ec1county"`
	Ec1PostCode        string         `db:"ec1postcode"`
	Ec1Country         string         `db:"ec1country"`
	Ec1HomePhone       string         `db:"ec1homephone"`
	Ec1WorkPhone       string         `db:"ec1workphone"`
	Ec2Title           string         `db:"ec2title"`
	Ec2FirstName       string         `db:"ec2firstname"`
	Ec2LastName        string         `db:"ec2lastname"`
	Ec2Relationship    string         `db:"ec2relationship"`
	Ec2AddressLine1    string         `db:"ec2addressline1"`
	Ec2HouseNameNumber string         `db:"ec2housenamenumber"`
	Ec2AddressLine2    string         `db:"ec2addressline2"`
	Ec2ApartmentSuite  string         `db:"ec2apartmentsuite"`
	Ec2Town            string         `db:"ec2town"`
	Ec2County          string         `db:"ec2county"`
	Ec2PostCode        string         `db:"ec2postcode"`
	Ec2Country         string         `db:"ec2country"`
	Ec2HomePhone       string         `db:"ec2homephone"`
	Ec2WorkPhone       string         `db:"ec2workphone"`
	Ec3Title           string         `db:"ec3title"`
	Ec3FirstName       string         `db:"ec3firstname"`
	Ec3LastName        string         `db:"ec3lastname"`
	Ec3Relationship    string         `db:"ec3relationship"`
	Ec3AddressLine1    string         `db:"ec3addressline1"`
	Ec3HouseNameNumber string         `db:"ec3housenamenumber"`
	Ec3AddressLine2    string         `db:"ec3addressline2"`
	Ec3ApartmentSuite  string         `db:"ec3apartmentsuite"`
	Ec3Town            string         `db:"ec3town"`
	Ec3County          string         `db:"ec3county"`
	Ec3PostCode        string         `db:"ec3postcode"`
	Ec3Country         string         `db:"ec3country"`
	Ec3HomePhone       string         `db:"ec3homephone"`
	Ec3WorkPhone       string         `db:"ec3workphone"`
	FutureUse1         string         `db:"futureuse1"`
	FutureUse2         string         `db:"futureuse2"`
	FutureUse3         string         `db:"futureuse3"`
	FutureUse4         string         `db:"futureuse4"`
	UserID             uint32         `db:"user_id"`
	CreatedAt          mysql.NullTime `db:"created_at"`
	UpdatedAt          mysql.NullTime `db:"updated_at"`
	DeletedAt          mysql.NullTime `db:"deleted_at"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByID gets an item by ID.
func ByID(db Connection, ID string, userID string) (Item, bool, error) {
	result := Item{}
	err := db.Get(&result, fmt.Sprintf(`
		SELECT id, firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1,nokhousenamenumber, nokaddressline2,nokapartmentsuite,noktown,nokcounty,	nokpostcode,	nokcountry,nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country, ec1homephone,ec1workphone,ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone,ec2workphone,ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec3homephone,ec3workphone,futureuse1, futureuse2, futureuse3, futureuse4, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ID, userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserID gets all items for a user.
func ByUserID(db Connection, userID string) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
		SELECT id, firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1,nokhousenamenumber, nokaddressline2,nokapartmentsuite,noktown,nokcounty,	nokpostcode,	nokcountry,nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country,ec1homephone,ec1workphone, ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone,ec2workphone,ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec3homephone,ec3workphone,futureuse1, futureuse2, futureuse3, futureuse4, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDPaginate gets items for a user based on page and max variables.
func ByUserIDPaginate(db Connection, userID string, max int, page int) ([]Item, bool, error) {
	var result []Item
	err := db.Select(&result, fmt.Sprintf(`
		SELECT id, firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1,nokhousenamenumber, nokaddressline2,nokapartmentsuite,noktown,nokcounty,	nokpostcode,	nokcountry,nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country, ec1homephone,ec1workphone,ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone,ec2workphone,ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec3homephone,ec3workphone, futureuse1, futureuse2, futureuse3, futureuse4, user_id, created_at, updated_at, deleted_at
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		LIMIT %v OFFSET %v
		`, table, max, page),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDCount counts the number of items for a user.
func ByUserIDCount(db Connection, userID string) (int, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
		SELECT count(*)
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err
}

// Create adds an item.
func Create(db Connection, firstname string, middlename string, lastname string, manager string, personalemail string, businessemail string, homephone string, mobilephone string, workphone string, addressline1 string, housenamenumber string, addressline2 string, apartmentsuite string, town string, county string, postcode string, country string, noktitle string, nokfirstname string, noklastname string, nokrelationship string, nokaddressline1 string, nokhousenamenumber string, nokaddressline2 string, nokapartmentsuite string, noktown string, nokcounty string, nokpostcode string, nokcountry string, nokhomephone string, nokworkphone string, ec1title string, ec1firstname string, ec1lastname string, ec1relationship string, ec1addressline1 string, ec1housenamenumber string, ec1addressline2 string, ec1apartmentsuite string, ec1town string, ec1county string, ec1postcode string, ec1country string, ec1homephone string, ec1workphone string, ec2title string, ec2firstname string, ec2lastname string, ec2relationship string, ec2addressline1 string, ec2housenamenumber string, ec2addressline2 string, ec2apartmentsuite string, ec2town string, ec2county string, ec2postcode string, ec2country string, ec2homephone string, ec2workphone string, ec3title string, ec3firstname string, ec3lastname string, ec3relationship string, ec3addressline1 string, ec3housenamenumber string, ec3addressline2 string, ec3apartmentsuite string, ec3town string, ec3county string, ec3postcode string, ec3country string, ec3homephone string, ec3workphone string, futureuse1 string, futureuse2 string, futureuse3 string, futureuse4 string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1,nokhousenamenumber, nokaddressline2,nokapartmentsuite,noktown,nokcounty,	nokpostcode,	nokcountry,nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country, ec1homephone,ec1workphone,ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone,ec2workphone,ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec3homephone,ec3workphone,futureuse1, futureuse2, futureuse3, futureuse4, user_id)
		VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
		`, table),
		firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1, nokhousenamenumber, nokaddressline2, nokapartmentsuite, noktown, nokcounty, nokpostcode, nokcountry, nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country, ec1homephone, ec1workphone, ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone, ec2workphone, ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec3homephone, ec3workphone, futureuse1, futureuse2, futureuse3, futureuse4, userID)
	return result, err
}

// Update makes changes to an existing item.
func Update(db Connection, firstname string, middlename string, lastname string, manager string, personalemail string, businessemail string, homephone string, mobilephone string, workphone string, addressline1 string, housenamenumber string, addressline2 string, apartmentsuite string, town string, county string, postcode string, country string, noktitle string, nokfirstname string, noklastname string, nokrelationship string, nokaddressline1 string, nokhousenamenumber string, nokaddressline2 string, nokapartmentsuite string, noktown string, nokcounty string, nokpostcode string, nokcountry string, nokhomephone string, nokworkphone string, ec1title string, ec1firstname string, ec1lastname string, ec1relationship string, ec1addressline1 string, ec1housenamenumber string, ec1addressline2 string, ec1apartmentsuite string, ec1town string, ec1county string, ec1postcode string, ec1country string, ec1homephone string, ec1workphone string, ec2title string, ec2firstname string, ec2lastname string, ec2relationship string, ec2addressline1 string, ec2housenamenumber string, ec2addressline2 string, ec2apartmentsuite string, ec2town string, ec2county string, ec2postcode string, ec2country string, ec2homephone string, ec2workphone string, ec3title string, ec3firstname string, ec3lastname string, ec3relationship string, ec3addressline1 string, ec3housenamenumber string, ec3addressline2 string, ec3apartmentsuite string, ec3town string, ec3county string, ec3postcode string, ec3country string, ec3homephone string, ec3workphone string, futureuse1 string, futureuse2 string, futureuse3 string, futureuse4 string, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET firstname = ?,
		middlename = ?,
		lastname  = ?,
		manager = ?,
		personalemail = ?,
		businessemail = ?,
		homephone = ?,
		mobilephone = ?,
		workphone = ?,
		addressline1 = ?,
		housenamenumber = ?,
		addressline2 = ?,
		apartmentsuite = ?,
		town = ?,
		county = ?,
		postcode = ?,
		country = ?,
		noktitle = ?,
		nokfirstname = ?,
		noklastname = ?,
		nokrelationship = ?,
		nokaddressline1  = ?,
		nokhousenamenumber  = ?,
		nokaddressline2 = ?,
		nokapartmentsuite = ?,
		noktown = ?,
		nokcounty = ?,
		nokpostcode = ?,
		nokcountry = ?,
		nokhomephone = ?,
		nokworkphone = ?,
		ec1title = ?,
		ec1firstname = ?,
		ec1lastname = ?,
		ec1relationship = ?,
		ec1addressline1 = ?,
		ec1housenamenumber = ?,
		ec1addressline2 = ?,
		ec1apartmentsuite = ?,
		ec1town = ?,
		ec1county = ?,
		ec1postcode = ?,
		ec1country = ?,
		ec1homephone = ?,
		ec1workphone = ?,
		ec2title = ?,
		ec2firstname = ?,
		ec2lastname = ?,
		ec2relationship = ?,
		ec2addressline1 = ?,
		ec2housenamenumber = ?,
		ec2addressline2 = ?,
		ec2apartmentsuite = ?,
		ec2town = ?,
		ec2county = ?,
		ec2postcode = ?,
		ec2country = ?,
		ec2homephone  = ?,
		ec2workphone  = ?,
		ec3title = ?,
		ec3firstname = ?,
		ec3lastname = ?,
		ec3relationship = ?,
		ec3addressline1 = ?,
		ec3housenamenumber = ?,
		ec3addressline2 = ?,
		ec3apartmentsuite = ?,
		ec3town = ?,
		ec3county = ?,
		ec3postcode = ?,
		ec3country = ?,
		ec3homephone  = ?,
		ec3workphone  = ?,
		futureuse1 = ?,
		futureuse2 = ?,
		futureuse3 = ?,
		futureuse4 =?
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		firstname, middlename, lastname, manager, personalemail, businessemail, homephone, mobilephone, workphone, addressline1, housenamenumber, addressline2, apartmentsuite, town, county, postcode, country, noktitle, nokfirstname, noklastname, nokrelationship, nokaddressline1, nokhousenamenumber, nokaddressline2, nokapartmentsuite, noktown, nokcounty, nokpostcode, nokcountry, nokhomephone, nokworkphone, ec1title, ec1firstname, ec1lastname, ec1relationship, ec1addressline1, ec1housenamenumber, ec1addressline2, ec1apartmentsuite, ec1town, ec1county, ec1postcode, ec1country, ec1homephone, ec1workphone, ec2title, ec2firstname, ec2lastname, ec2relationship, ec2addressline1, ec2housenamenumber, ec2addressline2, ec2apartmentsuite, ec2town, ec2county, ec2postcode, ec2country, ec2homephone, ec2workphone,ec3title, ec3firstname, ec3lastname, ec3relationship, ec3addressline1, ec3housenamenumber, ec3addressline2, ec3apartmentsuite, ec3town, ec3county, ec3postcode, ec3country, ec1homephone, ec1workphone, futureuse1, futureuse2, futureuse3, futureuse4, ID, userID)
	return result, err
}

// DeleteHard removes an item.
func DeleteHard(db Connection, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		DELETE FROM %v
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		`, table),
		ID, userID)
	return result, err
}

// DeleteSoft marks an item as removed.
func DeleteSoft(db Connection, ID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET deleted_at = NOW()
		WHERE id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ID, userID)
	return result, err
}
