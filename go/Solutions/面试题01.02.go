package Solutions

func CheckPermutation(s1 string, s2 string) bool {
    var mp=make(map[byte]int,len(s1))
    for i:=0;i<len(s1);i++{
        mp[s1[i]]++
    }
    for i:=0;i<len(s2);i++{
        if count,ok:=mp[s2[i]];!ok|| count <=0{
            return false
        }else{
            mp[s2[i]]--
        }
    }
    return true
}