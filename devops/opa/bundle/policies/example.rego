package example

default allow = false

allow if {
  input.user == "alice"
  input.action == "read"
}

allow if {
  input.user == "bob"
  input.action == "write"
}

message = sprintf("User %s is allowed to perform %s: %v", [input.user, input.action, allow])
