/*
 * @lc app=leetcode.cn id=14 lang=javascript
 *
 * [14] 最长公共前缀
 */

// @lc code=start
/**
 * @param {string[]} strs
 * @return {string}
 */
/**
 * @param {string[]} strs
 * @return {string}
 */
 var longestCommonPrefix = function(strs) {
    var b=strs[0];
    var c;
    var d;
    var m;
    var n1=strs.length;
    var k;
    var sw=1;
    if(strs.length===0){
        return('');
    }
    for(var i=0,n=b.length;i<n;i++){
        d=b.slice(0,i+1);
        c=new RegExp(`^${d}`);
        for(m=0;m<n1;m++){
            k=strs[m].match(c);
            if(k===null){
                sw=0;
                break;
            }
        }
        if(sw===0){
            break;
        }
    }
    return(b.slice(0,i));


};
// @lc code=end

