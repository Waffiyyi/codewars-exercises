  }
​
​
  for i := 0; i < len(sl); i++ {
    temp := sl[i]
    prime := 2
    for temp > 1 {
      if temp%prime == 0 {
        mp[prime]++
      } else {
        prime = findNextPrime(prime)
        continue
      }
      temp = temp / prime
    }
    prime = 2
  }
  var res strings.Builder
  first := true
​
  for _, v := range sl {
    elem, ok := mp[v]
    if !ok {
        continue
    }
​
    if !first {
        res.WriteString(" * ")
    }
​
    if elem != 1 {
        res.WriteString(fmt.Sprintf("%d^%d", v, elem))
    } else {
        res.WriteString(fmt.Sprintf("%d", v))
    }
​
    first = false
}
  return res.String()
}