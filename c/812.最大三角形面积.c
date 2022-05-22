/*
 * @lc app=leetcode.cn id=812 lang=c
 *
 * [812] 最大三角形面积
 */

// @lc code=start

double area(int x1,int x2,int x3,int y1,int y2,int y3);
double largestTriangleArea(int** points, int pointsSize, int* pointsColSize){
    double max=0;
    for (int i=0;i<pointsSize;i++){
        for(int j=i+1;j<pointsSize;j++){
            for (int k=j+1;k<pointsSize;k++){
                double tem=area(points[i][0],points[j][0],points[k][0],points[i][1],points[j][1],points[k][1]);
                max=max<tem?tem:max;
            }
        }
    }
    return max;
}
double area(int x1,int x2,int x3,int y1,int y2,int y3){
    return 0.5*(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2);
}
// @lc code=end

