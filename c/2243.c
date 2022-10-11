#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char *digitSum(char *s, int k) {
  int len = strlen(s);
  char *sTop = s;
  int add;
  char *s1 = malloc(k);
  char *st = s1;
  while (len > k) {
    for (int i = 0; i < len;) {
      add = 0;
      for (int p = 0; p < k && i < len; p++, i++) {
        add += s[i] - '0';
      }
      if (add == 0) {
        *st = 0;
        st++;
      } else {
        while (add > 0) {
          *st = add % 10;
          st++;
          add /= 10;
        }
      }

      st--;
      while (st >= s1) {
        *sTop = *st + '0';
        sTop++;
        st--;
      }
      st++;
    }
    len = sTop - s;
    sTop = s;
  }
  sTop[len] = 0;

  return s;
}

int main() {
  char s[] = "11111222223";
  digitSum(s, 3);
}