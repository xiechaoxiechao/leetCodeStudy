/*
 * @lc app=leetcode.cn id=933 lang=c
 *
 * [933] 最近的请求次数
 */

// @lc code=start

typedef struct {
  int *buf;
  int *top;
} RecentCounter;

RecentCounter *recentCounterCreate() {
  RecentCounter *re = malloc(sizeof(RecentCounter));
  re->buf = malloc(sizeof(int) * 3400);
  re->top = re->buf;
  return re;
}

int recentCounterPing(RecentCounter *obj, int t) {
  *(obj->top) = t;
  int *b = obj->buf;
  int k = t - 3000;
  while (*b < k) {
    b++;
  }
  int size = obj->top - b + 1;
  memmove(obj->buf, b, size * sizeof(int));
  obj->top = obj->buf + size;
  return size;
}

void recentCounterFree(RecentCounter *obj) {
  free(obj->buf);
  free(obj);
}

/**
 * Your RecentCounter struct will be instantiated and called as such:
 * RecentCounter* obj = recentCounterCreate();
 * int param_1 = recentCounterPing(obj, t);

 * recentCounterFree(obj);
*/
// @lc code=end
