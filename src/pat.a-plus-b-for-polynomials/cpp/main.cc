//
// Created by up0 on 2020/10/14.
//

/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805526272000000",
    reference:[],
    description:"A+B for polynomials",
    solved:true,
}
*/


#include <iostream>
#include <iomanip>
#include <map>
int main() {
  using namespace std;
  int K = 0;
  while (cin >> K) {
    std::map<int, double> dict;
    int a;
    double b;
    while (K--) {
      cin >> a >> b;
      dict[a] = b;
    }
    int KK = 0;
    cin >> KK;
    int zero = 0;
    while (KK--) {
      cin >> a >> b;
      dict[a] += b;
      if (dict[a] == 0) {
        zero++;
      }
    }
    cout << dict.size() - zero;
    for (auto p = dict.rbegin(); p != dict.rend(); p++) {
      if (p->second != 0) {
        cout
            << ' ' << p->first << ' '
            << fixed << setprecision(1) << p->second;
      }
    }
    cout << endl;
  }
}

