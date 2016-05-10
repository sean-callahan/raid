package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/sean-callahan/raid"
)

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/index.tmpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		level := raid.Level0
		capacity := uint64(300)
		drives := uint64(2)

		tc, _ := level.TotalCapacity(capacity, drives)
		eff, _ := level.SpaceEfficiency(capacity, drives)
		ft, _ := level.FaultTolerance(drives)

		t.Execute(w, map[string]interface{}{
			"RAIDLevels":             raid.LevelText,
			"TotalCapacity":          tc,
			"TotalCapacityByteSize":  raid.ByteSize(tc * raid.Gigabyte),
			"SpaceEfficiency":        eff,
			"SpaceEfficiencyPercent": raid.Percent(eff),
			"FaultTolerance":         ft,
			"LastCapacity":           capacity,
			"LastDrives":             drives,
			"LastType":               level,
		})
	}
	if req.Method == http.MethodPost {
		t, err := template.ParseFiles("templates/index.tmpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = req.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rawLevel, err := strconv.Atoi(req.Form.Get("type"))
		capacity, err := strconv.ParseUint(req.Form.Get("capacity"), 10, 64)
		drives, err := strconv.ParseUint(req.Form.Get("drives"), 10, 64)
		if err != nil {
			http.Error(w, "malformed request", http.StatusBadRequest)
			return
		}

		level := raid.Level(rawLevel)

		tc, err := level.TotalCapacity(capacity, drives)
		eff, err := level.SpaceEfficiency(capacity, drives)
		ft, err := level.FaultTolerance(drives)

		t.Execute(w, map[string]interface{}{
			"Error":                  err,
			"RAIDLevels":             raid.LevelText,
			"TotalCapacity":          tc,
			"TotalCapacityByteSize":  raid.ByteSize(tc * raid.Gigabyte),
			"SpaceEfficiency":        eff,
			"SpaceEfficiencyPercent": raid.Percent(eff),
			"FaultTolerance":         ft,
			"LastCapacity":           capacity,
			"LastDrives":             drives,
			"LastType":               level,
		})
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
