using System.Collections.Generic;
/*
 * @lc app=leetcode.cn id=68 lang=csharp
 *
 * [68] 文本左右对齐
 */

// @lc code=start
public partial class Solutions {
    public IList<string> FullJustify(string[] words, int maxWidth) {
        var res=new List<string>();
        int len;
        var last=0;
        var spaceNum=0;
        int left;
        bool isFina=false;
        var tem=string.Empty;
        for(var i=0;i<words.Length;){
            len=0;
            while(true){
                len+=words[i].Length;
                if(len>maxWidth){
                    len-=(words[i].Length+1);
                    i--;
                    isFina=false;
                    break;
                }
                if(i<words.Length-1){
                    i++;
                    len++;
                }else{
                    isFina=true;
                    break;
                }
            }
            if(!isFina){
                if(i!=last){
                     spaceNum=(maxWidth-len)/(i-last);
                     left=(maxWidth-len)-spaceNum*(i-last);
                     while(last<i){
                         tem+=words[last];
                         if(left!=0){
                             tem+=new string(' ',spaceNum+2);
                             left--;
                         }else{
                            tem+=new string(' ',spaceNum+1);
                         }
                         last++;
                         
                     }
                     tem+=words[last];
                     last++;
                     i++;
                     res.Add(tem);
                     tem=string.Empty;

                }else{
                    tem+=words[i];
                    i++;
                    last++;
                    tem+=new string(' ',maxWidth-tem.Length);
                    res.Add(tem);
                    tem=string.Empty;
                }
               
            }else{
                while(last!=i){
                    tem+=words[last]+" ";
                    last++;
                }
                tem+=words[last];
                tem+=new string(' ',maxWidth-tem.Length);
                res.Add(tem);
                i++;
            }

            
             
        }
        return res;

    }
}
// @lc code=end

