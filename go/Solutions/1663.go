package Solutions

func getSmallestString(n int, k int) string {
    var ans =make([]byte,n)
    for i:=0;i<n;i++{
        ans[i]='a'
        k--
    }
    for i:=n-1;i>=0;i--{
        var sub int='z'-'a'
        if k>sub{
            ans[i]='z'
            k-=int(sub)
        }else{
            ans[i]='a'+byte(k)
            break
        }
    }
    return string(ans)
}