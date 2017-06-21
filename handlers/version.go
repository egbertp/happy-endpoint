package handlers

//import (
//"encoding/json"
//"net/http"
//)
//
//type VersionResponse struct {
//	Version        string      `json:"version"`
//	BuildTime           string      `json:"build_time"`
//	CommitHash          string      `json:"commit_hash"`
//}
//
//func VersionHandler(w http.ResponseWriter, r *http.Request) {
//
//	response := VersionResponse{
//		Version: maiVersionTag,
//		BuildTime: BuildTime,
//		CommitHash: CommitHash,
//
//	}
//
//
//	json.NewEncoder(w).Encode(response)
//	return
//}
//
