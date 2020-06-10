package authorization

default isAccessGranted = false

isAccessGranted {
  some i
  re_match(data.cache[i].regExp, input.path)
  data.cache[i].methods[_] == input.method
  data.cache[i].roles[_] == token.payload.roles[_]
}

token = {"payload": payload} {
  [header, payload, signature] := io.jwt.decode(input.token)
}

