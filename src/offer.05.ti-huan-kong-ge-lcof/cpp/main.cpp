#include <iostream>
#include <string>

/*
{
  from:"
  https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/" ,
  description:"替换空格"
}
*/
using namespace std;

class Solution {
 public:
  string replaceSpace(string s) {
    string res;
    for (char &i : s) {
      if (i == ' ') {
        res += "%20";
      } else {
        res += i;
      }
    }
    return res;
  };
};

int main() {

  return 0;
}
