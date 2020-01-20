// Code generated by "scopegen"; DO NOT EDIT.
package branch

var Scopes = map[string]string{
	"https://auth.bnk.to/branch.read": "View branch data",
	"https://auth.bnk.to/branch.write": "Manage branch data",
}

var AuthScopes = map[string][]string{
	"/branch.BranchService/CreateBranch": []string{"https://auth.bnk.to/branch.write"},
	"/branch.BranchService/DeleteBranch": []string{"https://auth.bnk.to/branch.write"},
	"/branch.BranchService/GetBranch": []string{"https://auth.bnk.to/branch.read"},
	"/branch.BranchService/GetBranches": []string{"https://auth.bnk.to/branch.read"},
	"/branch.BranchService/UpdateBranch": []string{"https://auth.bnk.to/branch.write"},
}
