// Code generated by go generate; DO NOT EDIT.

package migration

var contractsDataSQL = `
INSERT INTO "1_contracts" (id, name, value, token_id, conditions, app_id, ecosystem)
VALUES
	(next_id('1_contracts'), 'AdminCondition', '// This contract is used to set "admin" rights.
// Usually the "admin" role is used for this.
// The role ID is written to the ecosystem parameter and can be changed.
// The contract requests the role ID from the ecosystem parameter and the contract checks the rights.

contract AdminCondition {
    conditions {
        if EcosysParam("founder_account") == $key_id {
            return
        }

        var role_id_param string
        role_id_param = EcosysParam("role_admin")
        if Size(role_id_param) == 0 {
            warning "Sorry, you do not have access to this action."
        }

        var role_id int
        role_id = Int(role_id_param)
        
        if !RoleAccess(role_id) {
            warning "Sorry, you do not have access to this action."
        }      
    }
}
', '{{.Ecosystem}}', 'ContractConditions("MainCondition")', '{{.AppID}}', '{{.Ecosystem}}'),
	(next_id('1_contracts'), 'DeveloperCondition', '// This contract is used to set "developer" rights.
// Usually the "developer" role is used for this.
// The role ID is written to the ecosystem parameter and can be changed.
// The contract requests the role ID from the ecosystem parameter and the contract checks the rights.

contract DeveloperCondition {
    conditions {
        // check for Founder
        if EcosysParam("founder_account") == AddressToId($account_id) {
            return
        }

        // check for Developer role
        var app_id int role_id string
        app_id = Int(DBFind("@1applications").Where({"ecosystem": $ecosystem_id, "name": "System"}).One("id"))
        role_id = AppParam(app_id, "role_developer", $ecosystem_id)

        if Size(role_id) == 0 {
            warning Sprintf(LangRes("@1x_not_access_action"),"DeveloperCondition")
        }
        if !RoleAccess(Int(role_id)) {
            warning Sprintf(LangRes("@1x_not_access_action"),"DeveloperCondition")
        }
    }
}
', '{{.Ecosystem}}', 'ContractConditions("MainCondition")', '{{.AppID}}', '{{.Ecosystem}}'),
	(next_id('1_contracts'), 'MainCondition', 'contract MainCondition {
	conditions {
		if EcosysParam("founder_account")!=$key_id
		{
			warning "Sorry, you do not have access to this action."
		}
	}
}
