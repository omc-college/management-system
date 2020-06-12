package authorization

default isAccessGranted = false

isAccessGranted {
  some i
  re_match(input.cache.rules[i].pathRegExp, input.path)
  input.cache.rules[i].methods[j].name == input.method
  input.cache.rules[i].methods[j].roles[_] == token.payload.roles[_]
}

token = {"payload": payload} {
  [header, payload, signature] := io.jwt.decode(input.token)
}

