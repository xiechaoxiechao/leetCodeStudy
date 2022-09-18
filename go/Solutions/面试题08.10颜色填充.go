package Solutions 
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc]==newColor{
        return image
    }
    temColor:=image[sr][sc]
    image[sr][sc] = newColor
	if sr-1 >= 0 && image[sr-1][sc] ==  temColor {
		floodFill(image, sr-1, sc, newColor)
	}
	if sr+1 < len(image) && image[sr+1][sc] ==  temColor {
		floodFill(image, sr+1, sc, newColor)
	}
	if sc-1 >= 0 && image[sr][sc-1] ==  temColor {
		floodFill(image, sr, sc-1, newColor)
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] ==  temColor {
		floodFill(image, sr, sc+1, newColor)
	}
	
	return image
}
