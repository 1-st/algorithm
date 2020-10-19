//
// Created by up0 on 2020/10/14.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805528788582400",
    reference:[],
    description:"A+B Format",
    solved:true,
}
*/

#include <iostream>
#include <cmath>
#include <iomanip>

int main() {
  using namespace std;
  int a, b;
  while (cin >> a >> b) {
    auto n = a + b;
    if (n < 0) {
      cout << "-";
      n = -n;
    }
    int div = 0;
    {
      unsigned int nn = n;
      for (; nn >= 1000;) {
        nn /= 1000;
        div++;
      }
    }
    cout << n / (int) (pow(1000, div--));
    for (; div > 0;) {
      cout << ','
           << setw(3) << setfill('0')
           << (n / (int) (pow(1000, div--))) % 1000;
    }
    if (n >= 1000) {
      cout << ','
           << setw(3) << setfill('0')
           << n % 1000;
    }
    cout << endl;
  }
}