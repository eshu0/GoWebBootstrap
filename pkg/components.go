package Components 

type NavLink struct {
	Active bool
	Href string
	DisplayName string
}

type BreadCrumb struct {
	Active bool
	Href string
	DisplayName string
}

type Button struct {
	DisplayName string
	Href string
	DataTarget bool
	OnClick string
	Outline string
}

type Card struct {
	Active bool
	DisplayName string
	Description string
	Mutedtext string
	Buttons []*Button
}


type Modal struct {
	Id string

	HeaderText string
	HasHeader bool

	BodyText template.HTML
	HasBody bool

	FooterText string //template.HTML
	HasFooter bool
	
	FooterButtons []*Button
}

type SiteDetails struct{
	Name string
	NavItems []*NavLink
	CSSIncludes []*CSSLink
	JSIncludes []*JSInclude
}

type CSSLink struct{
	Integrity string
	Href string
}

type JSInclude struct{
	Integrity string
	Src string
}

type PageDetails struct{
	PageTitle string
	SiteDetails *SiteDetails
	HasAddButton bool
	AddButtonHref string
	AddButtonText string
	Job *Job
	IsCardpage bool	
	IsItempage bool
	IsModalpage bool
	BreadCrumbs []*BreadCrumb
	Cards []*Card
	Modals []*Modal	
}
