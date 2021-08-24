using System.Threading.Tasks;
/*
 * @lc app=leetcode.cn id=59 lang=csharp
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
public partial class Solution {
    public int[][] GenerateMatrix(int n) {
        var res=new int[n][];
        for(int i=0;i<n;i++){
            res[i]=new int[n];
        }
        var length=new int[n/2+1];
         if (length.Length>0){
            length[0]=0;
        }
        for(int i=1;i<length.Length;i++){
            length[i]=(n-2*i+1)*4+length[i-1];
        }
        var layer=0;
        for(var i=0;i<n;++i){
            for(var j=0;j<n;++j){
                layer=(i<j?i:j)<(i>j?n-i-1:n-j-1)?(i<j?i:j):(i>j?n-i-1:n-j-1);
                if(i==layer||j==n-1-layer){
                    res[i][j]=length[layer]+i+j-2*layer+1;
                }else{
                     res[i][j]=length[layer+1]-(i+j-2*layer)+1;
                     
                }
                
            }

        }

        return res;
    }
}
// @lc code=end

