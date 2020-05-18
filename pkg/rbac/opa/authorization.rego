package authorization

default isAccessGranted = false

isAccessGranted {
  isPathAllowed
  isMethodAllowed
}

isPathAllowed {
  some roleID
  role := data.roles[roleID]
  roleID == token.payload.roles[_]
  path := role.entries[_].endpoints[_].path
  path == input.path
}

isMethodAllowed {
  some roleID
  role := data.roles[roleID]
  roleID == token.payload.roles[_]
  method := role.entries[_].endpoints[_].method
  method == input.method
}

token = {"payload": payload} {
  [header, payload, signature] := io.jwt.decode(input.token)
}

