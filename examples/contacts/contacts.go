package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import (
	"github.com/dragonfax/gltr/gltr"
)

var _ gltr.StatelessWidget = &ContactCategory{}

type ContactCategory struct {
	Icon     *gltr.IconData
	Children []gltr.Widget
}

func (cc *ContactCategory) Build(context gltr.BuildContext) (gltr.Widget, error) {
	return &gltr.Container{
		Padding: gltr.EdgeInsetsSymmetric(16, 0),
		Decoration: gltr.BoxDecoration{
			Border: gltr.Border{Bottom: gltr.BorderSide{}},
		},
		Child: &gltr.Row{
			CrossAxisAlignment: gltr.CrossAxisAlignmentStart,
			Children: []gltr.Widget{
				&gltr.Container{
					Padding: gltr.EdgeInsetsSymmetric(24, 0),
					Width:   72.0,
					Child:   &gltr.Icon{Icon: cc.Icon},
				},
				&gltr.Expanded{Child: &gltr.Column{Children: cc.Children}},
			},
		},
	}, nil
}

var _ gltr.StatelessWidget = &ContactItem{}

type ContactItem struct {
	Icon      *gltr.IconData
	Lines     []string
	Tooltip   string
	OnPressed func()
}

func (ci *ContactItem) Build(context gltr.BuildContext) (gltr.Widget, error) {
	var columnChildren []gltr.Widget = make([]gltr.Widget, 0, len(ci.Lines))
	for _, line := range ci.Lines[0 : len(ci.Lines)-1] {
		columnChildren = append(columnChildren, &gltr.Text{Text: line})
	}
	columnChildren = append(columnChildren, &gltr.Text{Text: ci.Lines[len(ci.Lines)-1]})

	rowChildren := []gltr.Widget{
		&gltr.Expanded{
			Child: &gltr.Column{
				CrossAxisAlignment: gltr.CrossAxisAlignmentStart,
				Children:           columnChildren,
			},
		},
	}

	if ci.Icon != nil {
		rowChildren = append(rowChildren, &gltr.SizedBox{
			Width: 72.0,
			Child: &gltr.IconButton{
				Icon:      &gltr.Icon{Icon: ci.Icon},
				OnPressed: ci.OnPressed,
			},
		})
	}
	return &gltr.Padding{
		Padding: gltr.EdgeInsetsSymmetric(16, 0),
		Child: &gltr.Row{
			MainAxisAlignment: gltr.MainAxisAlignmentSpaceBetween,
			Children:          rowChildren,
		},
	}, nil
}

var _ gltr.StatefulWidget = &ContactsDemo{}

type ContactsDemo struct {
}

func (cd *ContactsDemo) CreateState() gltr.State {
	return &ContactsDemoState{}
}

type AppBarBehavior uint8

const (
	AppBarBehaviorNormal AppBarBehavior = iota
	AppBarBehaviorPinned
	AppBarBehaviorFloating
	AppBarBehaviorSnapping
)

var _ gltr.State = &ContactsDemoState{}

type ContactsDemoState struct {
	appBarHeight uint16
}

// Use Init() in lue of a constructor
func (cds *ContactsDemoState) Init() {
	cds.appBarHeight = 256.0
}

func (cds *ContactsDemoState) Build(context gltr.BuildContext) (gltr.Widget, error) {
	var appBarBehavior AppBarBehavior = AppBarBehaviorPinned

	return &gltr.Scaffold{
		Body: &gltr.CustomScrollView{
			Slivers: []gltr.Widget{
				&gltr.SliverAppBar{
					ExpandedHeight: cds.appBarHeight,
					Pinned:         appBarBehavior == AppBarBehaviorPinned,
					Floating:       appBarBehavior == AppBarBehaviorFloating || appBarBehavior == AppBarBehaviorSnapping,
					Snap:           appBarBehavior == AppBarBehaviorSnapping,
					Actions: []gltr.Widget{
						&gltr.IconButton{
							Icon:    &gltr.Icon{Icon: gltr.IconsCreate},
							Tooltip: "Edit",
							OnPressed: func() {
								showSnackBar(context, &gltr.SnackBar{
									Content: &gltr.Text{"Editing isn't supported in this screen."},
								})
							},
						},
						&gltr.PopupMenuButton{
							OnSelected: func(value interface{}) {
								context.(gltr.StatefulContext).SetState(func() {
									appBarBehavior = value.(AppBarBehavior)
								})
							},
							ItemBuilder: func(context gltr.BuildContext) ([]*gltr.PopupMenuItem, error) {
								return []*gltr.PopupMenuItem{
									&gltr.PopupMenuItem{
										Value: AppBarBehaviorNormal,
										Child: &gltr.Text{"App bar scrolls away"},
									},
									&gltr.PopupMenuItem{
										Value: AppBarBehaviorPinned,
										Child: &gltr.Text{"App bar stays put"},
									},
									&gltr.PopupMenuItem{
										Value: AppBarBehaviorFloating,
										Child: &gltr.Text{"App bar floats"},
									},
									&gltr.PopupMenuItem{
										Value: AppBarBehaviorSnapping,
										Child: &gltr.Text{"App bar snaps"},
									},
								}, nil
							},
						},
					},
					FlexibleSpace: &gltr.FlexibleSpaceBar{
						Title: &gltr.Text{"Ali Connors"},
						Background: &gltr.Stack{
							Fit: gltr.StackFitExpand,
							Children: []gltr.Widget{
								gltr.NewImageFromAsset(
									gltr.Asset{
										File:    "people/ali_landscape.png",
										Package: "flutter_gallery_assets",
									},
									&gltr.Image{
										Fit:    gltr.BoxFitCover,
										Height: cds.appBarHeight,
									},
								),
								// This gradient ensures that the toolbar icons are distinct
								// against the background image.
								&gltr.DecoratedBox{
									Decoration: gltr.BoxDecoration{
										Gradient: gltr.LinearGradient{
											Begin:  gltr.Alignment{0.0, -1.0},
											End:    gltr.Alignment{0.0, -0.4},
											Colors: []gltr.Color{gltr.Color{0x60000000}, gltr.Color{0x00000000}},
										},
									},
								},
							},
						},
					},
				},
				&gltr.SliverList{
					Delegate: &gltr.SliverChildListDelegate{
						Children: []gltr.Widget{
							&ContactCategory{
								Icon: gltr.IconsCall,
								Children: []gltr.Widget{
									&ContactItem{
										Icon:    gltr.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"Pretend that this opened your SMS application."},
											})
										},
										Lines: []string{
											"(650) 555-1234",
											"Mobile",
										},
									},
									&ContactItem{
										Icon:    gltr.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"A messaging app appears."},
											})
										},
										Lines: []string{
											"(323) 555-6789",
											"Work",
										},
									},
									&ContactItem{
										Icon:    gltr.IconsMessage,
										Tooltip: "Send message",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"Imagine if you will, a messaging application."},
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
								Icon: gltr.IconsContactMail,
								Children: []gltr.Widget{
									&ContactItem{
										Icon:    gltr.IconsEmail,
										Tooltip: "Send personal e-mail",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"Here, your e-mail application would open."},
											})
										},
										Lines: []string{
											"ali_connors@example.com",
											"Personal",
										},
									},
									&ContactItem{
										Icon:    gltr.IconsEmail,
										Tooltip: "Send work e-mail",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"Summon your favorite e-mail application here."},
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
								Icon: gltr.IconsLocationOn,
								Children: []gltr.Widget{
									&ContactItem{
										Icon:    gltr.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"This would show a map of San Francisco."},
											})
										},
										Lines: []string{
											"2000 Main Street",
											"San Francisco, CA",
											"Home",
										},
									},
									&ContactItem{
										Icon:    gltr.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"This would show a map of Mountain View."},
											})
										},
										Lines: []string{
											"1600 Amphitheater Parkway",
											"Mountain View, CA",
											"Work",
										},
									},
									&ContactItem{
										Icon:    gltr.IconsMap,
										Tooltip: "Open map",
										OnPressed: func() {
											showSnackBar(context, &gltr.SnackBar{
												Content: &gltr.Text{"This would also show a map, if this was not a demo."},
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
								Icon: gltr.IconsToday,
								Children: []gltr.Widget{
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

func showSnackBar(context gltr.BuildContext, snackBar *gltr.SnackBar) {
	gltr.ContextOf(context, &gltr.Scaffold{}).GetWidget().(*gltr.Scaffold).ShowSnackBar(snackBar)
}
