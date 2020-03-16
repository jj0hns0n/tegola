package server

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/url"

	"github.com/dimfeld/httptreemux"

	"github.com/go-spatial/tegola/atlas"
	"github.com/go-spatial/tegola/internal/log"
)

type HandleLayerSchema struct {
	// required
	mapName string
	// required
	layerName string
	// the requests extension defaults to "json"
	extension string
}

// returns layer schema
//
// URI scheme: /:map_name/:layer_name/queryables
// 	layer_name - layer name in the config file
func (req HandleLayerSchema) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	params := httptreemux.ContextParams(r.Context())

    fmt.Println(req)
	// read the map_name value from the request
	req.mapName = params["map_name"]
	req.layerName = params["layer_name"]

	// lookup our Map
	m, err := atlas.GetMap(req.mapName)
	if err != nil {
		log.Errorf("map (%v) not configured. check your config file", req.mapName)
		http.Error(w, "map ("+req.mapName+") not configured. check your config file", http.StatusNotFound)
		return
	}


    if req.layerName != "" {
       m = m.FilterLayersByName(req.layerName)
       if len(m.Layers) == 0 {
            logAndError(w, http.StatusNotFound, "map (%v) has no layers, for LayerName %v", req.mapName, req.layerName)
            return
        }
    }

    for _, l := range m.Layers {
        fmt.Printf("LLL: %s -- %s\n", l.Name, l.ProviderLayerName)
        if l.Name == req.layerName {
            players, _ := l.Provider.LayerSchema(req.layerName, {})
        }
    }

	// if we have a debug param add it to our URLs
	debugQuery := url.Values{}
	if r.URL.Query().Get("debug") == "true" {
		debugQuery.Set("debug", "true")

		// update our map to include the debug layers
		m = m.AddDebugLayers()
	}

	// mimetype for protocol buffers
	w.Header().Add("Content-Type", "application/json")

	// cache control headers (no-cache)
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Expires", "0")

	if err = json.NewEncoder(w).Encode(req); err != nil {
		log.Errorf("error encoding tileJSON for map (%v)", req.mapName)
	}
}
