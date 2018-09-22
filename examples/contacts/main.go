package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import (
	. "github.com/dragonfax/squares/squares"
)

func main() {

	err := RunApp(
		Center{
			Child: MaterialApp{
				Title: "Contacts App",
				Color: ColorsGrey,
				Child: ContactsDemo{},
			}})

	if err != nil {
		panic(err)
	}
}
