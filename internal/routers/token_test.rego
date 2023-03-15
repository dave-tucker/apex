package token

import future.keywords

mock_decode_verify("org-write-jwt", _) := [true, {"scope": "openid profile email write:organizations"}, {}]
mock_decode_verify("org-read-jwt", _) := [true, {"scope": "openid profile email read:organizations"}, {}]
mock_decode_verify("device-write-jwt", _) := [true, {"scope": "openid profile email write:devices"}, {}]
mock_decode_verify("device-read-jwt", _) := [true, {"scope": "openid profile email read:devices"}, {}]
mock_decode_verify("user-write-jwt", _) := [true, {"scope": "openid profile email write:users"}, {}]
mock_decode_verify("user-read-jwt", _) := [true, {"scope": "openid profile email read:users"}, {}]
mock_decode_verify("bad-jwt", _) := [false, {}, {}]

test_org_get_allowed if {
    allow
        with input.path as "/organizations"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "org-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_org_post_allowed if {
    allow
        with input.path as "/organizations"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "org-write-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_org_post_with_read_scope_denied if {
    not allow
        with input.path as "/organizations"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "org-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_org_get_anonymous_denied if {
     not allow
        with input.path as "/organizations"
        with input.method as "GET"
        with io.jwt.decode_verify as mock_decode_verify
}

test_org_get_bad_jwt_denied if {
     not allow
        with input.path as "/organizations"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with io.jwt.decode_verify as mock_decode_verify
}


test_device_get_allowed if {
    allow
        with input.path as "/devices"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "device-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_device_post_allowed if {
    allow
        with input.path as "/devices"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "device-write-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_device_post_with_read_scope_denied if {
    not allow
        with input.path as "/devices"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "device-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_device_get_anonymous_denied if {
     not allow
        with input.path as "/devices"
        with input.method as "GET"
        with io.jwt.decode_verify as mock_decode_verify
}

test_device_get_bad_jwt_denied if {
     not allow
        with input.path as "/devices"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "bad-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_user_get_allowed if {
    allow
        with input.path as "/users"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "user-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_user_post_allowed if {
    allow
        with input.path as "/users"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "user-write-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_user_post_with_read_scope_denied if {
    not allow
        with input.path as "/users"
        with input.method as "POST"
        with input.jwks as "my-cert"
        with input.access_token as "user-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_user_get_anonymous_denied if {
     not allow
        with input.path as "/users"
        with input.method as "GET"
        with io.jwt.decode_verify as mock_decode_verify
}

test_user_get_bad_jwt_denied if {
     not allow
        with input.path as "/users"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "bad-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}


test_get_fflags if {
     allow
        with input.path as "/fflags"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "user-read-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}

test_get_fflags_anonymous_denied if {
     not allow
        with input.path as "/fflags"
        with input.method as "GET"
        with io.jwt.decode_verify as mock_decode_verify
}

test_get_fflags_bad_jwt_denied if {
     not allow
        with input.path as "/fflags"
        with input.method as "GET"
        with input.jwks as "my-cert"
        with input.access_token as "bad-jwt"
        with io.jwt.decode_verify as mock_decode_verify
}
