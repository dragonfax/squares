package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import (
	"github.com/dragonfax/squares/squares"
)

var _ squares.StatelessWidget = &ContactCategory{}

type ContactCategory struct {
	Icon     *squares.IconData
	Children []squares.Widget
}

func (cc *ContactCategory) Build(context squares.BuildContext) (squares.Widget, error) {
	return &squares.Container{
		Padding: squares.EdgeInsetsSymmetric(16, 0),
		Decoration: squares.BoxDecoration{
			Border: squares.Border{Bottom: squares.BorderSide{}},
		},
		Child: &squares.Row{
			CrossAxisAlignment: squares.CrossAxisAlignmentStart,
			Children: []squares.Widget{
				&squares.Container{
					Padding: squares.EdgeInsetsSymmetric(24, 0),
					Width:   72.0,
					Child:   &squares.Icon{Icon: cc.Icon},
				},
				&squares.Expanded{Child: &squares.Column{Children: cc.Children}},
			},
		},
	}, nil
}

var _ squares.StatelessWidget = &ContactItem{}

type ContactItem struct {
	Icon      *squares.IconData
	Lines     []string
	Tooltip   string
	OnPressed func()
}

func (ci *ContactItem) Build(context squares.BuildContext) (squares.Widget, error) {
	var columnChildren []squares.Widget = make([]squares.Widget, 0, len(ci.Lines))
	for _, line := range ci.Lines[0 : len(ci.Lines)-1] {
		columnChildren = append(columnChildren, &squares.Text{Text: line})
	}
	columnChildren = append(columnChildren, &squares.Text{Text: ci.Lines[len(ci.Lines)-1]})

	rowChildren := []squares.Widget{
		&squares.Expanded{
			Child: &squares.Column{
				CrossAxisAlignment: squares.CrossAxisAlignmentStart,
				Children:           columnChildren,
			},
		},
	}

	if ci.Icon != nil {
		rowChildren = append(rowChildren, &squares.SizedBox{
			Width: 72.0,
			Child: &squares.IconButton{
				Icon:      &squares.Icon{Icon: ci.Icon},
				OnPressed: ci.OnPressed,
			},
		})
	}
	return &squares.Padding{
		Padding: squares.EdgeInsetsSymmetric(16, 0),
		Child: &squares.Row{
			MainAxisAlignment: squares.MainAxisAlignmentSpaceBetween,
			Children:          rowChildren,
		},
	}, nil
}

var _ squares.StatefulWidget = &ContactsDemo{}

type ContactsDemo struct {
}

func (cd *ContactsDemo) CreateState() squares.State {
	return &ContactsDemoState{}
}

type AppBarBehavior uint8

const (
	AppBarBehaviorNormal AppBarBehavior = iota
	AppBarBehaviorPinned
	AppBarBehaviorFloating
	AppBarBehaviorSnapping
)

var _ squares.State = &ContactsDemoState{}

type ContactsDemoState struct {
	appBarHeight uint16
}

// Use Init() in lue of a constructor
func (cds *ContactsDemoState) Init() {
	cds.appBarHeight = 256.0
}

func (cds *ContactsDemoState) Build(context squares.BuildContext) (squares.Widget, error) {
	var appBarBehavior AppBarBehavior = AppBarBehaviorPinned

	return &squares.Scaffold{
		Body: &squares.CustomScrollView{
			Slivers: []squares.Widget{
				&squares.SliverAppBar{
					ExpandedHeight: cds.appBarHeight,
					Pinned:         appBarBehavior == AppBarBehaviorPinned,
					Floating:       appBarBehavior == AppBarBehaviorFloating || appBarBehavior == AppBarBehaviorSnapping,
					Snap:           appBarBehavior == AppBarBehaviorSnapping,
					Actions: []squares.Widget{
						&squares.IconButton{
							Icon:    &squares.Icon{Icon: squares.IconsCreate},
							Tooltip: "Edit",
							OnPressed: func() {
								showSnackBar(context, &squares.SnackBar{
									Content: &squares.Text{"Editing isn't supported in this screen."},
								})
							},
						},
						&squares.PopupMenuButton{
							OnSelected: func(value interface{}) {
								context.(squares.StatefulContext).SetState(func() {
									appBarBehavior = value.(AppBarBehavior)
								})
							},
							ItemBuilder: func(context squares.BuildContext) ([]*squares.PopupMenuItem, error) {
								return []*squares.PopupMenuItem{
									&squares.PopupMenuItem{
										Value: AppBarBehaviorNormal,
										Child: &squares.Text{"App bar scrolls away"},
									},
									&squares.PopupMenuItem{
										Value: AppBarBehaviorPinned,
										Child: &squares.Text{"App bar stays put"},
									},
									&squares.PopupMenuItem{
										Value: AppBarBehaviorFloating,
										Child: &squares.Text{"App bar floats"},
									},
									&squares.PopupMenuItem{
										Value: AppBarBehaviorSnapping,
										Child: &squares.Text{"App bar snaps"},
									},
								}, nil
							},
						},
					},
					FlexibleSpace: &squares.FlexibleSpaceBar{
						Title: &squares.Text{"Ali Connors"},
						Background: &squares.Stack{
							Fit: squares.StackFitExpand,
							Children: []squares.Widget{
								squares.NewImageFromAsset(
									squares.Asset{
										File:    "people/ali_landscape.png",
										Package: "flutter_gallery_assets",
									},
									&squares.Image{
										Fit:    squares.BoxFitCover,
										Height: cds.appBarHeight,
									},
								),
								// This gradient ensures that the toolbar icons are distinct
								// against the background image.
								&squares.DecoratedBox{
									Decoration: squares.BoxDecoration{
										Gradient: squares.LinearGradient{
											Begin:  squares.Alignment{0.0, -1.0},
											End:    squares.Alignment{0.0, -0.4},
											Colors: []squares.Color{squares.Color{0x60000000}, squares.Color{0x00000000}},
										},
									},
								},
							},
						},
					},
				},
				&squares.SliverList{
					Delegate: &squares.SliverChildListDelegate{
						Children: []squares.Widget{
							&ContactCategory{
								Icon: squares.IconsCall,
								Children: []squares.Widget{
									&ContactItem{
										Icon:    squares.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"Pretend that this opened your SMS application."},
											})
										},
										Lines: []string{
											"(650) 555-1234",
											"Mobile",
										},
									},
									&ContactItem{
										Icon:    squares.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"A messaging app appears."},
											})
										},
										Lines: []string{
											"(323) 555-6789",
											"Work",
										},
									},
									&ContactItem{
										Icon:    squares.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"Imagine if you will, a messaging application."},
											})
										},
										Lines: []string{
											"(650) 555-6789",
											"Home",
										},
									},
								},
							},
							&ContactCategory{
								Icon: squares.IconsContactMail,
								Children: []squares.Widget{
									&ContactItem{
										Icon:    squares.IconsEmail,
										Tooltip: "Send personal e-mail",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"Here, your e-mail application would open."},
											})
										},
										Lines: []string{
											"ali_connors@example.com",
											"Personal",
										},
									},
									&ContactItem{
										Icon:    squares.IconsEmail,
										Tooltip: "Send work e-mail",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"Summon your favorite e-mail application here."},
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
								Icon: squares.IconsLocationOn,
								Children: []squares.Widget{
									&ContactItem{
										Icon:    squares.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"This would show a map of San Francisco."},
											})
										},
										Lines: []string{
											"2000 Main Street",
											"San Francisco, CA",
											"Home",
										},
									},
									&ContactItem{
										Icon:    squares.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"This would show a map of Mountain View."},
											})
										},
										Lines: []string{
											"1600 Amphitheater Parkway",
											"Mountain View, CA",
											"Work",
										},
									},
									&ContactItem{
										Icon:    squares.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &squares.SnackBar{
												Content: &squares.Text{"This would also show a map, if this was not a demo."},
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
								Icon: squares.IconsToday,
								Children: []squares.Widget{
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
	}, nil
}

func showSnackBar(context squares.BuildContext, snackBar *squares.SnackBar) {
	squares.ContextOf(context, &squares.Scaffold{}).GetWidget().(*squares.Scaffold).ShowSnackBar(snackBar)
}
