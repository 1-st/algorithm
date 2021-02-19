//
// Created by up0 on 2020/10/14.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805519074574336",
    reference:[],
    description:"Spell It Right",
    solved:true,
}
*/

#include <iostream>
#include <string>
#include <sstream>

int main() {
  using namespace std;
  string arr[10]{"zero","one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
  string s;
  while (cin >> s) {
    int sum = 0;
    for (auto i:s) {
      sum += i - '0';
    }
    stringstream ss;
    ss << sum;
    for (uint i = 0; i < ss.str().size(); i++) {
      if (i == ss.str().size() - 1) {
        cout<<arr[ss.str()[i]-'0'];
      } else {
        cout<<arr[ss.str()[i]-'0']<<' ';
      }
    }
    cout<<"\n";
  }
}
