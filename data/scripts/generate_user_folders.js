const fs = require('fs');

// Read the users to get their UUIDs
const users = JSON.parse(fs.readFileSync('data/user_statuses.json', 'utf8'));

// Folder UUIDs from case_folders.json
const folders = [
  { id: 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', name: 'High Risk Customers' },
  { id: 'b2c3d4e5-f6a7-5b6c-9d0e-1f2a3b4c5d6e', name: 'Pending Review' },
  { id: 'c3d4e5f6-a7b8-6c7d-0e1f-2a3b4c5d6e7f', name: 'Suspended Accounts' },
  { id: 'd4e5f6a7-b8c9-7d8e-1f2a-3b4c5d6e7f8a', name: 'VIP Customers' },
  { id: 'e5f6a7b8-c9d0-8e9f-2a3b-4c5d6e7f8a9b', name: 'New Accounts Under Review' },
  { id: 'f6a7b8c9-d0e1-9f0a-3b4c-5d6e7f8a9b0c', name: 'Medium Risk Profile' },
  { id: 'a7b8c9d0-e1f2-0a1b-4c5d-6e7f8a9b0c1d', name: 'Escalated Cases' },
  { id: 'b8c9d0e1-f2a3-1b2c-5d6e-7f8a9b0c1d2e', name: 'Document Verification Queue' },
  { id: 'c9d0e1f2-a3b4-2c3d-6e7f-8a9b0c1d2e3f', name: 'Responsible Gambling Concerns' },
  { id: 'd0e1f2a3-b4c5-3d4e-7f8a-9b0c1d2e3f4a', name: 'Threshold Breach Monitoring' }
];

const userFolders = [];
let assignmentIndex = 0;

// Assign users to folders - at least 3 per folder
folders.forEach((folder, folderIndex) => {
  // Determine how many users to assign (between 3 and 10)
  const numUsers = 3 + (folderIndex % 8);
  
  for (let i = 0; i < numUsers; i++) {
    const userIndex = (folderIndex * 7 + i) % users.length;
    const user = users[userIndex];
    
    const assignment = {
      ecdd_user_case_management_folder_pk: `uf${String(assignmentIndex + 1).padStart(3, '0')}-${folder.id.substring(0, 8)}`,
      folder_pk: folder.id,
      user_status_pk: user.ecdd_user_status_pk,
      created_at: `2026-0${(folderIndex % 2) + 1}-${String(10 + i).padStart(2, '0')}T${String(9 + (i % 12)).padStart(2, '0')}:00:00Z`,
      updated_at: `2026-0${(folderIndex % 2) + 1}-${String(10 + i).padStart(2, '0')}T${String(9 + (i % 12)).padStart(2, '0')}:00:00Z`,
      updated_by: `compliance@company.com`
    };
    
    userFolders.push(assignment);
    assignmentIndex++;
  }
});

// Write to file
fs.writeFileSync('data/user_case_folders.json', JSON.stringify(userFolders, null, 2));
console.log(`Generated ${userFolders.length} user-folder assignments`);
console.log(`Folders: ${folders.length}`);
console.log(`Average users per folder: ${(userFolders.length / folders.length).toFixed(1)}`);

// Show distribution
folders.forEach(folder => {
  const count = userFolders.filter(uf => uf.folder_pk === folder.id).length;
  console.log(`  ${folder.name}: ${count} users`);
});