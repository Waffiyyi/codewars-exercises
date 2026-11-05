def hex_to_dec(s):
    mp: dict[str,int] = {"a": 10, "b" : 11, "c": 12, "d": 13, 
                     "e" : 14, "f" : 15}
    res: int = 0
    j=0
    for i in range(len(s) -1, -1, -1):
        ch: char = s[i]
        if ch.lower() in mp:
            num: int = mp[ch.lower()]
            res += num * (16**j)
        else :
            res +=int(ch) * (16**(j))
        j+=1
            
    return res
        