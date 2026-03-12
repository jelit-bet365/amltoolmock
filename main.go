package main

import (
	"amltoolmock/handlers"
	"amltoolmock/middleware"
	"amltoolmock/services"
	"amltoolmock/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Load all mock data
	ds := services.GetDataService()
	if err := ds.LoadAllData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Create HTTP server
	mux := http.NewServeMux()

	// User Status endpoints
	mux.HandleFunc("/api/ecdd/userstatus", handleUsers)
	mux.HandleFunc("/api/ecdd/userstatus/", handleUserByID)

	// Threshold Config endpoints
	mux.HandleFunc("/api/ecdd/thresholdconfig", handleThresholds)
	mux.HandleFunc("/api/ecdd/thresholdconfig/", handleThresholdByID)

	// Case Management Folder endpoints (folder CRUD only)
	mux.HandleFunc("/api/ecdd/casemanagementfolder", handleFolders)
	mux.HandleFunc("/api/ecdd/casemanagementfolder/", handleCaseManagementFolderByID)

	// Multiplier Config endpoints
	mux.HandleFunc("/api/ecdd/multiplierconfig", handleMultipliers)
	mux.HandleFunc("/api/ecdd/multiplierconfig/", handleMultiplierByID)

	// Business Profile endpoints
	mux.HandleFunc("/api/ecdd/businessprofile", handleBusinessProfiles)
	mux.HandleFunc("/api/ecdd/businessprofile/", handleBusinessProfileByID)

	// User Case Management endpoints (folder assignments, folder users, folder stats, bulk ops)
	mux.HandleFunc("/api/ecdd/usercasemanagement", handleFolderAssignments)
	mux.HandleFunc("/api/ecdd/usercasemanagement/", handleUserCaseManagement)

	// Health check
	mux.HandleFunc("/health", middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"ECDD Mock API"}`)
	}))

	port := ":3001"
	fmt.Printf("\nECDD Mock API Server starting on port %s\n", port)
	fmt.Println("API Documentation: http://localhost:3001/health")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  - GET    /api/ecdd/userstatus")
	fmt.Println("  - GET    /api/ecdd/userstatus/{userstatuspk}")
	fmt.Println("  - POST   /api/ecdd/userstatus")
	fmt.Println("  - PUT    /api/ecdd/userstatus/{userstatuspk}")
	fmt.Println("  - PATCH  /api/ecdd/userstatus/{userstatuspk}")
	fmt.Println("  - DELETE /api/ecdd/userstatus/{userstatuspk}")
	fmt.Println("  - GET    /api/ecdd/userstatus/{userstatuspk}/folders")
	fmt.Println("  - GET    /api/ecdd/thresholdconfig")
	fmt.Println("  - GET    /api/ecdd/thresholdconfig/{thresholdconfigpk}")
	fmt.Println("  - POST   /api/ecdd/thresholdconfig")
	fmt.Println("  - PUT    /api/ecdd/thresholdconfig/{thresholdconfigpk}")
	fmt.Println("  - PATCH  /api/ecdd/thresholdconfig/{thresholdconfigpk}")
	fmt.Println("  - DELETE /api/ecdd/thresholdconfig/{thresholdconfigpk}")
	fmt.Println("  - GET    /api/ecdd/casemanagementfolder")
	fmt.Println("  - GET    /api/ecdd/casemanagementfolder/{folderpk}")
	fmt.Println("  - POST   /api/ecdd/casemanagementfolder")
	fmt.Println("  - PUT    /api/ecdd/casemanagementfolder/{folderpk}")
	fmt.Println("  - DELETE /api/ecdd/casemanagementfolder/{folderpk}")
	fmt.Println("  - GET    /api/ecdd/multiplierconfig")
	fmt.Println("  - GET    /api/ecdd/multiplierconfig/{multiplierconfigpk}")
	fmt.Println("  - POST   /api/ecdd/multiplierconfig")
	fmt.Println("  - PUT    /api/ecdd/multiplierconfig/{multiplierconfigpk}")
	fmt.Println("  - PATCH  /api/ecdd/multiplierconfig/{multiplierconfigpk}")
	fmt.Println("  - DELETE /api/ecdd/multiplierconfig/{multiplierconfigpk}")
	fmt.Println("  - GET    /api/ecdd/businessprofile")
	fmt.Println("  - GET    /api/ecdd/businessprofile/{businessprofilepk}")
	fmt.Println("  - POST   /api/ecdd/businessprofile")
	fmt.Println("  - PUT    /api/ecdd/businessprofile/{businessprofilepk}")
	fmt.Println("  - PATCH  /api/ecdd/businessprofile/{businessprofilepk}")
	fmt.Println("  - DELETE /api/ecdd/businessprofile/{businessprofilepk}")
	fmt.Println("  - GET    /api/ecdd/usercasemanagement")
	fmt.Println("  - POST   /api/ecdd/usercasemanagement")
	fmt.Println("  - DELETE /api/ecdd/usercasemanagement?folder_pk=")
	fmt.Println("  - DELETE /api/ecdd/usercasemanagement/{usercasemanagementpk}")
	fmt.Println("  - GET    /api/ecdd/usercasemanagement/folder/{folderpk}/users")
	fmt.Println("  - DELETE /api/ecdd/usercasemanagement/folder/{folderpk}/users/{userstatuspk}")
	fmt.Println("  - POST   /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-delete")
	fmt.Println("  - POST   /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-add")
	fmt.Println("  - GET    /api/ecdd/usercasemanagement/folder/{folderpk}/stats")
	fmt.Println("  - GET    /api/ecdd/usercasemanagement/stats")
	fmt.Println("\nServer ready!")

	log.Fatal(http.ListenAndServe(port, mux))
}

// Route handlers with method routing

func handleUsers(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllUsers(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/userstatus/")
		if path == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}

		// Support nested user routes:
		//   GET /api/ecdd/userstatus/{userstatuspk}/folders
		if strings.Contains(path, "/") {
			parts := strings.SplitN(path, "/", 2)
			if len(parts) == 2 && parts[1] == "folders" {
				if r.Method == http.MethodGet {
					handlers.GetUserFolders(w, r)
					return
				}
				utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
				return
			}
		}

		switch r.Method {
		case http.MethodGet:
			handlers.GetUserByID(w, r)
		case http.MethodPut:
			handlers.UpdateUser(w, r)
		case http.MethodPatch:
			handlers.PatchUser(w, r)
		case http.MethodDelete:
			handlers.DeleteUser(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleThresholds(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllThresholds(w, r)
		case http.MethodPost:
			handlers.CreateThreshold(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleThresholdByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/ecdd/thresholdconfig/") == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}
		switch r.Method {
		case http.MethodGet:
			handlers.GetThresholdByID(w, r)
		case http.MethodPut:
			handlers.UpdateThreshold(w, r)
		case http.MethodPatch:
			handlers.PatchThreshold(w, r)
		case http.MethodDelete:
			handlers.DeleteThreshold(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleFolders(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllFolders(w, r)
		case http.MethodPost:
			handlers.CreateFolder(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

// handleCaseManagementFolderByID handles folder CRUD only:
//
//	GET    /api/ecdd/casemanagementfolder/{folderpk}
//	PUT    /api/ecdd/casemanagementfolder/{folderpk}
//	DELETE /api/ecdd/casemanagementfolder/{folderpk}
func handleCaseManagementFolderByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/casemanagementfolder/")
		if path == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}

		switch r.Method {
		case http.MethodGet:
			handlers.GetFolderByID(w, r)
		case http.MethodPut:
			handlers.UpdateFolder(w, r)
		case http.MethodDelete:
			handlers.DeleteFolder(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleMultipliers(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllMultipliers(w, r)
		case http.MethodPost:
			handlers.CreateMultiplier(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleMultiplierByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/ecdd/multiplierconfig/") == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}
		switch r.Method {
		case http.MethodGet:
			handlers.GetMultiplierByID(w, r)
		case http.MethodPut:
			handlers.UpdateMultiplier(w, r)
		case http.MethodPatch:
			handlers.PatchMultiplier(w, r)
		case http.MethodDelete:
			handlers.DeleteMultiplier(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleBusinessProfiles(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllBusinessProfiles(w, r)
		case http.MethodPost:
			handlers.CreateBusinessProfile(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleBusinessProfileByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/ecdd/businessprofile/") == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}
		switch r.Method {
		case http.MethodGet:
			handlers.GetBusinessProfileByID(w, r)
		case http.MethodPut:
			handlers.UpdateBusinessProfile(w, r)
		case http.MethodPatch:
			handlers.PatchBusinessProfile(w, r)
		case http.MethodDelete:
			handlers.DeleteBusinessProfile(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

func handleFolderAssignments(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllFolderAssignments(w, r)
		case http.MethodPost:
			handlers.CreateFolderAssignment(w, r)
		case http.MethodDelete:
			handlers.DeleteFolderAssignmentsByFolder(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}

// handleUserCaseManagement routes all sub-paths under /api/ecdd/usercasemanagement/:
//
//	DELETE /api/ecdd/usercasemanagement/{usercasemanagementpk}
//	GET    /api/ecdd/usercasemanagement/stats
//	GET    /api/ecdd/usercasemanagement/folder/{folderpk}/users
//	DELETE /api/ecdd/usercasemanagement/folder/{folderpk}/users/{userstatuspk}
//	POST   /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-delete
//	POST   /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-add
//	GET    /api/ecdd/usercasemanagement/folder/{folderpk}/stats
func handleUserCaseManagement(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/ecdd/usercasemanagement/")
		if path == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID or sub-path required")
			return
		}

		// GET /api/ecdd/usercasemanagement/stats — bulk stats for all folders
		if path == "stats" {
			if r.Method == http.MethodGet {
				handlers.GetAllFolderAndStats(w, r)
				return
			}
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		// Routes under /api/ecdd/usercasemanagement/folder/{folderpk}/...
		if strings.HasPrefix(path, "folder/") {
			// Strip the "folder/" prefix so parts[0] == folderpk
			folderPath := strings.TrimPrefix(path, "folder/")
			parts := strings.SplitN(folderPath, "/", 4)

			if len(parts) >= 2 {
				switch parts[1] {
				case "stats":
					// GET /api/ecdd/usercasemanagement/folder/{folderpk}/stats
					if r.Method == http.MethodGet {
						handlers.GetFolderAndStats(w, r)
						return
					}
					utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
					return

				case "users":
					switch len(parts) {
					case 2:
						// GET /api/ecdd/usercasemanagement/folder/{folderpk}/users
						if r.Method == http.MethodGet {
							handlers.GetFolderUsers(w, r)
							return
						}
						utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
						return
					case 3:
						switch parts[2] {
						case "bulk-delete":
							// POST /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-delete
							if r.Method == http.MethodPost {
								handlers.BulkDeleteFolderUsers(w, r)
								return
							}
							utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
							return
						case "bulk-add":
							// POST /api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-add
							if r.Method == http.MethodPost {
								handlers.BulkAddFolderUsers(w, r)
								return
							}
							utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
							return
						default:
							// DELETE /api/ecdd/usercasemanagement/folder/{folderpk}/users/{userstatuspk}
							if parts[2] != "" {
								if r.Method == http.MethodDelete {
									handlers.DeleteFolderUser(w, r)
									return
								}
								utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
								return
							}
						}
					}
				}
			}
		}

		// DELETE /api/ecdd/usercasemanagement/{usercasemanagementpk}
		// path has no slashes — it is a plain PK
		if !strings.Contains(path, "/") && path != "" {
			if r.Method == http.MethodDelete {
				handlers.DeleteFolderAssignment(w, r)
				return
			}
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		utils.WriteJSONError(w, http.StatusNotFound, "Route not found")
	})
	handler(w, r)
}
