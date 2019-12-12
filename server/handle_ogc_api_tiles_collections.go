package server

import (
	"encoding/json"
	"net/http"
    "net/url"
    //"log"
    "github.com/go-spatial/geom"
    "github.com/go-spatial/tegola/atlas"
)

type OgcApiTilesCollections struct {
    Title string                       `json:"title"`
    Description string                 `json:"description"`
    Collections []CollectionMap        `json:"collections"`
    Links []LinkMap                    `json:"links"`
}

type CollectionMap struct {
	Id string               `json:"id"`
    Title string            `json:"title"`
    Description string      `json:"description"`
    Keywords []string       `json:"keywords"`
    Attribution string      `json:"attribution"`
    Extent *geom.Extent     `json:"extent"`
    Links []LinkMap         `json:"links"`
}

type HandleOgcApiTilesCollections struct{}

func (req HandleOgcApiTilesCollections) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	collections := OgcApiTilesCollections{
        Title: "OGC-API-Tiles",
        Description: "OGC API Tiles",
	}
    // parse our query string
    var query = r.URL.Query()

    var layerNames = make(map[string]string)

    // iterate our registered maps
    for _, m := range atlas.AllMaps() {
		debugQuery := url.Values{}
		// if we have a debug param add it to our URLs
		if query.Get("debug") == "true" {
			debugQuery.Set("debug", "true")

			// update our map to include the debug layers
			m = m.AddDebugLayers()
		}

        for i := range m.Layers {
            _, exists := layerNames[m.Layers[i].Name]
            if(!exists) {
        		cMap := CollectionMap{
        			Id:        m.Layers[i].Name,
        		    Title:     m.Layers[i].Name,
        		}
                tilesLink := LinkMap{
                    Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "collections", m.Layers[i].Name, "tiles"}, debugQuery),
                    Rel:        "tiles",
                    Type:       "application/json",
                }
                cMap.Links = append(cMap.Links, tilesLink)

        		// add the map to the capabilities struct
        		collections.Collections = append(collections.Collections, cMap)
                layerNames[m.Layers[i].Name] = ""
            }
        }
	}

    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")

	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(collections)
}
