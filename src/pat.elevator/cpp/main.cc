//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805511923286016",
    reference:[],
    description:"Elevator",
    solved:false,
}
*/

/*

 */
#include <iostream>

int main() {
  using namespace std;
  int N;
  while (cin >> N) {
    int M;
    int f = 0;
    int time = 0;
    while (N--) {
      cin >> M;
      if (M > f) {
        time += 6 * (M - f) + 5;
        f = M;
      } else {
        time += 4 * (f - M) + 5;
        f = M;
      }
    }
    cout<<time<<"\n";
  }
}

