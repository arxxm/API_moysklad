package commands

import "net/url"

type Meta struct {
	Href         url.URL
	MetadataHref url.URL
	Type         string
	MediaType    string
	UuidHref     url.URL
	DownloadHref url.URL
}

type Employee struct {
	meta    Meta
	context Meta
	rows    []rowsEmployee
}

type rowsEmployee struct {
	id           string
	accountId    string
	updated      string
	name         string
	externalCode string
	archived     string
	uid          string
	email        string
	phone        string
	firstName    string
	middleName   string
	lastName     string
	fullName     string
	shortFio     string
	inn          string
	position     string
}
