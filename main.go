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
	mux.HandleFunc("/api/v1/users", handleUsers)
	mux.HandleFunc("/api/v1/users/", handleUserByID)

	// Threshold Config endpoints
	mux.HandleFunc("/api/v1/thresholds", handleThresholds)
	mux.HandleFunc("/api/v1/thresholds/", handleThresholdByID)

	// Case Folder endpoints
	mux.HandleFunc("/api/v1/folders", handleFolders)
	mux.HandleFunc("/api/v1/folders/", handleFolderByID)

	// Multiplier Config endpoints
	mux.HandleFunc("/api/v1/multipliers", handleMultipliers)
	mux.HandleFunc("/api/v1/multipliers/", handleMultiplierByID)

	// Business Profile endpoints
	mux.HandleFunc("/api/v1/business-profiles", handleBusinessProfiles)
	mux.HandleFunc("/api/v1/business-profiles/", handleBusinessProfileByID)

	// Folder Assignment endpoints
	mux.HandleFunc("/api/v1/folder-assignments", handleFolderAssignments)
	mux.HandleFunc("/api/v1/folder-assignments/", handleFolderAssignmentByID)

	// Health check
	mux.HandleFunc("/health", middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"ECDD Mock API"}`)
	}))

	port := ":3001"
	fmt.Printf("\n🚀 ECDD Mock API Server starting on port %s\n", port)
	fmt.Println("📝 API Documentation: http://localhost:3001/health")
	fmt.Println("\n Available endpoints:")
	fmt.Println("  - GET    /api/v1/users")
	fmt.Println("  - GET    /api/v1/users/{id}")
	fmt.Println("  - POST   /api/v1/users")
	fmt.Println("  - PUT    /api/v1/users/{id}")
	fmt.Println("  - PATCH  /api/v1/users/{id}")
	fmt.Println("  - DELETE /api/v1/users/{id}")
	fmt.Println("  - GET    /api/v1/users/{id}/folders")
	fmt.Println("  - GET    /api/v1/thresholds")
	fmt.Println("  - GET    /api/v1/thresholds/{id}")
	fmt.Println("  - POST   /api/v1/thresholds")
	fmt.Println("  - PUT    /api/v1/thresholds/{id}")
	fmt.Println("  - PATCH  /api/v1/thresholds/{id}")
	fmt.Println("  - DELETE /api/v1/thresholds/{id}")
	fmt.Println("  - GET    /api/v1/folders")
	fmt.Println("  - GET    /api/v1/folders/{id}")
	fmt.Println("  - GET    /api/v1/folders/{id}/users")
	fmt.Println("  - GET    /api/v1/folders/{id}/stats")
	fmt.Println("  - GET    /api/v1/folders/stats")
	fmt.Println("  - DELETE /api/v1/folders/{id}/users/{user_id}")
	fmt.Println("  - POST   /api/v1/folders/{id}/users/bulk-delete")
	fmt.Println("  - POST   /api/v1/folders/{id}/users/bulk-add")
	fmt.Println("  - POST   /api/v1/folders")
	fmt.Println("  - PUT    /api/v1/folders/{id}")
	fmt.Println("  - DELETE /api/v1/folders/{id}")
	fmt.Println("  - GET    /api/v1/multipliers")
	fmt.Println("  - GET    /api/v1/multipliers/{id}")
	fmt.Println("  - POST   /api/v1/multipliers")
	fmt.Println("  - PUT    /api/v1/multipliers/{id}")
	fmt.Println("  - PATCH  /api/v1/multipliers/{id}")
	fmt.Println("  - DELETE /api/v1/multipliers/{id}")
	fmt.Println("  - GET    /api/v1/business-profiles")
	fmt.Println("  - GET    /api/v1/business-profiles/{id}")
	fmt.Println("  - POST   /api/v1/business-profiles")
	fmt.Println("  - PUT    /api/v1/business-profiles/{id}")
	fmt.Println("  - PATCH  /api/v1/business-profiles/{id}")
	fmt.Println("  - DELETE /api/v1/business-profiles/{id}")
	fmt.Println("  - GET    /api/v1/folder-assignments")
	fmt.Println("  - POST   /api/v1/folder-assignments")
	fmt.Println("  - DELETE /api/v1/folder-assignments?folder_pk=")
	fmt.Println("  - DELETE /api/v1/folder-assignments/{id}")
	fmt.Println("\n✨ Server ready!")

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
		path := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
		if path == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}

		// Support nested user routes:
		//   GET /api/v1/users/{id}/folders
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
		if strings.TrimPrefix(r.URL.Path, "/api/v1/thresholds/") == "" {
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

func handleFolderByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/v1/folders/")
		if path == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}

		// GET /api/v1/folders/stats — bulk stats for all folders
		if path == "stats" {
			if r.Method == http.MethodGet {
				handlers.GetAllFolderAndStats(w, r)
				return
			}
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		// Support nested folder routes:
		//   GET    /api/v1/folders/{id}/users
		//   DELETE /api/v1/folders/{id}/users/{user_id}
		//   POST   /api/v1/folders/{id}/users/bulk-delete
		//   GET    /api/v1/folders/{id}/stats
		if strings.Contains(path, "/") {
			parts := strings.SplitN(path, "/", 4)
			if len(parts) >= 2 {
				switch parts[1] {
				case "stats":
					// GET /api/v1/folders/{id}/stats — stats for a single folder
					if r.Method == http.MethodGet {
						handlers.GetFolderAndStats(w, r)
						return
					}
					utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
					return

				case "users":
					switch len(parts) {
					case 2:
						// /api/v1/folders/{id}/users
						if r.Method == http.MethodGet {
							handlers.GetFolderUsers(w, r)
							return
						}
						utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
						return
					case 3:
						if parts[2] == "bulk-delete" {
							// /api/v1/folders/{id}/users/bulk-delete
							if r.Method == http.MethodPost {
								handlers.BulkDeleteFolderUsers(w, r)
								return
							}
							utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
							return
						}
						if parts[2] == "bulk-add" {
							// /api/v1/folders/{id}/users/bulk-add
							if r.Method == http.MethodPost {
								handlers.BulkAddFolderUsers(w, r)
								return
							}
							utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
							return
						}
						// /api/v1/folders/{id}/users/{user_id}
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
		if strings.TrimPrefix(r.URL.Path, "/api/v1/multipliers/") == "" {
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
		if strings.TrimPrefix(r.URL.Path, "/api/v1/business-profiles/") == "" {
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

func handleFolderAssignmentByID(w http.ResponseWriter, r *http.Request) {
	handler := middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/v1/folder-assignments/") == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "ID required")
			return
		}
		switch r.Method {
		case http.MethodDelete:
			handlers.DeleteFolderAssignment(w, r)
		default:
			utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	handler(w, r)
}
