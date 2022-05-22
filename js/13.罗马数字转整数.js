/*
 * @lc app=leetcode.cn id=13 lang=javascript
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
 var romanToInt = function(s) {
    var arr=[];
    for(var i=0,b=s.length;i<b;i++){
        c(s[i]);
    } 
    function c(s){
        switch (s) {
            case 'M':
                arr.push(1000);
                break;
        case 'D':
               arr.push(500) 
            break;  
            case 'C':
                arr.push(100);
            break;  
            case 'L':
                arr.push(50)
            break;  
            case 'X':
                arr.push(10)
            break;  
            case 'V':
                arr.push(5)
            break;  
            case 'I':
                arr.push(1)
            break;  
            default:
                break;
        }
    }
    var res=0;
    for(i=0;i<b;i++){
        if(i===b-1){
            res+=arr[i];
            break;
        }
        if(arr[i]>=arr[i+1]){
            res+=arr[i];
        }else{
            res-=arr[i];
        }
    }

    return(res);
};

// @lc code=end

