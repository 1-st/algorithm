//
// Created by up0 on 2020/10/26.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805516654460928",
    reference:[],
    description:"Sign In And Sign Out",
    solved:true,
}
*/



/*
 * tips: if max==0 output max max_start max_end
 *
 * */

#include <vector>
#include <iostream>



int main() {
  using namespace std;
  int sum, max, first, start, end, max_start, max_end;
  int M;
  while (cin >> M) {
    cin >> end;
    M--;
    sum = max = first = max_start = max_end = start = end;
    while (M--) {
      cin >> end;
      if (sum < 0) {
        sum = 0;
        start = end;
      }
      sum += end;
      if (sum > max) {
        max = sum;
        max_start = start;
        max_end = end;
      }
    }
    if (max < 0) {
      cout << 0 << " " << first << " " << end << "\n";
    } else {
      cout << max << " " << max_start << " " << max_end << "\n";
    }
  }
}

