# Squares

_Any UI is really just a collection of overlapping colored squares_

Declarative, React-style UI framework for Golang, desktop applications.

## Example

```
Column{
    CrossAxisAlignment: CrossAxisAlignmentCenter,
    MainAxisAlignment:  MainAxisAlignmentCenter,
    Children: []Widget{
        Padding{
            Padding: EdgeInsetsAll(8), 
            Child: Text{Text: "Hello, World."}},
        Text{Text: "another row."},
        Text{Text: "yet another row."},
    },
},
```

## Status

Early alpha. Not really usable. You can create an app, it will render and even scroll. But it has no real visual style, and there are few widgets defined. Even those that are defined have few features implemented.

There is a simple version of composition (rendered widget caching) implemented in the form of a CompositionWidget that you can insert anywhere in the widget tree.

There is a scrollable widget, but it only scrolls vertically.

Stateful and Stateless apps work as expected, including using setState to efficiently rebuild the UI.

## Plans

Rather than translating each class or file in full from Dart to Go, I'm just re-implementing the classes one at a time including only minimal features, incrementally buidling up the widgets, and the features they support.

I'm currently working to port over the Contacts Demo from the Flutter Examples. Including all the widgets necessary to make that Demo work.

Rather than trying to copy the L&F of native widgets on any platform, I plan to keep with the Material Design . I feel Material Design is actually a good visual style that matches the feel of working with Go.

## Issues

### Default Values for Widgets

Using default values can't be auto-detected. Leaving out a properties in a struct just results in its Zero-value. And for some properties, the Zero-value might be a valid value for that property. Pointers to such properties could be used, but I think a pointer to an int or float is just bad form.

Within the code I've tried to keep Zero-values matching the standard default values used for each Widget, but thats not always possible. An example is the Flexible widget which defaults its Flex property to 1. A 0 Flex value makes no sense for this widget, so Init() could detect that and return an error when the user leaves this property out. But some widgets have properties where their zero-value is not the default, but at the same time, still a valid value for the property.

## TODO

* SetState should work concurrently via a work-queue
* bugs in Flex so that Expanded and Flexible widgets don't layout correctly
* more widgets and features
* layout and render are inefficient. redoing the whole tree every frame.


## Requirements

* SDL: 
  * `go get github.com/veandco/go-sdl2/sdl`
  * `go get github.com/veandco/go-sdl2/img`

## Screenshot
progress...

<img src="Contacts_App.png" width="200" >

## Inspiration

Squares is heavily inspired by Flutter (flutter.io).
