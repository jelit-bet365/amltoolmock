const fs = require('fs');
const path = require('path');

// Region to Country ID mapping
const regions = {
  MALTA: { name: 'Malta', countryIds: [134, 135, 136] },
  USA: { name: 'USA', countryIds: [231, 232, 233] },
  AUSTRALIA: { name: 'Australia', countryIds: [13, 14, 15] },
  GIBRALTAR: { name: 'Gibraltar', countryIds: [83, 84, 85] }
};

// Names by region - expanded to support 30+ users per region
const names = {
  MALTA: [
    'Matthew Borg', 'Maria Vella', 'Paul Azzopardi', 'Grace Camilleri', 'Joseph Farrugia',
    'Sarah Mizzi', 'Daniel Zammit', 'Rebecca Galea', 'Anthony Bonnici', 'Lisa Schembri',
    'Mark Caruana', 'Emma Fenech', 'Christopher Attard', 'Nicole Scerri', 'David Cutajar',
    'Anna Barbara', 'James Formosa', 'Rachel Cassar', 'Peter Gatt', 'Joanne Mifsud',
    'Andrew Pace', 'Claudia Debono', 'Simon Grech', 'Michelle Baldacchino', 'Stefan Brincat',
    'Diane Spiteri', 'Kevin Bugeja', 'Tania Mercieca', 'Robert Abela', 'Christine Micallef'
  ],
  USA: [
    'Christopher Davis', 'Emily Rodriguez', 'Daniel Miller', 'William Jackson', 'Sophia Williams',
    'Michael Anderson', 'Olivia Martinez', 'James Taylor', 'Isabella Garcia', 'Robert Wilson',
    'Mia Thompson', 'John Moore', 'Charlotte Lee', 'David Brown', 'Amelia White',
    'Joseph Harris', 'Ava Clark', 'Benjamin Scott', 'Madison King', 'Alexander Wright',
    'Grace Adams', 'Samuel Green', 'Lily Baker', 'Henry Nelson', 'Chloe Hill',
    'Sebastian Rivera', 'Zoey Campbell', 'Jack Mitchell', 'Ella Turner', 'Owen Phillips'
  ],
  AUSTRALIA: [
    'Sophia Williams', 'Oliver Thompson', 'Charlotte Harris', 'Ava Mitchell', 'Liam O\'Brien',
    'Emma Johnston', 'Noah Murray', 'Mia Kelly', 'William Fraser', 'Isabella Campbell',
    'Jack Robertson', 'Sophie McKenzie', 'Lucas Stewart', 'Chloe Patterson', 'Mason Anderson',
    'Ella Hughes', 'Henry Russell', 'Amelia Watson', 'James Cooper', 'Harper Bailey',
    'Thomas Edwards', 'Evelyn Collins', 'Oscar Turner', 'Grace Morgan', 'Charlie Jenkins',
    'Lily Richardson', 'Max Sullivan', 'Ruby Howard', 'Archie Barnes', 'Zoe Foster'
  ],
  GIBRALTAR: [
    'George Caruana', 'Isabella Santos', 'Carlos Perez', 'Sofia Ramirez', 'Juan Martinez',
    'Elena Garcia', 'Miguel Lopez', 'Carmen Rodriguez', 'Antonio Fernandez', 'Laura Gonzalez',
    'Diego Sanchez', 'Maria Torres', 'Pablo Ruiz', 'Ana Morales', 'Francisco Jimenez',
    'Rosa Alvarez', 'Manuel Diaz', 'Lucia Hernandez', 'Pedro Romero', 'Isabel Navarro',
    'Javier Dominguez', 'Cristina Munoz', 'Alejandro Serrano', 'Patricia Ortiz', 'Fernando Molina',
    'Teresa Castillo', 'Raul Vargas', 'Beatriz Ramos', 'Alberto Mendez', 'Monica Guerrero'
  ]
};

// USA State IDs - expanded for more users
const usaStates = [5, 12, 36, 48, 6, 33, 48, 12, 5, 36, 6, 48, 33, 12, 5, 36, 48, 4, 8, 17, 21, 25, 29, 39, 42, 47, 51, 53, 55, 10];

// Simple UUID generator
function generateUUID() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    const r = Math.random() * 16 | 0;
    const v = c === 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

const users = [];
let userId = 100001;
let userIndex = 0;

// Minimum users per region
const USERS_PER_REGION = 30;

// Generate users for each region
Object.keys(regions).forEach(regionKey => {
  const region = regions[regionKey];
  const regionNames = names[regionKey];
  
  // Generate at least 30 users per region, distributed across country IDs
  for (let i = 0; i < USERS_PER_REGION; i++) {
    const countryId = region.countryIds[i % region.countryIds.length];
    const name = regionNames[i % regionNames.length];
    const isUSA = regionKey === 'USA';
    
    const user = {
      ecdd_user_status_pk: `u${String(userIndex + 1).padStart(3, '0')}-${regionKey.toLowerCase()}-${countryId}-${String((i % 10) + 1).padStart(3, '0')}`,
      user_id: userId++,
      user_name: name,
      country_id: countryId,
      state_id: isUSA ? usaStates[i % usaStates.length] : null,
      ecdd_status: (i % 5) + 1, // Cycle through statuses 1-5
      ecdd_threshold: 10000 + (i * 2000) + (Math.random() * 5000),
      ecdd_review_trigger: 3 + (i % 15),
      ecdd_suspension_due_date: (i % 5 === 4) ? `2026-02-${15 + (i % 10)}T00:00:00Z` : null,
      ecdd_multiplier: (i % 4 === 0) ? 0.50 : 1.00,
      ecdd_multiplier_rg_flag: i % 4 === 0,
      user_lt_engg_threshold_gbp: 2000 + (i * 500) + (Math.random() * 1000),
      user_lt_deposit_threshold_gbp: 4000 + (i * 1000) + (Math.random() * 2000),
      user_12month_drop_threshold_gbp: 1000 + (i * 250) + (Math.random() * 500),
      info_source: (i % 8) + 1,
      sign_off_status: (i % 3) + 1,
      date_last_ecdd_sign_off: (i % 3 === 0) ? `2026-0${(i % 2) + 1}-${String(10 + (i % 18)).padStart(2, '0')}T00:00:00Z` : null,
      ecdd_rg_review_status: (i % 4) + 1,
      date_last_ecdd_rg_sign_off: (i % 2 === 0) ? `2026-0${(i % 2) + 1}-${String(8 + (i % 20)).padStart(2, '0')}T00:00:00Z` : null,
      ecdd_report_status: (i % 4) + 1,
      ecdd_review_status: (i % 4) + 1,
      ecdd_document_status: (i % 4) + 1,
      ecdd_escalation_status: (i % 3) + 1,
      uar_status: (i % 3) + 1,
      logged_at: `2026-02-${String(1 + (i % 28)).padStart(2, '0')}T${String(9 + (i % 12)).padStart(2, '0')}:${String((i * 5) % 60).padStart(2, '0')}:00Z`,
      updated_by: `${regionKey.toLowerCase()}.compliance@company.com`
    };
    
    // Round numbers to 2 decimal places
    user.ecdd_threshold = Math.round(user.ecdd_threshold * 100) / 100;
    user.user_lt_engg_threshold_gbp = Math.round(user.user_lt_engg_threshold_gbp * 100) / 100;
    user.user_lt_deposit_threshold_gbp = Math.round(user.user_lt_deposit_threshold_gbp * 100) / 100;
    user.user_12month_drop_threshold_gbp = Math.round(user.user_12month_drop_threshold_gbp * 100) / 100;
    
    users.push(user);
    userIndex++;
  }
});

// Write users to file
const dataDir = path.join(__dirname, '..');
fs.writeFileSync(path.join(dataDir, 'user_statuses.json'), JSON.stringify(users, null, 2));
console.log(`Generated ${users.length} users across 4 regions`);
console.log(`Malta: ${users.filter(u => [134, 135, 136].includes(u.country_id)).length} users`);
console.log(`USA: ${users.filter(u => [231, 232, 233].includes(u.country_id)).length} users`);
console.log(`Australia: ${users.filter(u => [13, 14, 15].includes(u.country_id)).length} users`);
console.log(`Gibraltar: ${users.filter(u => [83, 84, 85].includes(u.country_id)).length} users`);

// Now generate user_case_folders mappings
// Read case folders
const caseFolders = JSON.parse(fs.readFileSync(path.join(dataDir, 'case_folders.json'), 'utf8'));
console.log(`\nLoaded ${caseFolders.length} case folders`);

// Generate user_case_folders - assign each user to 1-3 case folders
const userCaseFolders = [];
users.forEach((user, idx) => {
  // Determine how many folders this user belongs to (1-3)
  const numFolders = 1 + (idx % 3);
  
  // Assign to folders based on user index to ensure distribution
  const assignedFolderIndices = new Set();
  for (let f = 0; f < numFolders; f++) {
    const folderIndex = (idx + f * 3) % caseFolders.length;
    assignedFolderIndices.add(folderIndex);
  }
  
  assignedFolderIndices.forEach(folderIndex => {
    const folder = caseFolders[folderIndex];
    const userCaseFolder = {
      ecdd_user_case_management_folder_pk: generateUUID(),
      folder_id: folder.ecdd_case_management_folder_pk,
      user_id: user.ecdd_user_status_pk,
      logged_at: user.logged_at,
      updated_by: user.updated_by
    };
    userCaseFolders.push(userCaseFolder);
  });
});

// Write user_case_folders to file
fs.writeFileSync(path.join(dataDir, 'user_case_folders.json'), JSON.stringify(userCaseFolders, null, 2));
console.log(`\nGenerated ${userCaseFolders.length} user-case-folder mappings`);

// Show distribution of users per folder
console.log('\nUsers per case folder:');
caseFolders.forEach(folder => {
  const count = userCaseFolders.filter(ucf => ucf.folder_id === folder.ecdd_case_management_folder_pk).length;
  console.log(`  ${folder.folder_name}: ${count} users`);
});