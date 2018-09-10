package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.


import	"github.com/dragonfax/glitter/glt"

func main() {

  glt.RunApp(&glt.MaterialApp{
    Title: 'Contacts App',
    Color: Colors.grey,
    Child: &ContactsDemo{}
  ))
}
