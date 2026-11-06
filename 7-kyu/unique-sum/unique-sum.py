def unique_sum(lst: list[int]) -> int | None:
    if len(lst) == 0:
        return None
    mp: dict[int, int] = {}
    res: int = 0
    for n in lst:
        mp[n] = mp.get(n, 0) + 1
    
    for k in mp:
        res += k
        
    return res
        