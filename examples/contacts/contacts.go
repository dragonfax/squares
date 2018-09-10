package main

// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import	"github.com/dragonfax/glitter/glt"

var _ StatelessWidget = &ContactCategory{} 

type ContactCategory struct {
  icon IconData
  Children []Widget
}

func (cc *ContactCategory) Build(context glt.BuildContext) (Widget, error) {
    return &glt.Container{
      Padding: EdgeInsets{Vertical: 16.0},
      Decoration: &BoxDecoration{
        Border: &Border{Bottom: &BorderSide{Color: themeData.dividerColor}}
      },
      Child: &DefaultTextStyle{
        Style: Theme.of(context).textTheme.subhead,
        Child: &SafeArea{
          Top: false,
          Bottom: false,
          Child: &Row{
            CrossAxisAlignment: CrossAxisAlignment.start,
            Children: []Widget{
              &Container{
                Padding: EdgeInsets{Vertical: 24.0},
                Width: 72.0,
                Child: &Icon{Icon: icon, Color: themeData.primaryColor},
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
  Icon IconData
  lines []string
  Tooltip string
  onPressed VoidCallback
}

func (ci *ContactItem) Build(context glt.BuildContext) (Widget,error) {
    // ThemeData themeData = Theme.of(context);
    var columnChildren []Widget = make([]Widget,0, len(lines))
    for _, line := range lines[0:-1] {
      columnChildren = append(columnChildren,&glt.Text{Text: line})
    }
    columnChildren = append(columnChildren, &Text{Text: lines[-1], Style: themeData.textTheme.caption})

    rowChildren := []Widget{
      &glt.Expanded{
        Child: &Column{
          CrossAxisAlignment: CrossAxisAlignment.start,
          Children: columnChildren,
        }
      }
    };

    if icon != nil {
      rowChildren = append(rowChildren, &SizedBox{
        Width: 72.0,
        Child: &IconButton{
          Icon: &Icon{Icon: icon},
          Color: themeData.primaryColor,
          OnPressed: onPressed,
        }
      });
    }
    return &MergeSemantics{
      Child: &Padding{
        Padding: &EdgeInsets{Vertical: 16.0},
        Child: &Row{
          MainAxisAlignment: MainAxisAlignment.spaceBetween,
          Children: rowChildren,
        },
      },
    };
}

class ContactsDemo extends StatefulWidget {
  static const String routeName = '/contacts';

  @override
  ContactsDemoState createState() => new ContactsDemoState();
}

enum AppBarBehavior { normal, pinned, floating, snapping }

class ContactsDemoState extends State<ContactsDemo> {
  static final GlobalKey<ScaffoldState> _scaffoldKey = new GlobalKey<ScaffoldState>();
  final double _appBarHeight = 256.0;

  AppBarBehavior _appBarBehavior = AppBarBehavior.pinned;

  @override
  Widget build(BuildContext context) {
    return new Theme(
      data: new ThemeData(
        brightness: Brightness.light,
        primarySwatch: Colors.indigo,
        platform: Theme.of(context).platform,
      ),
      child: new Scaffold(
        key: _scaffoldKey,
        body: new CustomScrollView(
          slivers: <Widget>[
            new SliverAppBar(
              expandedHeight: _appBarHeight,
              pinned: _appBarBehavior == AppBarBehavior.pinned,
              floating: _appBarBehavior == AppBarBehavior.floating || _appBarBehavior == AppBarBehavior.snapping,
              snap: _appBarBehavior == AppBarBehavior.snapping,
              actions: <Widget>[
                new IconButton(
                  icon: const Icon(Icons.create),
                  tooltip: 'Edit',
                  onPressed: () {
                    _scaffoldKey.currentState.showSnackBar(const SnackBar(
                      content: Text("Editing isn't supported in this screen.")
                    ));
                  },
                ),
                new PopupMenuButton<AppBarBehavior>(
                  onSelected: (AppBarBehavior value) {
                    setState(() {
                      _appBarBehavior = value;
                    });
                  },
                  itemBuilder: (BuildContext context) => <PopupMenuItem<AppBarBehavior>>[
                    const PopupMenuItem<AppBarBehavior>(
                      value: AppBarBehavior.normal,
                      child: Text('App bar scrolls away')
                    ),
                    const PopupMenuItem<AppBarBehavior>(
                      value: AppBarBehavior.pinned,
                      child: Text('App bar stays put')
                    ),
                    const PopupMenuItem<AppBarBehavior>(
                      value: AppBarBehavior.floating,
                      child: Text('App bar floats')
                    ),
                    const PopupMenuItem<AppBarBehavior>(
                      value: AppBarBehavior.snapping,
                      child: Text('App bar snaps')
                    ),
                  ],
                ),
              ],
              flexibleSpace: new FlexibleSpaceBar(
                title: const Text('Ali Connors'),
                background: new Stack(
                  fit: StackFit.expand,
                  children: <Widget>[
                    new Image.asset(
                      'people/ali_landscape.png',
                      package: 'flutter_gallery_assets',
                      fit: BoxFit.cover,
                      height: _appBarHeight,
                    ),
                    // This gradient ensures that the toolbar icons are distinct
                    // against the background image.
                    const DecoratedBox(
                      decoration: BoxDecoration(
                        gradient: LinearGradient(
                          begin: Alignment(0.0, -1.0),
                          end: Alignment(0.0, -0.4),
                          colors: <Color>[Color(0x60000000), Color(0x00000000)],
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ),
            new SliverList(
              delegate: new SliverChildListDelegate(<Widget>[
                new AnnotatedRegion<SystemUiOverlayStyle>(
                  value: SystemUiOverlayStyle.dark,
                  child: new _ContactCategory(
                    icon: Icons.call,
                    children: <Widget>[
                      new _ContactItem(
                        icon: Icons.message,
                        tooltip: 'Send message',
                        onPressed: () {
                          _scaffoldKey.currentState.showSnackBar(const SnackBar(
                            content: Text('Pretend that this opened your SMS application.')
                          ));
                        },
                        lines: const <String>[
                          '(650) 555-1234',
                          'Mobile',
                        ],
                      ),
                      new _ContactItem(
                        icon: Icons.message,
                        tooltip: 'Send message',
                        onPressed: () {
                          _scaffoldKey.currentState.showSnackBar(const SnackBar(
                            content: Text('A messaging app appears.')
                          ));
                        },
                        lines: const <String>[
                          '(323) 555-6789',
                          'Work',
                        ],
                      ),
                      new _ContactItem(
                        icon: Icons.message,
                        tooltip: 'Send message',
                        onPressed: () {
                          _scaffoldKey.currentState.showSnackBar(const SnackBar(
                            content: Text('Imagine if you will, a messaging application.')
                          ));
                        },
                        lines: const <String>[
                          '(650) 555-6789',
                          'Home',
                        ],
                      ),
                    ],
                  ),
                ),
                new _ContactCategory(
                  icon: Icons.contact_mail,
                  children: <Widget>[
                    new _ContactItem(
                      icon: Icons.email,
                      tooltip: 'Send personal e-mail',
                      onPressed: () {
                        _scaffoldKey.currentState.showSnackBar(const SnackBar(
                          content: Text('Here, your e-mail application would open.')
                        ));
                      },
                      lines: const <String>[
                        'ali_connors@example.com',
                        'Personal',
                      ],
                    ),
                    new _ContactItem(
                      icon: Icons.email,
                      tooltip: 'Send work e-mail',
                      onPressed: () {
                        _scaffoldKey.currentState.showSnackBar(const SnackBar(
                          content: Text('Summon your favorite e-mail application here.')
                        ));
                      },
                      lines: const <String>[
                        'aliconnors@example.com',
                        'Work',
                      ],
                    ),
                  ],
                ),
                new _ContactCategory(
                  icon: Icons.location_on,
                  children: <Widget>[
                    new _ContactItem(
                      icon: Icons.map,
                      tooltip: 'Open map',
                      onPressed: () {
                        _scaffoldKey.currentState.showSnackBar(const SnackBar(
                          content: Text('This would show a map of San Francisco.')
                        ));
                      },
                      lines: const <String>[
                        '2000 Main Street',
                        'San Francisco, CA',
                        'Home',
                      ],
                    ),
                    new _ContactItem(
                      icon: Icons.map,
                      tooltip: 'Open map',
                      onPressed: () {
                        _scaffoldKey.currentState.showSnackBar(const SnackBar(
                          content: Text('This would show a map of Mountain View.')
                        ));
                      },
                      lines: const <String>[
                        '1600 Amphitheater Parkway',
                        'Mountain View, CA',
                        'Work',
                      ],
                    ),
                    new _ContactItem(
                      icon: Icons.map,
                      tooltip: 'Open map',
                      onPressed: () {
                        _scaffoldKey.currentState.showSnackBar(const SnackBar(
                          content: Text('This would also show a map, if this was not a demo.')
                        ));
                      },
                      lines: const <String>[
                        '126 Severyns Ave',
                        'Mountain View, CA',
                        'Jet Travel',
                      ],
                    ),
                  ],
                ),
                new _ContactCategory(
                  icon: Icons.today,
                  children: <Widget>[
                    new _ContactItem(
                      lines: const <String>[
                        'Birthday',
                        'January 9th, 1989',
                      ],
                    ),
                    new _ContactItem(
                      lines: const <String>[
                        'Wedding anniversary',
                        'June 21st, 2014',
                      ],
                    ),
                    new _ContactItem(
                      lines: const <String>[
                        'First day in office',
                        'January 20th, 2015',
                      ],
                    ),
                    new _ContactItem(
                      lines: const <String>[
                        'Last day in office',
                        'August 9th, 2018',
                      ],
                    ),
                  ],
                ),
              ]),
            ),
          ],
        ),
      ),
    );
  }
}
