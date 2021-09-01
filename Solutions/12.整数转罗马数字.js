/*
 * @lc app=leetcode.cn id=12 lang=javascript
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
/**
 * @param {number} num
 * @return {string}
 */
 var intToRoman = function(num) {
    var i=Math.floor(num/1000);
    var o=Math.floor((num-1000*i)/500);
    var p=Math.floor((num-i*1000-o*500)/100);
    var a=Math.floor((num-i*1000-o*500-p*100)/50);
    var s=Math.floor((num-i*1000-o*500-p*100-a*50)/10);
    var d=Math.floor((num-i*1000-o*500-p*100-a*50-s*10)/5);
    var f=Math.floor((num-i*1000-o*500-p*100-a*50-s*10-d*5));
    var res='';
    for(let r=0;r<i;r++){
        res+='M'
    }
    if(o*500+p*100===900){
        res+='CM';
    }else if(o*500+p*100===400){
        res+='CD';
    }else{
        if(o===1){
            res+='D';
        }
        for(let i=0;i<p;i++){
            res+='C';
        }
    }
    if((a*50+s*10)===90){
        res+='XC';
    }else if((a*50+s*10)===40){
        res+='XL';
    }else{
        if(a===1){
            res+='L';
        }
        for(let i=0;i<s;i++){
            res+='X';
        }
    }
    if(f+d*5===9){
        res+='IX';
    }else if(f+d*5===4){
        res+='IV';
    }else{
        if(d===1){
            res+='V';
        }
        for(let i=0;i<f;i++){
            res+='I';
        }
    }
    return(res);

};
// @lc code=end

