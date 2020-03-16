package server

import (
	"encoding/json"
	"net/http"
    //"strings"
    "net/url"
    "github.com/dimfeld/httptreemux"
    //"github.com/go-spatial/tegola/atlas"
    //"fmt"
)

type TileMatrixSetLinkMap struct {
	TileMatrixSet string            `json:"tileMatrixSet"`
    TileMatrixSetURI string         `json:"tileMatrixSetURI"`
}

type OgcApiTilesTiles struct {
    Title string                                `json:"title"`
    Description string                          `json:"description"`
    Links []LinkMap                             `json:"links"`
    TileMatrixSetLinks []TileMatrixSetLinkMap   `json:"tileMatrixSetLinks"`
}

type HandleOgcApiTilesTiles struct{
}

func (req HandleOgcApiTilesTiles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    params := httptreemux.ContextParams(r.Context())

    mapTiles := OgcApiTilesTiles{
        Title: "OGC-API-Tiles",
        Description: "OGC API Tiles",
	}
    // parse our query string
	//var query = r.URL.Query()

	debugQuery := url.Values{}

    layerName := params["layer_name"]

    wgs84Link := TileMatrixSetLinkMap{
        TileMatrixSet:       "WorldCRS84Quad",
        TileMatrixSetURI:    "http://schemas.opengis.net/tms/1.0/json/examples/WorldCRS84Quad.json",
    }
    mercatorLink := TileMatrixSetLinkMap{
        TileMatrixSet:       "WebMercatorQuad",
        TileMatrixSetURI:    "http://schemas.opengis.net/tms/1.0/json/examples/WebMercatorQuad.json",
    }
    mapTiles.TileMatrixSetLinks = append(mapTiles.TileMatrixSetLinks, wgs84Link)
    mapTiles.TileMatrixSetLinks = append(mapTiles.TileMatrixSetLinks, mercatorLink)

    tilesLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"maps", "{tileMatrixSetId}", layerName, "{tileMatrix}/{tileCol}/{tileRow}.pbf"}, debugQuery),
        Rel:        "item",
        Type:       "application/vnd.mapbox-vector-tile",
        Title:      "Mapbox vector tiles",
    }
    mapTiles.Links = append(mapTiles.Links, tilesLink)

    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")

	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(mapTiles)
}
