package Solutions

func toHex(num int) string {
    if num<0{
        num=1<<32+num
    }
    return fmt.Sprintf("%x",num)
}