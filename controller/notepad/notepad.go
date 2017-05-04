// Package notepad provides a simple CRUD application in a web page.
package notepad

import (
	"net/http"

	"github.com/nerfmiester/blue-jay/blueprint/lib/flight"
	"github.com/nerfmiester/blue-jay/blueprint/middleware/acl"
	"github.com/nerfmiester/blue-jay/blueprint/model/note"

	"github.com/nerfmiester/blue-jay/core/pagination"
	"github.com/nerfmiester/blue-jay/core/router"
)

var (
	uri = "/notepad"
)

// Load the routes.
func Load() {
	c := router.Chain(acl.DisallowAnon)
	router.Get(uri, Index, c...)
	router.Get(uri+"/create", Create, c...)
	router.Post(uri+"/create", Store, c...)
	router.Get(uri+"/view/:id", Show, c...)
	router.Get(uri+"/edit/:id", Edit, c...)
	router.Patch(uri+"/edit/:id", Update, c...)
	router.Delete(uri+"/:id", Destroy, c...)
}

// Index displays the items.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Create a pagination instance with a max of 10 results.
	p := pagination.New(r, 10)

	items, _, err := note.ByUserIDPaginate(c.DB, c.UserID, p.PerPage, p.Offset)
	if err != nil {
		c.FlashErrorGeneric(err)
		items = []note.Item{}
	}

	count, err := note.ByUserIDCount(c.DB, c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	}

	// Calculate the number of pages.
	p.CalculatePages(count)

	v := c.View.New("note/index")
	v.Vars["items"] = items
	v.Vars["pagination"] = p
	v.Render(w, r)
}

// Create displays the create form.
func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("note/create")
	c.Repopulate(v.Vars, "firstname", "middlename", "lastname", "manager", "personalemail", "businessemail", "homephone", "mobilephone", "workphone", "addressline1", "housenamenumber", "addressline2", "apartmentsuite", "town", "county", "postcode", "country", "noktitle", "nokfirstname", "noklastname", "nokrelationship", "nokhomephone", "nokworkphone", "ec1title", "ec1firstname", "ec1lastname", "ec1relationship", "ec1addressline1", "ec1housenamenumber", "ec1addressline2", "ec1apartmentsuite", "ec1town", "ec1county", "ec1postcode", "ec1country", "ec2title", "ec2firstname", "ec2lastname", "ec2relationship", "ec2addressline1", "ec2housenamenumber", "ec2addressline2", "ec2apartmentsuite", "ec2town", "ec2county", "ec2postcode", "ec2country", "ec3title", "ec3firstname", "ec3lastname", "ec3relationship", "ec3addressline1", "ec3housenamenumber", "ec3addressline2", "ec3apartmentsuite", "ec3town", "ec3county", "ec3postcode", "ec3country", "futureuse1", "futureuse2", "futureuse3", "futureuse4")
	v.Render(w, r)
}

// Store handles the create form submission.
func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("firstname") && !c.FormValid("lastname") && !c.FormAllNums(r.FormValue("futureuse4")) {
		Create(w, r)
		return
	}

	_, err := note.Create(c.DB, r.FormValue("firstname"), r.FormValue("middlename"), r.FormValue("lastname"), r.FormValue("manager"), r.FormValue("personalemail"), r.FormValue("businessemail"), r.FormValue("homephone"), r.FormValue("mobilephone"), r.FormValue("workphone"), r.FormValue("addressline1"), r.FormValue("housenamenumber"), r.FormValue("addressline2"), r.FormValue("apartmentsuite"), r.FormValue("town"), r.FormValue("county"), r.FormValue("postcode"), r.FormValue("country"), r.FormValue("noktitle"), r.FormValue("nokfirstname"), r.FormValue("noklastname"), r.FormValue("nokrelationship"), r.FormValue("nokhomephone"), r.FormValue("nokworkphone"), r.FormValue("ec1title"), r.FormValue("ec1firstname"), r.FormValue("ec1lastname"), r.FormValue("ec1relationship"), r.FormValue("ec1addressline1"), r.FormValue("ec1housenamenumber"), r.FormValue("ec1addressline2"), r.FormValue("ec1apartmentsuite"), r.FormValue("ec1town"), r.FormValue("ec1county"), r.FormValue("ec1postcode"), r.FormValue("ec1country"), r.FormValue("ec2title"), r.FormValue("ec2firstname"), r.FormValue("ec2lastname"), r.FormValue("ec2relationship"), r.FormValue("ec2addressline1"), r.FormValue("ec2housenamenumber"), r.FormValue("ec2addressline2"), r.FormValue("ec2apartmentsuite"), r.FormValue("ec2town"), r.FormValue("ec2county"), r.FormValue("ec2postcode"), r.FormValue("ec2country"), r.FormValue("ec3title"), r.FormValue("ec3firstname"), r.FormValue("ec3lastname"), r.FormValue("ec3relationship"), r.FormValue("ec3addressline1"), r.FormValue("ec3housenamenumber"), r.FormValue("ec3addressline2"), r.FormValue("ec3apartmentsuite"), r.FormValue("ec3town"), r.FormValue("ec3county"), r.FormValue("ec3postcode"), r.FormValue("ec3country"), r.FormValue("futureuse1"), r.FormValue("futureuse2"), r.FormValue("futureuse3"), r.FormValue("futureuse4"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		Create(w, r)
		return
	}

	c.FlashSuccess("Item added.")
	c.Redirect(uri)
}

// Show displays a single item.
func Show(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := note.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("note/show")
	v.Vars["item"] = item
	v.Render(w, r)
}

// Edit displays the edit form.
func Edit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := note.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("note/edit")
	c.Repopulate(v.Vars, "firstname", "middlename", "lastname", "manager", "personalemail", "businessemail", "homephone", "mobilephone", "workphone", "addressline1", "housenamenumber", "addressline2", "apartmentsuite", "town", "county", "postcode", "country", "noktitle", "nokfirstname", "noklastname", "nokrelationship", "nokhomephone", "nokworkphone", "ec1title", "ec1firstname", "ec1lastname", "ec1relationship", "ec1addressline1", "ec1housenamenumber", "ec1addressline2", "ec1apartmentsuite", "ec1town", "ec1county", "ec1postcode", "ec1country", "ec2title", "ec2firstname", "ec2lastname", "ec2relationship", "ec2addressline1", "ec2housenamenumber", "ec2addressline2", "ec2apartmentsuite", "ec2town", "ec2county", "ec2postcode", "ec2country", "ec3title", "ec3firstname", "ec3lastname", "ec3relationship", "ec3addressline1", "ec3housenamenumber", "ec3addressline2", "ec3apartmentsuite", "ec3town", "ec3county", "ec3postcode", "ec3country", "futureuse1", "futureuse2", "futureuse3", "futureuse4")
	//c.Repopulate(v.Vars, "name")
	v.Vars["item"] = item
	v.Render(w, r)
}

// Update handles the edit form submission.
func Update(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	//
	//if !c.FormValid("name") {
	//	Edit(w, r)
	//	return
	//}

	if !c.FormValid("firstname") && !c.FormValid("lastname") {
		Create(w, r)
		return
	}

	_, err := note.Update(c.DB, r.FormValue("firstname"), r.FormValue("middlename"), r.FormValue("lastname"), r.FormValue("manager"), r.FormValue("personalemail"), r.FormValue("businessemail"), r.FormValue("homephone"), r.FormValue("mobilephone"), r.FormValue("workphone"), r.FormValue("addressline1"), r.FormValue("housenamenumber"), r.FormValue("addressline2"), r.FormValue("apartmentsuite"), r.FormValue("town"), r.FormValue("county"), r.FormValue("postcode"), r.FormValue("country"), r.FormValue("noktitle"), r.FormValue("nokfirstname"), r.FormValue("noklastname"), r.FormValue("nokrelationship"), r.FormValue("nokhomephone"), r.FormValue("nokworkphone"), r.FormValue("ec1title"), r.FormValue("ec1firstname"), r.FormValue("ec1lastname"), r.FormValue("ec1relationship"), r.FormValue("ec1addressline1"), r.FormValue("ec1housenamenumber"), r.FormValue("ec1addressline2"), r.FormValue("ec1apartmentsuite"), r.FormValue("ec1town"), r.FormValue("ec1county"), r.FormValue("ec1postcode"), r.FormValue("ec1country"), r.FormValue("ec2title"), r.FormValue("ec2firstname"), r.FormValue("ec2lastname"), r.FormValue("ec2relationship"), r.FormValue("ec2addressline1"), r.FormValue("ec2housenamenumber"), r.FormValue("ec2addressline2"), r.FormValue("ec2apartmentsuite"), r.FormValue("ec2town"), r.FormValue("ec2county"), r.FormValue("ec2postcode"), r.FormValue("ec2country"), r.FormValue("ec3title"), r.FormValue("ec3firstname"), r.FormValue("ec3lastname"), r.FormValue("ec3relationship"), r.FormValue("ec3addressline1"), r.FormValue("ec3housenamenumber"), r.FormValue("ec3addressline2"), r.FormValue("ec3apartmentsuite"), r.FormValue("ec3town"), r.FormValue("ec3county"), r.FormValue("ec3postcode"), r.FormValue("ec3country"), r.FormValue("futureuse1"), r.FormValue("futureuse2"), r.FormValue("futureuse3"), r.FormValue("futureuse4"), c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		Edit(w, r)
		return
	}

	c.FlashSuccess("Item updated.")
	c.Redirect(uri)
}

// Destroy handles the delete form submission.
func Destroy(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	_, err := note.DeleteSoft(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	} else {
		c.FlashNotice("Item deleted.")
	}

	c.Redirect(uri)
}
