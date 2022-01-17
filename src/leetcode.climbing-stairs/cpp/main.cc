//爬楼梯问题，解决方法是动态规划，f(x) = f(x-1) + f(x-2)
#include <cstdio>

class Solution {
public:
  int climbStairs(int n) {
    if (n <= 2) {
      return n;
    }
    int ret_1 = 1;
    int ret_2 = 2;
    int ret = 0;
    for (int i = 2; i < n; i++) {
      ret = ret_1 + ret_2;
      ret_1 = ret_2;
      ret_2 = ret;
    }
    return ret;
  }
};
int main() {
  Solution S{};
  printf("%d\n", S.climbStairs(3));
  printf("%d\n", S.climbStairs(4));
  return 0;
}
