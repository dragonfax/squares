package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import "github.com/dragonfax/glitter/glt"

var _ glt.StatelessWidget = &ContactCategory{}

type ContactCategory struct {
	Icon     glt.IconData
	Children []glt.Widget
}

func (cc *ContactCategory) Build(context glt.BuildContext) (glt.Widget, error) {
	// var themeData glt.ThemeData = glt.Theme.of(context)
	return &glt.Container{
		Padding: glt.EdgeInsets{Vertical: 16.0},
		Decoration: &glt.BoxDecoration{
			Border: &glt.Border{Bottom: &glt.BorderSide{Color: themeData.dividerColor}},
		},
		Child: &glt.DefaultTextStyle{
			Style: glt.Theme.of(context).textTheme.subhead,
			Child: &glt.SafeArea{
				Top:    false,
				Bottom: false,
				Child: &glt.Row{
					CrossAxisAlignment: glt.CrossAxisAlignmentStart,
					Children: []glt.Widget{
						&glt.Container{
							Padding: glt.EdgeInsets{Vertical: 24.0},
							Width:   72.0,
							Child:   &glt.Icon{Icon: icon, Color: themeData.primaryColor},
						},
						&glt.Expanded{Child: &glt.Column{Children: children}},
					},
				},
			},
		},
	}, nil
}

type VoidCallback func()

var _ glt.StatelessWidget = &ContactItem{}

type ContactItem struct {
	Icon      glt.IconData
	lines     []string
	Tooltip   string
	onPressed VoidCallback
}

func (ci *ContactItem) Build(context glt.BuildContext) (glt.Widget, error) {
	// ThemeData themeData = Theme.of(context);
	var columnChildren []glt.Widget = make([]glt.Widget, 0, len(lines))
	for _, line := range lines[0:-1] {
		columnChildren = append(columnChildren, &glt.Text{Text: line})
	}
	columnChildren = append(columnChildren, &glt.Text{Text: lines[-1], Style: themeData.textTheme.caption})

	rowChildren := []glt.Widget{
		&glt.Expanded{
			Child: &glt.Column{
				CrossAxisAlignment: glt.CrossAxisAlignmentStart,
				Children:           columnChildren,
			},
		},
	}

	if icon != nil {
		rowChildren = append(rowChildren, &glt.SizedBox{
			Width: 72.0,
			Child: &glt.IconButton{
				Icon:      &glt.Icon{Icon: icon},
				Color:     themeData.primaryColor,
				OnPressed: onPressed,
			},
		})
	}
	return &glt.MergeSemantics{
		Child: &glt.Padding{
			Padding: &glt.EdgeInsets{Vertical: 16.0},
			Child: &glt.Row{
				MainAxisAlignment: glt.MainAxisAlignmentSpaceBetween,
				Children:          rowChildren,
			},
		},
	}
}

var _ glt.StatefulWidget = &ContactsDemo{}

type ContactsDemo struct {
}

func (cd *ContactsDemo) CreateState() glt.State {
	return &ContactsDemoState{}
}

type AppBarBehavior uint8

const (
	AppBarBehaviorNormal AppBarBehavior = iota
	AppBarBehaviorPinned
	AppBarBehaviorFloating
	AppBarBehaviorSnapping
)

var _ glt.State = &ContactsDemoState{}

type ContactsDemoState struct {
	scaffoldKey  glt.GlobalKey
	appBarHeight float32
}

// Use Init() in lue of a constructor
func (cds *ContactsDemoState) Init() {
	cds.scaffoldKey = glt.NewGlobalKey()
	cds.appBarHeight = 256.0
}

func (cds *ContactsDemoState) Build(context glt.BuildContext) (glt.Widget, error) {
	var appBarBehavior AppBarBehavior = AppBarBehaviorPinned

	return &glt.Theme{
		Data: &glt.ThemeData{
			Brightness:    glt.BrightnessLight,
			PrimarySwatch: glt.ColorsIndigo,
			Platform:      glt.Theme.of(context).platform,
		},
		Child: &glt.Scaffold{
			Key: cds.scaffoldKey,
			Body: &glt.CustomScrollView{
				Slivers: []glt.Widget{
					&glt.SliverAppBar{
						ExpandedHeight: cds.appBarHeight,
						Pinned:         appBarBehavior == AppBarBehaviorPinned,
						Floating:       appBarBehavior == AppBarBehaviorFloating || appBarBehavior == AppBarBehaviorSnapping,
						Snap:           appBarBehavior == AppBarBehaviorSnapping,
						Actions: []glt.Widget{
							&glt.IconButton{
								Icon:    &glt.Icon{Icon: glt.IconsCreate},
								Tooltip: "Edit",
								OnPressed: func() {
									cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
										Content: &glt.Text{"Editing isn't supported in this screen."},
									})
								},
							},
							&glt.PopupMenuButton{
								OnSelected: func(value AppBarBehavior) {
									context.setState(func() {
										appBarBehavior = value
									})
								},
								ItemBuilder: func(context glt.BuildContext) {
									return []glt.PopupMenuItem{
										&glt.PopupMenuItem{
											Value: AppBarBehaviorNormal,
											Child: &glt.Text{"App bar scrolls away"},
										},
										&glt.PopupMenuItem{
											Value: AppBarBehaviorPinned,
											Child: &glt.Text{"App bar stays put"},
										},
										&glt.PopupMenuItem{
											Value: AppBarBehaviorFloating,
											Child: &glt.Text{"App bar floats"},
										},
										&glt.PopupMenuItem{
											Value: AppBarBehaviorSnapping,
											Child: &glt.Text{"App bar snaps"},
										},
									}
								},
							},
						},
						FlexibleSpace: &glt.FlexibleSpaceBar{
							Title: &glt.Text{"Ali Connors"},
							Background: &glt.Stack{
								Fit: glt.StackFitExpand,
								Children: []glt.Widget{
									glt.NewImageFromAsset(
										glt.Asset{
											File:    "people/ali_landscape.png",
											Package: "flutter_gallery_assets",
										},
										&glt.Image{
											Fit:    glt.BoxFitCover,
											Height: cds.appBarHeight,
										},
									),
									// This gradient ensures that the toolbar icons are distinct
									// against the background image.
									&glt.DecoratedBox{
										Decoration: &glt.BoxDecoration{
											Gradient: &glt.LinearGradient{
												Begin:  glt.Alignment(0.0, -1.0),
												End:    glt.Alignment(0.0, -0.4),
												Colors: []glt.Color{glt.Color{0x60000000}, glt.Color{0x00000000}},
											},
										},
									},
								},
							},
						},
					},
					&glt.SliverList{
						Delegate: &glt.SliverChildListDelegate{[]glt.Widget{
							&glt.AnnotatedRegion{
								Value: glt.SystemUiOverlayStyleDark,
								Child: &glt.ContactCategory{
									Icon: glt.IconsCall,
									Children: []glt.Widget{
										&glt.ContactItem{
											Icon:    glt.IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
													Content: &glt.Text{"Pretend that this opened your SMS application."},
												})
											},
											Lines: []string{
												"(650) 555-1234",
												"Mobile",
											},
										},
										&ContactItem{
											Icon:    glt.IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
													Content: &glt.Text{"A messaging app appears."},
												})
											},
											Lines: []string{
												"(323) 555-6789",
												"Work",
											},
										},
										&ContactItem{
											Icon:    glt.IconsMessage,
											Tooltip: "Send message",
											OnPressed: func() {
												cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
													Content: &glt.Text{"Imagine if you will, a messaging application."},
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
								Icon: glt.IconsContactMail,
								Children: []glt.Widget{
									&ContactItem{
										Icon:    glt.IconsEmail,
										Tooltip: "Send personal e-mail",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
												Content: &glt.Text{"Here, your e-mail application would open."},
											})
										},
										Lines: []string{
											"ali_connors@example.com",
											"Personal",
										},
									},
									&ContactItem{
										Icon:    glt.IconsEmail,
										Tooltip: "Send work e-mail",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
												Content: &glt.Text{"Summon your favorite e-mail application here."},
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
								Icon: glt.IconsLocationOn,
								Children: []glt.Widget{
									&ContactItem{
										Icon:    glt.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
												Content: &glt.Text{"This would show a map of San Francisco."},
											})
										},
										Lines: []string{
											"2000 Main Street",
											"San Francisco, CA",
											"Home",
										},
									},
									&ContactItem{
										Icon:    glt.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
												Content: &glt.Text{"This would show a map of Mountain View."},
											})
										},
										Lines: []string{
											"1600 Amphitheater Parkway",
											"Mountain View, CA",
											"Work",
										},
									},
									&ContactItem{
										Icon:    glt.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											cds.scaffoldKey.currentState.showSnackBar(&glt.SnackBar{
												Content: &glt.Text{"This would also show a map, if this was not a demo."},
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
								Icon: glt.IconsToday,
								Children: []glt.Widget{
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
