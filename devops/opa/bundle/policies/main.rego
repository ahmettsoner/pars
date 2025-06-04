package main

default allow = false

allow if {
    input.method == "GET"
    data.authz.access.allow_read
}