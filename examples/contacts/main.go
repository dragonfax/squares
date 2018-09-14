package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import "github.com/dragonfax/gltr/gltr"

func main() {

	glt.RunApp(&gltr.MaterialApp{
		Title: "Contacts App",
		Color: gltr.ColorsGrey,
		Child: &ContactsDemo{},
	})
}
