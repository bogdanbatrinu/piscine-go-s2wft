package main

func LastRUne (s string) rune {
  if len(s) == 0 {
    return 0
    }
  var last rune
  for _, r := range s {
    last = r
    }
  return last
  }
