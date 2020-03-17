package server

import (
	"encoding/json"
	"net/http"
    //"strings"
    "net/url"
    "github.com/dimfeld/httptreemux"
    "github.com/go-spatial/tegola/atlas"
    "github.com/go-spatial/geom"
    "github.com/go-spatial/tegola/internal/log"
    //"fmt"
)

type SpatialExtentStruct struct {
    Bbox []*geom.Extent                           `json:"bbox"`
    Crs string                                  `json:"crs"`
}

type ExtentStruct struct {
    SpatialExtent SpatialExtentStruct           `json:"spatial"`
}


type HandleOgcApiTilesCollection struct{
}

func (req HandleOgcApiTilesCollection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    params := httptreemux.ContextParams(r.Context())

    layerName := params["layer_name"]
    mapName := "WebMercatorQuad"
    m, err := atlas.GetMap(mapName)
    if err != nil {
        log.Errorf("map (%v) not configured. check your config file", mapName)
        http.Error(w, "map ("+mapName+") not configured. check your config file", http.StatusNotFound)
        return
    }

    extent := ExtentStruct{}
    spatialExtent := SpatialExtentStruct{}
    spatialExtent.Bbox = append(spatialExtent.Bbox, m.Bounds)
    spatialExtent.Crs = "http://www.opengis.net/def/crs/OGC/1.3/CRS84"
    extent.SpatialExtent = spatialExtent


    collection := CollectionMap{
        Id: layerName,
        Title: layerName,
        Description: layerName,
        Extent: extent,
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
        Type:       "application/json",
        Title:      "Queryable attributes",
    }
    collection.Links = append(collection.Links, queryablesLink)

    tilesLink := LinkMap{
        Href:       buildCapabilitiesURL(r, []string{"ogc-api-tiles", "collections", layerName, "tiles"}, debugQuery),
        Rel:        "tiles",
        Type:       "application/json",
        Title:      "Access the data as vector tiles",
    }
    collection.Links = append(collection.Links, tilesLink)

    w.Header().Add("Content-Type", "application/json")

    // cache control headers (no-cache)
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Add("Pragma", "no-cache")
    w.Header().Add("Expires", "0")

	// setup a new json encoder and encode our capabilities
	json.NewEncoder(w).Encode(collection)
}
