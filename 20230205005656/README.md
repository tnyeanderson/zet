# Parsing XML maps in go

Parsing XML with the `encoding/xml` go package is not always straightforward.

Say you have an XML element `<mapelement>` which essentially represents a map:
child elements of this will have unique tags, each representing an instance of
some resource/item (`<item1>`, `<item2>`).

A minimum XML example element is:

```xml
<mapelement>
  <item1>
    <value>hello</value>
  </item1>
  <item2>
    <value>goodbye</value>
  </item2>
</mapelement>
```

You conceptualize this as the equivalent JSON:

```json
{
  "mapelement": {
    "item1": {
      "value": "hello"
    },
    "item2": {
      "value": "hello"
    }
  }
}
```

In this case, see the annotated go code below:

```go
package main

import (
	"encoding/xml"
	"fmt"
)

var xmltest = []byte(`<?xml version="1.0"?>
<rootelement>
  <mapelement>
    <item1>
      <itemprop>
				<value>hello</value>
      </itemprop>
      <itemprop>
				<value>goodbye</value>
      </itemprop>
		</item1>
		<item2>
      <itemprop>
				<value>hola</value>
      </itemprop>
      <itemprop>
				<value>adios</value>
      </itemprop>
		</item2>
  </mapelement>
</rootelement>
`)

// The XML <rootelement>. This is what will be passed to the xml.Unmarshal
// function!
type RootElement struct {
	MapElement MapElement `xml:"mapelement"`
}

// This element contains what is essentially a map, where the XML tag of each
// child element is the key, and the object contained in that child element is
// its value. The XML package does not support maps like this natively.
// However, the `xml:",any"` struct tag will grab all of the remaining
// unmatched child elements (all of them in this case) and assemble them into
// an array. The XMLName for each Item will have the name of the XML tag of the
// child element (what would be the map key if it were represented that way).
//
// Below in the main() function, there is an example of how to translate this
// XML-ified data structure into the map[string]Item that we conceptualize it
// as.
type MapElement struct {
	Items []Item `xml:",any"`
}

// Each item will be unmarshaled here. The XMLName propery (of type xml.Name)
// will contain the tag name of the containing element (in this case "item1",
// "item2", etc).
type Item struct {
	XMLName   xml.Name
	ItemProps []ItemProp `xml:"itemprop"`
}

// A repeated item property, for the sake of example.
type ItemProp struct {
	Value string `xml:"value"`
}

func main() {
	// The root element is always where the XML parser starts. Don't try to
	// unmarshal into a "document-level" struct with RootElement as a property,
	// it won't work! The marshaling is done starting at the contents of the root
	// element.
	root := RootElement{}

	// Not checking for errors here since this is an example
	xml.Unmarshal(xmltest, &root)

	// At this point, we can loop through the Items array and do whatever it is
	// we plan on doing. Unforunately, the data structure isn't exactly how we
	// conceptualize it (a map[string]Item), but this can be fixed by looping
	// through the properties later and assembling that map.
	itemsMap := map[string]Item{}
	for _, item := range root.MapElement.Items {
		// Be sure to use XMLName.Local to get the local tag name
		name := item.XMLName.Local
		itemsMap[name] = item
	}
	fmt.Printf("itemsMap: %+v\n", itemsMap)

	// If we want a map of item names to item properties, looping through the
	// results is required.
	itemProps := map[string][]ItemProp{}
	for _, item := range root.MapElement.Items {
		// Be sure to use XMLName.Local to get the local tag name
		name := item.XMLName.Local
		itemProps[name] = item.ItemProps
	}
	fmt.Printf("itemProps: %+v\n", itemProps)
}
```

This will output:

```
itemsMap: map[item1:{XMLName:{Space: Local:item1} ItemProps:[{Value:hello} {Value:goodbye}]} item2:{XMLName:{Space: Local:item2} ItemProps:[{Value:hola} {Value:adios}]}]
itemProps: map[item1:[{Value:hello} {Value:goodbye}] item2:[{Value:hola} {Value:adios}]]
```

    #go #golang #xml
