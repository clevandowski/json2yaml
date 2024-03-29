package main

import (
  "testing"
)

func TestJson2YamlEmpty(t *testing.T) {
  json := ``

  want := ``

  got, _ := json2yaml(json)

  if want != got {
    t.Fatalf("json2yaml(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestJson2YamlBasicMap(t *testing.T) {
  json := `{ "coincoin": 5 }
`

  want := `coincoin: 5
`

  got, _ := json2yaml(json)

  if want != got {
    t.Fatalf("json2yaml(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestJson2YamlBasicArray(t *testing.T) {
  json := `[ "coincoin", "cuicui", "coucou" ]
`

  want := `- coincoin
- cuicui
- coucou
`

  got, _ := json2yaml(json)

  if want != got {
    t.Fatalf("json2yaml(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestJson2YamlAdvanced(t *testing.T) {
  json := `{
    "a": "Easy!",
    "b": {
      "c": 2,
      "d": true,
      "my.super/key.test": [
        "une\\vilaine` + "`string`" + `",
        4, {
          "another.key.too": {
            "x": 1,
            "y": 5,
            "z": 6.666666666666666666666666
          }
        }
      ],
      "e": [
        true,
        false
      ],
      "f": "coincoin\n  cuicui\\npouetpouet\nje\\mets\\des\\antislash\n",
      "g": 3.1415,
      "h": 3.141592654,
      "i": "true"
    }
  }  
`

  want := `a: Easy!
b:
    c: 2
    d: true
    e:
        - true
        - false
    f: |
        coincoin
          cuicui\npouetpouet
        je\mets\des\antislash
    g: 3.1415
    h: 3.141592654
    i: "true"
    my.super/key.test:
        - une\vilaine` + "`string`" + `
        - 4
        - another.key.too:
            x: 1
            "y": 5
            z: 6.666666666666667
`

  got, _ := json2yaml(json)

  if want != got {
    t.Fatalf("json2yaml(): want:\n%v\n, got:\n%v", want, got)
  }
}
