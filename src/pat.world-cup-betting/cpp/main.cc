//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805504927186944",
    reference:[],
    description:"World Cup Betting",
    solved:true,
}
*/

#include <iostream>
#include <algorithm>
#include <vector>
#include <cstdio>

using namespace std;

int getMax3(vector<double>::iterator it) {
  int idx = 0;
  double max = *it;
  for (int i = 0; i < 3; i++) {
    if (*it > max) {
      max = *it;
      idx = i;
    }
    it++;
  }
  return idx;
}

int main() {
  vector<double> B(9);
  char name[3] = {'W', 'T', 'L'};
  while (cin >> B[0]) {
    int i = 0;
    while (i < 8) {
      i++;
      cin >> B[i];
    }
    int i1 = getMax3(B.begin());
    int i2 = getMax3(B.begin() + 3);
    int i3 = getMax3(B.begin() + 6);
    cout << name[i1] << ' ' << name[i2] << ' ' << name[i3] << ' ';
    printf("%.2lf\n",(B[i1] * B[i2 + 3] * B[i3 + 6] * 0.65 - 1) * 2);
  }
}


