package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import (
	"path"
	"runtime"

	. "github.com/dragonfax/squares/squares"
)

var _ StatelessWidget = ContactCategory{}

type ContactCategory struct {
	Icon     *IconData
	Children []Widget
}

func (cc ContactCategory) Build(context StatelessContext) (Widget, error) {
	return DecoratedBox{
		Decoration: BoxDecoration{
			Border: Border{Bottom: BorderSide{}},
		},
		Child: Padding{
			Padding: EdgeInsetsSymmetric(16, 0),
			Child: Row{
				CrossAxisAlignment: CrossAxisAlignmentStart,
				Children: []Widget{
					SizedBox{
						Size: Size{Width: 72.0, Height: -1},
						Child: Padding{
							Padding: EdgeInsetsSymmetric(24, 0),
							Child:   Icon{Icon: *cc.Icon},
						},
					},
					Expanded{Child: Column{Children: cc.Children}},
				},
			},
		},
	}, nil
}

var _ StatelessWidget = ContactItem{}

type ContactItem struct {
	Icon      *IconData
	Lines     []string
	Tooltip   string
	OnPressed func()
}

func (ci ContactItem) Build(context StatelessContext) (Widget, error) {
	var columnChildren []Widget = make([]Widget, 0, len(ci.Lines))
	for _, line := range ci.Lines[0 : len(ci.Lines)-1] {
		columnChildren = append(columnChildren, Text{Text: line})
	}
	columnChildren = append(columnChildren, Text{Text: ci.Lines[len(ci.Lines)-1]})

	rowChildren := []Widget{
		Expanded{
			Child: Column{
				CrossAxisAlignment: CrossAxisAlignmentStart,
				Children:           columnChildren,
			},
		},
	}

	if ci.Icon != nil {
		rowChildren = append(rowChildren, SizedBox{
			Size: Size{
				Width:  72.0,
				Height: -1,
			},
			Child: IconButton{
				Icon:      Icon{Icon: *ci.Icon},
				OnPressed: ci.OnPressed,
			},
		})
	}
	return Padding{
		Padding: EdgeInsetsSymmetric(16, 0),
		Child: Row{
			MainAxisAlignment: MainAxisAlignmentSpaceBetween,
			Children:          rowChildren,
		},
	}, nil
}

var _ StatefulWidget = ContactsDemo{}

type ContactsDemo struct {
}

func (cd ContactsDemo) CreateState() State {
	return &ContactsDemoState{
		appBarHeight: 256.0,
	}
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
	appBarHeight float64
}

func (cds *ContactsDemoState) Build(context StatefulContext) (Widget, error) {
	var appBarBehavior AppBarBehavior = AppBarBehaviorPinned

	_, file, _, _ := runtime.Caller(0)
	assetPath := path.Dir(file) + "/assets/"

	return Scaffold{
		Body: CustomScrollView{
			Slivers: []Widget{
				SliverAppBar{
					ExpandedHeight: cds.appBarHeight,
					Pinned:         appBarBehavior == AppBarBehaviorPinned,
					Floating:       appBarBehavior == AppBarBehaviorFloating || appBarBehavior == AppBarBehaviorSnapping,
					Snap:           appBarBehavior == AppBarBehaviorSnapping,
					Actions: []Widget{
						IconButton{
							Icon:    Icon{Icon: IconsCreate},
							Tooltip: "Edit",
							OnPressed: func() {
								showSnackBar(context, SnackBar{
									Content: Text{"Editing isn't supported in this screen."},
								})
							},
						},
						PopupMenuButton{
							OnSelected: func(value interface{}) {
								context.SetState(func() {
									appBarBehavior = value.(AppBarBehavior)
								})
							},
							ItemBuilder: func(context StatelessContext) ([]PopupMenuItem, error) {
								return []PopupMenuItem{
									PopupMenuItem{
										Value: AppBarBehaviorNormal,
										Child: Text{"App bar scrolls away"},
									},
									PopupMenuItem{
										Value: AppBarBehaviorPinned,
										Child: Text{"App bar stays put"},
									},
									PopupMenuItem{
										Value: AppBarBehaviorFloating,
										Child: Text{"App bar floats"},
									},
									PopupMenuItem{
										Value: AppBarBehaviorSnapping,
										Child: Text{"App bar snaps"},
									},
								}, nil
							},
						},
					},
					FlexibleSpace: FlexibleSpaceBar{
						Title: Text{"Ali Connors"},
						Background: Stack{
							Fit: StackFitExpand,
							Children: []Widget{
								SizedBox{
									Size: Size{Height: cds.appBarHeight, Width: -1},
									Child: Image{
										File: assetPath + "people/ali_landscape.png",
										Fit:  BoxFitCover,
									},
								},
								// This gradient ensures that the toolbar icons are distinct
								// against the background image.
								DecoratedBox{
									Decoration: BoxDecoration{
										Gradient: LinearGradient{
											Begin:  Alignment{0.0, -1.0},
											End:    Alignment{0.0, -0.4},
											Colors: []Color{Color{0x60000000}, Color{0x00000000}},
										},
									},
								},
							},
						},
					},
				},
				SliverList{
					Delegate: SliverChildListDelegate{
						Children: []Widget{
							ContactCategory{
								Icon: &IconsCall,
								Children: []Widget{
									ContactItem{
										Icon:    &IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"Pretend that this opened your SMS application."},
											})
										},
										Lines: []string{
											"(650) 555-1234",
											"Mobile",
										},
									},
									ContactItem{
										Icon:    &IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"A messaging app appears."},
											})
										},
										Lines: []string{
											"(323) 555-6789",
											"Work",
										},
									},
									ContactItem{
										Icon:    &IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"Imagine if you will, a messaging application."},
											})
										},
										Lines: []string{
											"(650) 555-6789",
											"Home",
										},
									},
								},
							},
							ContactCategory{
								Icon: &IconsContactMail,
								Children: []Widget{
									ContactItem{
										Icon:    &IconsEmail,
										Tooltip: "Send personal e-mail",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"Here, your e-mail application would open."},
											})
										},
										Lines: []string{
											"ali_connors@example.com",
											"Personal",
										},
									},
									ContactItem{
										Icon:    &IconsEmail,
										Tooltip: "Send work e-mail",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"Summon your favorite e-mail application here."},
											})
										},
										Lines: []string{
											"aliconnors@example.com",
											"Work",
										},
									},
								},
							},
							ContactCategory{
								Icon: &IconsLocationOn,
								Children: []Widget{
									ContactItem{
										Icon:    &IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"This would show a map of San Francisco."},
											})
										},
										Lines: []string{
											"2000 Main Street",
											"San Francisco, CA",
											"Home",
										},
									},
									ContactItem{
										Icon:    &IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"This would show a map of Mountain View."},
											})
										},
										Lines: []string{
											"1600 Amphitheater Parkway",
											"Mountain View, CA",
											"Work",
										},
									},
									ContactItem{
										Icon:    &IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, SnackBar{
												Content: Text{"This would also show a map, if this was not a demo."},
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
							ContactCategory{
								Icon: &IconsToday,
								Children: []Widget{
									ContactItem{
										Lines: []string{
											"Birthday",
											"January 9th, 1989",
										},
									},
									ContactItem{
										Lines: []string{
											"Wedding anniversary",
											"June 21st, 2014",
										},
									},
									ContactItem{
										Lines: []string{
											"First day in office",
											"January 20th, 2015",
										},
									},
									ContactItem{
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
	}, nil
}

func showSnackBar(context BuildContext, snackBar SnackBar) {
	ContextOf(context, Scaffold{}).GetWidget().(Scaffold).ShowSnackBar(snackBar)
}
