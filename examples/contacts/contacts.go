package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import "github.com/dragonfax/glitter/glt"

var _ StatelessWidget = &ContactCategory{}

type ContactCategory struct {
	Icon     IconData
	Children []Widget
}

func (cc *ContactCategory) Build(context glt.BuildContext) (Widget, error) {
	var themeData ThemeData = Theme.of(context)
	return &glt.Container{
		Padding: EdgeInsets{Vertical: 16.0},
		Decoration: &BoxDecoration{
			Border: &Border{Bottom: &BorderSide{Color: themeData.dividerColor}},
		},
		Child: &DefaultTextStyle{
			Style: Theme.of(context).textTheme.subhead,
			Child: &SafeArea{
				Top:    false,
				Bottom: false,
				Child: &Row{
					CrossAxisAlignment: CrossAxisAlignment.start,
					Children: []Widget{
						&Container{
							Padding: EdgeInsets{Vertical: 24.0},
							Width:   72.0,
							Child:   &Icon{Icon: icon, Color: themeData.primaryColor},
						},
						&Expanded{Child: &Column{Children: children}},
					},
				},
			},
		},
	}
}

var _ StatelessWidget = &ContactItem

type ContactItem struct {
	Icon      IconData
	lines     []string
	Tooltip   string
	onPressed VoidCallback
}

func (ci *ContactItem) Build(context glt.BuildContext) (Widget, error) {
	// ThemeData themeData = Theme.of(context);
	var columnChildren []Widget = make([]Widget, 0, len(lines))
	for _, line := range lines[0:-1] {
		columnChildren = append(columnChildren, &glt.Text{Text: line})
	}
	columnChildren = append(columnChildren, &Text{Text: lines[-1], Style: themeData.textTheme.caption})

	rowChildren := []Widget{
		&glt.Expanded{
			Child: &Column{
				CrossAxisAlignment: CrossAxisAlignment.start,
				Children:           columnChildren,
			},
		},
	}

	if icon != nil {
		rowChildren = append(rowChildren, &SizedBox{
			Width: 72.0,
			Child: &IconButton{
				Icon:      &Icon{Icon: icon},
				Color:     themeData.primaryColor,
				OnPressed: onPressed,
			},
		})
	}
	return &MergeSemantics{
		Child: &Padding{
			Padding: &EdgeInsets{Vertical: 16.0},
			Child: &Row{
				MainAxisAlignment: MainAxisAlignment.spaceBetween,
				Children:          rowChildren,
			},
		},
	}
}

var _ StatefulWidget = &ContactsDemo{}

type ContactsDemo struct {
}

func (cd *ContactsDemo) CreateState() State {
	return &ContactsDemoState{}
}

type AppBarBehavior uint8

const (
	AppBarBehaviorNormal AppBarBehavior = iota
	AppBarBehaviorPinned
	AppBarBehaviorFloating
	AppBarBehaviorSnapping
)

var _ State = &ContactsDemoState{}

type ContactsDemoState struct {
	scaffoldKey  GlobalKey
	appBarHeight float32
}

// Use Init() in lue of a constructor
func (cds *ContactsDemoState) Init() {
	cds.scaffoldKey = NewGlobalKey()
	cds.appBarHeight = 256.0
}

func (cds *ContactsDemoState) Build(context glt.BuildContext) (Widget, error) {
	var appBarBehavior AppBarBehavior = AppBarBehaviorPinned

	return &Theme{
		Data: &ThemeData{
			Brightness:    Brightness.light,
			PrimarySwatch: Colors.indigo,
			Platform:      Theme.of(context).platform,
		},
		Child: &Scaffold{
			Key: cds.scaffoldKey,
			Body: &CustomScrollView{
				Slivers: []Widget{
					&SliverAppBar{
						ExpandedHeight: cds.appBarHeight,
						Pinned:         appBarBehavior == AppBarBehaviorPinned,
						Floating:       appBarBehavior == AppBarBehaviorFloating || appBarBehavior == AppBarBehaviorSnapping,
						Snap:           appBarBehavior == AppBarBehaviorSnapping,
						Actions: []Widget{
							&IconButton{
								Icon:    &Icon{Icon: IconsCreate},
								Tooltip: "Edit",
								OnPressed: func() {
									cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
										Content: &Text{"Editing isn't supported in this screen."},
									})
								},
							},
							&PopupMenuButton{
								OnSelected: func(value AppBarBehavior) {
									context.setState(func() {
										appBarBehavior = value
									})
								},
								ItemBuilder: func(context glt.BuildContext) {
									return []PopupMenuItem{
										&PopupMenuItem{
											Value: AppBarBehaviorNormal,
											Child: &Text{"App bar scrolls away"},
										},
										&PopupMenuItem{
											Value: AppBarBehaviorPinned,
											Child: &Text{"App bar stays put"},
										},
										&PopupMenuItem{
											Value: AppBarBehaviorFloating,
											Child: &Text{"App bar floats"},
										},
										&PopupMenuItem{
											Value: AppBarBehaviorSnapping,
											Child: &Text{"App bar snaps"},
										},
									}
								},
							},
						},
						FlexibleSpace: &FlexibleSpaceBar{
							Title: &Text{"Ali Connors"},
							Background: &Stack{
								Fit: StackFitExpand,
								Children: []Widget{
									NewImageFromAsset(
										Asset{
											File:    "people/ali_landscape.png",
											Package: "flutter_gallery_assets",
										},
										&Image{
											Fit:    BoxFitCover,
											Height: cds.appBarHeight,
										},
									),
									// This gradient ensures that the toolbar icons are distinct
									// against the background image.
									&DecoratedBox{
										Decoration: &BoxDecoration{
											Gradient: &LinearGradient{
												Begin:  Alignment(0.0, -1.0),
												End:    Alignment(0.0, -0.4),
												Colors: []Color{Color{0x60000000}, Color{0x00000000}},
											},
										},
									},
								},
							},
						},
					},
					&SliverList{
						Delegate: &SliverChildListDelegate{[]Widget{
							&AnnotatedRegion{
								Value: SystemUiOverlayStyleDark,
								Child: &ContactCategory{
									Icon: IconsCall,
									Children: []Widget{
										&ContactItem{
											Icon:    IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												scaffoldKey.currentState.showSnackBar(&SnackBar{
													Content: &Text{"Pretend that this opened your SMS application."},
												})
											},
											Lines: []string{
												"(650) 555-1234",
												"Mobile",
											},
										},
										&ContactItem{
											Icon:    IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												scaffoldKey.currentState.showSnackBar(&SnackBar{
													Content: &Text{"A messaging app appears."},
												})
											},
											Lines: []string{
												"(323) 555-6789",
												"Work",
											},
										},
										&ContactItem{
											Icon:    IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
													Content: &Text{"Imagine if you will, a messaging application."},
												})
											},
											Lines: []string{
												"(650) 555-6789",
												"Home",
											},
										},
									},
								},
							},
							&ContactCategory{
								Icon: IconsContactMail,
								Children: []Widget{
									&ContactItem{
										Icon:    IconsEmail,
										Tooltip: "Send personal e-mail",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
												Content: &Text{"Here, your e-mail application would open."},
											})
										},
										Lines: []string{
											"ali_connors@example.com",
											"Personal",
										},
									},
									&ContactItem{
										Icon:    IconsEmail,
										Tooltip: "Send work e-mail",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
												Content: &Text{"Summon your favorite e-mail application here."},
											})
										},
										Lines: []string{
											"aliconnors@example.com",
											"Work",
										},
									},
								},
							},
							&ContactCategory{
								Icon: IconsLocationOn,
								Children: []Widget{
									&ContactItem{
										Icon:    IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
												Content: &Text{"This would show a map of San Francisco."},
											})
										},
										Lines: []string{
											"2000 Main Street",
											"San Francisco, CA",
											"Home",
										},
									},
									&ContactItem{
										Icon:    IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
												Content: &Text{"This would show a map of Mountain View."},
											})
										},
										Lines: []string{
											"1600 Amphitheater Parkway",
											"Mountain View, CA",
											"Work",
										},
									},
									&ContactItem{
										Icon:    IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&SnackBar{
												Content: &Text{"This would also show a map, if this was not a demo."},
											})
										},
										Lines: []string{
											"126 Severyns Ave",
											"Mountain View, CA",
											"Jet Travel",
										},
									},
								},
							},
							&ContactCategory{
								Icon: IconsToday,
								Children: []Widget{
									&ContactItem{
										Lines: []string{
											"Birthday",
											"January 9th, 1989",
										},
									},
									&ContactItem{
										Lines: []string{
											"Wedding anniversary",
											"June 21st, 2014",
										},
									},
									&ContactItem{
										Lines: []string{
											"First day in office",
											"January 20th, 2015",
										},
									},
									&ContactItem{
										Lines: []string{
											"Last day in office",
											"August 9th, 2018",
										},
									},
								},
							},
						}},
					},
				},
			},
		},
	}
}
