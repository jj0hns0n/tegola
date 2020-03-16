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

type OgcApiTilesCollection struct {
    Title string                                `json:"title"`
    Description string                          `json:"description"`
    Links []LinkMap                             `json:"links"`
    //TileMatrixSetLinks []TileMatrixSetLinkMap   `json:"tileMatrixSetLinks"`
}

type HandleOgcApiTilesCollection struct{
}

func (req HandleOgcApiTilesCollection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    params := httptreemux.ContextParams(r.Context())

    layerName := params["layer_name"]
    //mapName := params[":map_name"]

    mapTiles := OgcApiTilesTiles{
        Title: layerName,
        Description: layerName,
	}
    // parse our query string
	//var query = r.URL.Query()

	debugQuery := url.Values{}




    //wgs84Link := TileMatrixSetLinkMap{
    //    TileMatrixSet:       "WorldCRS84Quad",
    //    TileMatrixSetURI:    "http://schemas.opengis.net/tms/1.0/json/examples/WorldCRS84Quad.json",
    //}
    //mercatorLink := TileMatrixSetLinkMap{
    //    TileMatrixSet:       "WebMercatorQuad",
    //    TileMatrixSetURI:    "http://schemas.opengis.net/tms/1.0/json/examples/WebMercatorQuad.json",
    //}
    //mapTiles.TileMatrixSetLinks = append(mapTiles.TileMatrixSetLinks, wgs84Link)
    //mapTiles.TileMatrixSetLinks = append(mapTiles.TileMatrixSetLinks, mercatorLink)

    queryablesLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "collections", layerName, "queryables"}, debugQuery),
        Rel:        "queryables",
        Title:      "Queryable attributes",
    }
    mapTiles.Links = append(mapTiles.Links, queryablesLink)

    tilesLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "collections", layerName, "tiles"}, debugQuery),
        Rel:        "tiles",
        Title:      "Access the data as vector tiles",
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
