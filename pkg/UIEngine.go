package UIEngine 

import (
	"fmt"
	"os"
	"path/filepath"	
	
	sli "github.com/eshu0/simplelogger/interfaces"
	comps "github.com/eshu0/GoWebBootstrap/pkg/components"
)

type UIEngine struct {
	SiteDetails *comps.SiteDetails		`json:"-"`
	Log sli.ISimpleLogger				`json:"-"`
}

func NewUIEngine(sitename string, log sli.ISimpleLogger) *UIEngine {
	sd := comps.SiteDetails{ Name: sitename }
	navitems := []*comps.NavLink{}
	sd.NavItems = navitems
	uie := UIEngine{ SiteDetails:&sd, Log: log}
	return &uie
}

func (uie *UIEngine) LoadCSSFiles(myDir string) {

	var loadcss = func(path string, info os.FileInfo, err error) error {
		// first thing to do, check error. and decide what to do about it
		if err != nil {
			uie.Log.LogDebugf("LoadCSSFiles","error %v at a path %q\n", err, path)
			return err
		}
		
		uie.Log.LogDebugf("LoadCSSFiles","path: %s", path)

		// find out if it's a dir or file, if file, print info
		if info.IsDir() {
			uie.Log.LogDebug("LoadCSSFiles","is dir.")

		} else {

			uie.Log.LogDebugf("LoadCSSFiles","  dir: %v", filepath.Dir(path))
			uie.Log.LogDebugf("LoadCSSFiles","  file name %v ", info.Name())
			uie.Log.LogDebugf("LoadCSSFiles","  extenion: %v ", filepath.Ext(path))
			if filepath.Ext(path) == ".css" {
				cssIncludes := uie.SiteDetails.CSSIncludes
				cssIncludes = append(cssIncludes, &comps.CSSLink{ Href: "/"+path  } )
				uie.SiteDetails.CSSIncludes = cssIncludes
				uie.Log.LogDebugf("LoadCSSFiles","Added CSS include: %v", path)

			}
		}

		return nil
	}

	err := filepath.Walk(myDir, loadcss)

	if err != nil {
		uie.Log.LogDebugf("LoadCSSFiles","error walking the path %q: %v", myDir, err)
	}
}

func (uie *UIEngine) LoadJavascriptFiles(myDir string) {

	var loadjs = func(path string, info os.FileInfo, err error) error {
		// first thing to do, check error. and decide what to do about it
		if err != nil {
			uie.Log.LogErrorf("LoadJavascriptFiles","error %v at a path %q", err, path)
			return err
		}

		uie.Log.LogDebugf("LoadJavascriptFiles","path: %v\n", path)

		// find out if it's a dir or file, if file, print info
		if info.IsDir() {
			fmt.Printf("is dir.")
		} else {

			uie.Log.LogDebugf("LoadJavascriptFiles","  dir: %v", filepath.Dir(path))
			uie.Log.LogDebugf("LoadJavascriptFiles","  file name %v", info.Name())
			uie.Log.LogDebugf("LoadJavascriptFiles","  extenion: %v ", filepath.Ext(path))
			if filepath.Ext(path) == ".js" {
				jsIncludes := uie.SiteDetails.JSIncludes
				jsIncludes = append(jsIncludes, &comps.JSInclude{ Src: "/"+path  } )
				uie.SiteDetails.JSIncludes = jsIncludes
				uie.Log.LogDebugf("LoadJavascriptFiles","Added JS include: %v", path)
			}
		}

		return nil
	}

	err := filepath.Walk(myDir, loadjs)

	if err != nil {
		uie.Log.LogDebugf("LoadJavascriptFiles","error walking the path %q: %v", myDir, err)
	}
}

func (uie *UIEngine) NewModal(Id string) *comps.Modal {

	md := 	comps.Modal{ Id: Id } 
	
	md.HeaderText = ""
	md.HasHeader = true

	md.BodyText =""
	md.HasBody = true

	md.FooterText =  "" //template.HTML(" ")
	md.HasFooter = true
	md.FooterButtons =  []*comps.Button{}

	return &md
}

func (uie *UIEngine) NewPageDetails() *comps.PageDetails {

	pd := 	comps.PageDetails{ SiteDetails: uie.SiteDetails } 

	pd.IsCardpage = false
	pd.IsItempage = false
	pd.IsModalpage = false

	pd.BreadCrumbs = []*comps.BreadCrumb{}
	pd.Cards = []*comps.Card{}	
	pd.Modals = []*comps.Modal{}

	return &pd
}

func (uie *UIEngine) NewCardPage(title string, modals[]*comps.Modal, crumbs []*comps.BreadCrumb, cards []*comps.Card) *comps.PageDetails {

	pd := uie.NewPageDetails()

	pd.IsCardpage = len(cards) > 0
	pd.IsItempage = false
	pd.IsModalpage = len(modals) > 0
	pd.PageTitle = title

	pd.BreadCrumbs = crumbs
	pd.Cards = cards	
	pd.Modals = modals

	return pd
}

func (uie *UIEngine) NewItemPage(title string, modals[]*comps.Modal, crumbs []*comps.BreadCrumb) *comps.PageDetails {

	pd := uie.NewPageDetails()

	pd.IsCardpage = false
	pd.IsItempage = true
	pd.IsModalpage = len(modals) > 0
	pd.PageTitle = title

	pd.BreadCrumbs = crumbs
	pd.Cards = []*comps.Card{}	
	pd.Modals = modals

	return pd
}



