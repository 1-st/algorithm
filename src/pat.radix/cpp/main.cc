//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805507225665536",
    reference:["https://blog.csdn.net/CV_Jason/article/details/80993283"],
    description:"Radix",
    solved:false,
}
*/

/*

 */


#include <iostream>
#include <string>
#include <cmath>
#include <algorithm>

using namespace std;

long long str2num(string str, long long radix) {
  long long sum = 0;
  int index = 0;
  int per_digit = 0;
  for (auto t = str.rbegin(); t != str.rend(); t++) {
    per_digit = isdigit(*t) ? *t - '0' : *t - 'a' + 10;
    sum += per_digit * pow(radix, index++);
  }
  return sum;
}

long long find_radix(string str, long long num) {
  long long result_radix = -1;
  char it = *max_element(str.begin(), str.end());
  long long low = (isdigit(it) ? it - '0' : it - 'a' + 10) + 1;
  long long high = max(low, num);
  while (low <= high) {
    long long mid = (low + high) / 2;
    long long temp = str2num(str, mid);
    if (temp < 0 || temp > num) {
      high = mid - 1;
    } else if (temp < num) {
      low = mid + 1;
    } else {
      result_radix = mid;
      break;
    }
  }
  return result_radix;

}

int main() {
  string N1;
  string N2;
  int tag;
  long long radix;
  while (cin >> N1 >> N2 >> tag >> radix) {
    long long known_num = (tag == 1 ? str2num(N1, radix) : str2num(N2, radix));
    long long result = find_radix((tag == 1 ? N2 : N1), known_num);
    if (result != -1) {
      cout << result << endl;
    } else {
      cout << "Impossible" << endl;
    }
  }
  return 0;
}