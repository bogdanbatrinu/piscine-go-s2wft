package main

fun DivMod(a int, b int, *div int, *mod int) {
  *div = a/b
  *mod = a%b
  }
