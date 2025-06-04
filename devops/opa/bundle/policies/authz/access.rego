package authz.access

import data.authz.permissions

allow_read if {
    permissions.is_allowed(input.user.role, "read")
}