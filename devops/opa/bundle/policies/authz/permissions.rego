package authz.permissions

is_allowed(role, action) if {
    some i
    data.authz.roles.roles[role].permissions[i] == action
}
