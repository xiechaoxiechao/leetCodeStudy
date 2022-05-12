int minCount(int* coins, int coinsSize){
    int res=0;
    for(int i=0;i<coinsSize;i++){
        res+=(coins[i]-1)/2+1;
    }    
    return res;
}