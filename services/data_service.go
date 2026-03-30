package services

import (
	"amltoolmock/models"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

// DataService manages all in-memory data with thread-safe operations
type DataService struct {
	ThresholdConfigs  map[string]*models.ECDDThresholdConfig
	UserStatuses      map[string]*models.ECDDUserStatus
	CaseFolders       map[string]*models.ECDDCaseManagementFolder
	MultiplierConfigs map[string]*models.ECDDMultiplierConfig
	BusinessProfiles  map[string]*models.ECDDBusinessProfile
	UserCaseFolders   map[string]*models.ECDDUserCaseManagementFolder
	mu                sync.RWMutex
}

var (
	instance *DataService
	once     sync.Once
)

// GetDataService returns singleton instance of DataService
func GetDataService() *DataService {
	once.Do(func() {
		instance = &DataService{
			ThresholdConfigs:  make(map[string]*models.ECDDThresholdConfig),
			UserStatuses:      make(map[string]*models.ECDDUserStatus),
			CaseFolders:       make(map[string]*models.ECDDCaseManagementFolder),
			MultiplierConfigs: make(map[string]*models.ECDDMultiplierConfig),
			BusinessProfiles:  make(map[string]*models.ECDDBusinessProfile),
			UserCaseFolders:   make(map[string]*models.ECDDUserCaseManagementFolder),
		}
	})
	return instance
}

// LoadAllData loads all JSON files into memory
func (ds *DataService) LoadAllData() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Load case folders first (referenced by other tables)
	if err := ds.loadCaseFolders(); err != nil {
		return fmt.Errorf("failed to load case folders: %w", err)
	}

	// Load threshold configs
	if err := ds.loadThresholdConfigs(); err != nil {
		return fmt.Errorf("failed to load threshold configs: %w", err)
	}

	// Load user statuses
	if err := ds.loadUserStatuses(); err != nil {
		return fmt.Errorf("failed to load user statuses: %w", err)
	}

	// Load multiplier configs
	if err := ds.loadMultiplierConfigs(); err != nil {
		return fmt.Errorf("failed to load multiplier configs: %w", err)
	}

	// Load business profiles
	if err := ds.loadBusinessProfiles(); err != nil {
		return fmt.Errorf("failed to load business profiles: %w", err)
	}

	// Load user-case folder assignments
	if err := ds.loadUserCaseFolders(); err != nil {
		return fmt.Errorf("failed to load user case folders: %w", err)
	}

	fmt.Println("✓ All data loaded successfully")
	fmt.Printf("  - Threshold Configs: %d\n", len(ds.ThresholdConfigs))
	fmt.Printf("  - User Statuses: %d\n", len(ds.UserStatuses))
	fmt.Printf("  - Case Folders: %d\n", len(ds.CaseFolders))
	fmt.Printf("  - Multiplier Configs: %d\n", len(ds.MultiplierConfigs))
	fmt.Printf("  - Business Profiles: %d\n", len(ds.BusinessProfiles))
	fmt.Printf("  - User-Folder Assignments: %d\n", len(ds.UserCaseFolders))

	return nil
}

func (ds *DataService) loadCaseFolders() error {
	data, err := os.ReadFile("data/case_folders.json")
	if err != nil {
		return err
	}

	var folders []models.ECDDCaseManagementFolder
	if err := json.Unmarshal(data, &folders); err != nil {
		return err
	}

	for i := range folders {
		ds.CaseFolders[folders[i].ECDDCaseManagementFolderPK] = &folders[i]
	}

	return nil
}

func (ds *DataService) loadThresholdConfigs() error {
	data, err := os.ReadFile("data/threshold_configs.json")
	if err != nil {
		return err
	}

	var configs []models.ECDDThresholdConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return err
	}

	for i := range configs {
		ds.ThresholdConfigs[configs[i].ECDDThresholdConfigPK] = &configs[i]
	}

	return nil
}

func (ds *DataService) loadUserStatuses() error {
	data, err := os.ReadFile("data/user_statuses.json")
	if err != nil {
		return err
	}

	var statuses []models.ECDDUserStatus
	if err := json.Unmarshal(data, &statuses); err != nil {
		return err
	}

	for i := range statuses {
		// Ensure language is always populated so frontend doesn't need to hardcode it
		if statuses[i].Language == 0 {
			// For the mock service we default to English (1) if not provided in JSON
			statuses[i].Language = 1
		}
		ds.UserStatuses[statuses[i].ECDDUserStatusPK] = &statuses[i]
	}

	return nil
}

func (ds *DataService) loadMultiplierConfigs() error {
	data, err := os.ReadFile("data/multiplier_configs.json")
	if err != nil {
		return err
	}

	var configs []models.ECDDMultiplierConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return err
	}

	for i := range configs {
		ds.MultiplierConfigs[configs[i].ECDDMultiplierConfigPK] = &configs[i]
	}

	return nil
}

func (ds *DataService) loadBusinessProfiles() error {
	data, err := os.ReadFile("data/business_profiles.json")
	if err != nil {
		return err
	}

	var profiles []models.ECDDBusinessProfile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return err
	}

	for i := range profiles {
		ds.BusinessProfiles[profiles[i].ECDDBusinessProfilePK] = &profiles[i]
	}

	return nil
}

func (ds *DataService) loadUserCaseFolders() error {
	data, err := os.ReadFile("data/user_case_folders.json")
	if err != nil {
		return err
	}

	var assignments []models.ECDDUserCaseManagementFolder
	if err := json.Unmarshal(data, &assignments); err != nil {
		return err
	}

	for i := range assignments {
		ds.UserCaseFolders[assignments[i].ECDDUserCaseManagementFolderPK] = &assignments[i]
	}

	return nil
}

// Helper methods for CRUD operations

// GetAllThresholdConfigs returns all threshold configs (thread-safe)
func (ds *DataService) GetAllThresholdConfigs() []*models.ECDDThresholdConfig {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	configs := make([]*models.ECDDThresholdConfig, 0, len(ds.ThresholdConfigs))
	for _, config := range ds.ThresholdConfigs {
		configs = append(configs, config)
	}
	return configs
}

// GetThresholdConfigByID returns a threshold config by ID
func (ds *DataService) GetThresholdConfigByID(id string) *models.ECDDThresholdConfig {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.ThresholdConfigs[id]
}

// CreateThresholdConfig adds a new threshold config
func (ds *DataService) CreateThresholdConfig(config *models.ECDDThresholdConfig) *models.ECDDThresholdConfig {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	config.ECDDThresholdConfigPK = uuid.New().String()
	config.LoggedAt = time.Now()
	ds.ThresholdConfigs[config.ECDDThresholdConfigPK] = config
	return config
}

// UpdateThresholdConfig updates an existing threshold config
func (ds *DataService) UpdateThresholdConfig(id string, config *models.ECDDThresholdConfig) *models.ECDDThresholdConfig {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.ThresholdConfigs[id]; exists {
		config.ECDDThresholdConfigPK = id
		config.LoggedAt = time.Now()
		ds.ThresholdConfigs[id] = config
		return config
	}
	return nil
}

// DeleteThresholdConfig soft deletes a threshold config
func (ds *DataService) DeleteThresholdConfig(id string) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if config, exists := ds.ThresholdConfigs[id]; exists {
		config.IsActive = false
		config.LoggedAt = time.Now()
		return true
	}
	return false
}

// GetAllUserStatuses returns all user statuses
func (ds *DataService) GetAllUserStatuses() []*models.ECDDUserStatus {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	statuses := make([]*models.ECDDUserStatus, 0, len(ds.UserStatuses))
	for _, status := range ds.UserStatuses {
		statuses = append(statuses, status)
	}
	return statuses
}

// GetUserStatusByID returns a user status by ID
func (ds *DataService) GetUserStatusByID(id string) *models.ECDDUserStatus {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.UserStatuses[id]
}

// CreateUserStatus adds a new user status
func (ds *DataService) CreateUserStatus(status *models.ECDDUserStatus) *models.ECDDUserStatus {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	status.ECDDUserStatusPK = uuid.New().String()
	status.LoggedAt = time.Now()
	ds.UserStatuses[status.ECDDUserStatusPK] = status
	return status
}

// UpdateUserStatus updates an existing user status
func (ds *DataService) UpdateUserStatus(id string, status *models.ECDDUserStatus) *models.ECDDUserStatus {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.UserStatuses[id]; exists {
		status.ECDDUserStatusPK = id
		status.LoggedAt = time.Now()
		ds.UserStatuses[id] = status
		return status
	}
	return nil
}

// GetAllCaseFolders returns all case folders
func (ds *DataService) GetAllCaseFolders() []*models.ECDDCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	folders := make([]*models.ECDDCaseManagementFolder, 0, len(ds.CaseFolders))
	for _, folder := range ds.CaseFolders {
		folders = append(folders, folder)
	}
	return folders
}

// GetCaseFolderByID returns a case folder by ID
func (ds *DataService) GetCaseFolderByID(id string) *models.ECDDCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.CaseFolders[id]
}

// CreateCaseFolder adds a new case folder
func (ds *DataService) CreateCaseFolder(folder *models.ECDDCaseManagementFolder) *models.ECDDCaseManagementFolder {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	folder.ECDDCaseManagementFolderPK = uuid.New().String()
	folder.LoggedAt = time.Now()
	ds.CaseFolders[folder.ECDDCaseManagementFolderPK] = folder
	return folder
}

// UpdateCaseFolder updates an existing case folder
func (ds *DataService) UpdateCaseFolder(id string, folder *models.ECDDCaseManagementFolder) *models.ECDDCaseManagementFolder {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.CaseFolders[id]; exists {
		folder.ECDDCaseManagementFolderPK = id
		folder.LoggedAt = time.Now()
		ds.CaseFolders[id] = folder
		return folder
	}
	return nil
}

// GetAllMultiplierConfigs returns all multiplier configs
func (ds *DataService) GetAllMultiplierConfigs() []*models.ECDDMultiplierConfig {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	configs := make([]*models.ECDDMultiplierConfig, 0, len(ds.MultiplierConfigs))
	for _, config := range ds.MultiplierConfigs {
		configs = append(configs, config)
	}
	return configs
}

// GetMultiplierConfigByID returns a multiplier config by ID
func (ds *DataService) GetMultiplierConfigByID(id string) *models.ECDDMultiplierConfig {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.MultiplierConfigs[id]
}

// CreateMultiplierConfig adds a new multiplier config
func (ds *DataService) CreateMultiplierConfig(config *models.ECDDMultiplierConfig) *models.ECDDMultiplierConfig {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	config.ECDDMultiplierConfigPK = uuid.New().String()
	config.LoggedAt = time.Now()
	ds.MultiplierConfigs[config.ECDDMultiplierConfigPK] = config
	return config
}

// UpdateMultiplierConfig updates an existing multiplier config
func (ds *DataService) UpdateMultiplierConfig(id string, config *models.ECDDMultiplierConfig) *models.ECDDMultiplierConfig {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.MultiplierConfigs[id]; exists {
		config.ECDDMultiplierConfigPK = id
		config.LoggedAt = time.Now()
		ds.MultiplierConfigs[id] = config
		return config
	}
	return nil
}

// GetAllBusinessProfiles returns all business profiles
func (ds *DataService) GetAllBusinessProfiles() []*models.ECDDBusinessProfile {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	profiles := make([]*models.ECDDBusinessProfile, 0, len(ds.BusinessProfiles))
	for _, profile := range ds.BusinessProfiles {
		profiles = append(profiles, profile)
	}
	return profiles
}

// GetBusinessProfileByID returns a business profile by ID
func (ds *DataService) GetBusinessProfileByID(id string) *models.ECDDBusinessProfile {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.BusinessProfiles[id]
}

// CreateBusinessProfile adds a new business profile
func (ds *DataService) CreateBusinessProfile(profile *models.ECDDBusinessProfile) *models.ECDDBusinessProfile {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	profile.ECDDBusinessProfilePK = uuid.New().String()
	profile.LoggedAt = time.Now()
	ds.BusinessProfiles[profile.ECDDBusinessProfilePK] = profile
	return profile
}

// UpdateBusinessProfile updates an existing business profile
func (ds *DataService) UpdateBusinessProfile(id string, profile *models.ECDDBusinessProfile) *models.ECDDBusinessProfile {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.BusinessProfiles[id]; exists {
		profile.ECDDBusinessProfilePK = id
		profile.LoggedAt = time.Now()
		ds.BusinessProfiles[id] = profile
		return profile
	}
	return nil
}

// GetAllUserCaseFolders returns all user-folder assignments
func (ds *DataService) GetAllUserCaseFolders() []*models.ECDDUserCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	assignments := make([]*models.ECDDUserCaseManagementFolder, 0, len(ds.UserCaseFolders))
	for _, assignment := range ds.UserCaseFolders {
		assignments = append(assignments, assignment)
	}
	return assignments
}

// GetUserCaseFolderByID returns a user-folder assignment by ID
func (ds *DataService) GetUserCaseFolderByID(id string) *models.ECDDUserCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.UserCaseFolders[id]
}

// GetUserCaseFoldersByUserStatusPK returns all folder assignments for a user
func (ds *DataService) GetUserCaseFoldersByUserStatusPK(userStatusPK string) []*models.ECDDUserCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var assignments []*models.ECDDUserCaseManagementFolder
	for _, assignment := range ds.UserCaseFolders {
		if assignment.UserStatusPK == userStatusPK {
			assignments = append(assignments, assignment)
		}
	}
	return assignments
}

// GetUserCaseFoldersByFolderPK returns all user assignments for a folder
func (ds *DataService) GetUserCaseFoldersByFolderPK(folderPK string) []*models.ECDDUserCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var assignments []*models.ECDDUserCaseManagementFolder
	for _, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK == folderPK {
			assignments = append(assignments, assignment)
		}
	}
	return assignments
}

// GetUsersByFolderPK returns all user statuses assigned to a specific folder
func (ds *DataService) GetUsersByFolderPK(folderPK string) []*models.ECDDUserStatus {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var users []*models.ECDDUserStatus
	for _, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK == folderPK {
			if user, exists := ds.UserStatuses[assignment.UserStatusPK]; exists {
				users = append(users, user)
			}
		}
	}
	return users
}

// GetFolderAssignmentIndex builds a map from folder PK to its assignment records
// in a single O(M) pass over all assignments. Use this to avoid repeated O(M)
// scans when computing stats for multiple folders.
func (ds *DataService) GetFolderAssignmentIndex() map[string][]*models.ECDDUserCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	index := make(map[string][]*models.ECDDUserCaseManagementFolder)
	for _, ucf := range ds.UserCaseFolders {
		index[ucf.FolderPK] = append(index[ucf.FolderPK], ucf)
	}
	return index
}

// GetUserStatusMap returns the internal user statuses map for O(1) lookups.
// The returned map must not be modified by callers.
func (ds *DataService) GetUserStatusMap() map[string]*models.ECDDUserStatus {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.UserStatuses
}

// GetFoldersByUserPK returns all case folders that a user is assigned to.
func (ds *DataService) GetFoldersByUserPK(userStatusPK string) []*models.ECDDCaseManagementFolder {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var folders []*models.ECDDCaseManagementFolder
	for _, assignment := range ds.UserCaseFolders {
		if assignment.UserStatusPK == userStatusPK {
			if folder, exists := ds.CaseFolders[assignment.FolderPK]; exists {
				folders = append(folders, folder)
			}
		}
	}
	return folders
}

// CreateUserCaseFolder creates a new user-folder assignment
func (ds *DataService) CreateUserCaseFolder(assignment *models.ECDDUserCaseManagementFolder) *models.ECDDUserCaseManagementFolder {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	assignment.ECDDUserCaseManagementFolderPK = uuid.New().String()
	assignment.LoggedAt = time.Now()
	ds.UserCaseFolders[assignment.ECDDUserCaseManagementFolderPK] = assignment
	return assignment
}

// BulkCreateUserCaseFolders creates multiple user-folder assignments in a single
// atomic operation. It skips duplicates (user already assigned to the folder) and
// returns the list of created assignments plus a list of user IDs that were skipped.
func (ds *DataService) BulkCreateUserCaseFolders(folderPK string, userStatusPKs []string, updatedBy string) ([]*models.ECDDUserCaseManagementFolder, []string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Build set of existing user-folder pairs for O(1) duplicate check
	existingUsers := make(map[string]bool)
	for _, a := range ds.UserCaseFolders {
		if a.FolderPK == folderPK {
			existingUsers[a.UserStatusPK] = true
		}
	}

	var created []*models.ECDDUserCaseManagementFolder
	var skipped []string

	now := time.Now()
	for _, userPK := range userStatusPKs {
		if existingUsers[userPK] {
			skipped = append(skipped, userPK)
			continue
		}

		assignment := &models.ECDDUserCaseManagementFolder{
			ECDDUserCaseManagementFolderPK: uuid.New().String(),
			FolderPK:                       folderPK,
			UserStatusPK:                   userPK,
			UpdatedBy:                      updatedBy,
			LoggedAt:                       now,
		}
		ds.UserCaseFolders[assignment.ECDDUserCaseManagementFolderPK] = assignment
		created = append(created, assignment)
	}

	return created, skipped
}

// DeleteUserCaseFolder removes a user-folder assignment
func (ds *DataService) DeleteUserCaseFolder(id string) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.UserCaseFolders[id]; exists {
		delete(ds.UserCaseFolders, id)
		return true
	}
	return false
}

// DeleteUserCaseFoldersByFolderPK removes all user-folder assignments for a given folder
func (ds *DataService) DeleteUserCaseFoldersByFolderPK(folderPK string) int {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	count := 0
	for key, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK == folderPK {
			delete(ds.UserCaseFolders, key)
			count++
		}
	}
	return count
}

// DeleteUserCaseFolderByFolderPKAndUserStatusPK removes a user-folder assignment
// identified by the combination of folder PK and user status PK.
func (ds *DataService) DeleteUserCaseFolderByFolderPKAndUserStatusPK(folderPK, userStatusPK string) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	for key, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK == folderPK && assignment.UserStatusPK == userStatusPK {
			delete(ds.UserCaseFolders, key)
			return true
		}
	}
	return false
}

// BulkDeleteUserCaseFoldersByFolderPKAndUserStatusPKs removes multiple user-folder
// assignments for a given folder. It returns the count of successfully deleted
// assignments and a slice of user status PKs that were not found in the folder.
func (ds *DataService) BulkDeleteUserCaseFoldersByFolderPKAndUserStatusPKs(folderPK string, userStatusPKs []string) (int, []string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	// Build a set of user status PKs to delete for O(1) lookups
	toDelete := make(map[string]bool, len(userStatusPKs))
	for _, uid := range userStatusPKs {
		toDelete[uid] = false // false = not yet found/deleted
	}

	// Single pass over assignments: find and delete matches
	for key, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK != folderPK {
			continue
		}
		if _, wanted := toDelete[assignment.UserStatusPK]; wanted {
			delete(ds.UserCaseFolders, key)
			toDelete[assignment.UserStatusPK] = true // mark as deleted
		}
	}

	// Tally results
	deletedCount := 0
	var failedIDs []string
	for _, uid := range userStatusPKs {
		if toDelete[uid] {
			deletedCount++
		} else {
			failedIDs = append(failedIDs, uid)
		}
	}

	return deletedCount, failedIDs
}

// DeleteMultiplierConfig soft deletes a multiplier config by setting is_active = false
func (ds *DataService) DeleteMultiplierConfig(id string) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if config, exists := ds.MultiplierConfigs[id]; exists {
		config.IsActive = false
		config.LoggedAt = time.Now()
		return true
	}
	return false
}

// DeleteBusinessProfile soft deletes a business profile by setting enabled = false
func (ds *DataService) DeleteBusinessProfile(id string) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if profile, exists := ds.BusinessProfiles[id]; exists {
		profile.Enabled = false
		profile.LoggedAt = time.Now()
		return true
	}
	return false
}

// DeleteCaseFolder hard deletes a folder and cascade-deletes all related user_case_folders
func (ds *DataService) DeleteCaseFolder(id string) (bool, int) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.CaseFolders[id]; !exists {
		return false, 0
	}

	// Cascade delete all user-folder assignments for this folder
	cascadeCount := 0
	for key, assignment := range ds.UserCaseFolders {
		if assignment.FolderPK == id {
			delete(ds.UserCaseFolders, key)
			cascadeCount++
		}
	}

	// Delete the folder itself
	delete(ds.CaseFolders, id)
	return true, cascadeCount
}

// DeleteUserStatus hard deletes a user and cascade-deletes all related user_case_folders
func (ds *DataService) DeleteUserStatus(id string) (bool, int) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.UserStatuses[id]; !exists {
		return false, 0
	}

	// Cascade delete all user-folder assignments for this user
	cascadeCount := 0
	for key, assignment := range ds.UserCaseFolders {
		if assignment.UserStatusPK == id {
			delete(ds.UserCaseFolders, key)
			cascadeCount++
		}
	}

	// Delete the user itself
	delete(ds.UserStatuses, id)
	return true, cascadeCount
}
