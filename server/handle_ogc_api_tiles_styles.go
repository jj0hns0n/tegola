package server

import (
	"encoding/json"
	"net/http"
    "net/url"
)

type OgcApiTilesStyles struct {
	Links []LinkMap         `json:"links"`
    Styles []OgcApiStylesStruct `json:"styles"`
}

type OgcApiStylesStruct struct {
    Title string            `json:"title"`
    Links []LinkMap         `json:"links"`
    Id string               `json:"id"`
}

type HandleOgcApiTilesStyles struct{}

func (req HandleOgcApiTilesStyles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// new capabilities struct
	styles := OgcApiTilesStyles{
	}

    //Camo
    xrayStyleStruct := OgcApiStylesStruct{
        Title: "Xray",
        Id:    "xray",
    }

    xrayStyleLink := LinkMap{
        Href:       "https://ogc-vtp.gospatial.org/maps/WebMercatorQuad/style.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "Xray Style",
    }
    xrayStyleStruct.Links = append(xrayStyleStruct.Links, xrayStyleLink)

    styles.Styles = append(styles.Styles, xrayStyleStruct)

    //Camo
    camoStyleStruct := OgcApiStylesStruct{
        Title: "Camo",
        Id:    "camo",
    }

    camoStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/camo3d.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "Camo Style",
    }
    camoStyleStruct.Links = append(camoStyleStruct.Links, camoStyleLink)

    styles.Styles = append(styles.Styles, camoStyleStruct)

    //Go Spatial
    goSpatialStyleStruct := OgcApiStylesStruct{
        Title: "Go Spatial",
        Id:    "gospatial",
    }

    goSpatialStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/go-spatial.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "Go Spatial Style",
    }
    goSpatialStyleStruct.Links = append(goSpatialStyleStruct.Links, goSpatialStyleLink)

    styles.Styles = append(styles.Styles, goSpatialStyleStruct)

    //HOT OSM
    hotOSMStyleStruct := OgcApiStylesStruct{
        Title: "HOT OSM",
        Id:    "hotosm",
    }

    hotOSMStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/hot-osm3d.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "HOT OSM Style",
    }
    hotOSMStyleStruct.Links = append(hotOSMStyleStruct.Links, hotOSMStyleLink)

    styles.Styles = append(styles.Styles, hotOSMStyleStruct)

    //Mobility
    mobilityStyleStruct := OgcApiStylesStruct{
        Title: "Mobility",
        Id:    "mobility",
    }

    mobilityStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/mobility3d.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "Mobility Style",
    }
    mobilityStyleStruct.Links = append(mobilityStyleStruct.Links, mobilityStyleLink)

    styles.Styles = append(styles.Styles, mobilityStyleStruct)

    //Night
    nightStyleStruct := OgcApiStylesStruct{
        Title: "Night",
        Id:    "night",
    }

    nightStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/night-vision3d.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "Night Style",
    }
    nightStyleStruct.Links = append(nightStyleStruct.Links, nightStyleLink)

    styles.Styles = append(styles.Styles, nightStyleStruct)

    //OSGeo
    osgeoStyleStruct := OgcApiStylesStruct{
        Title: "OSGeo",
        Id:    "osgeo",
    }

    osgeoStyleLink := LinkMap{
        Href:       "https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/osgeo3d.json",
        Rel:        "stylesheet",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "OSGeo Style",
    }
    osgeoStyleStruct.Links = append(osgeoStyleStruct.Links, osgeoStyleLink)

    styles.Styles = append(styles.Styles, osgeoStyleStruct)

    debugQuery := url.Values{}

    selfLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "styles"}, debugQuery),
        Rel:        "self",
        Type:       "application/vnd.mapbox.style+json",
        Title:      "This Document",
    }
    styles.Links = append(styles.Links, selfLink)

    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")

	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(styles)
}
