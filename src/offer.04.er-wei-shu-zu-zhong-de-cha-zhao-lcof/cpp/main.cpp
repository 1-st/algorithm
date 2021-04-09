#include <iostream>
#include <vector>

/*
{
  from:"https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/" ,
  description:"二维数组中的查找"
}
*/
using namespace std;

class Solution {
 public:
  bool findNumberIn2DArray(vector<vector<int>> &matrix, int target) {
    if (matrix.size() == 0 || matrix[0].size() == 0) {
      return false;
    }
    int rows = matrix.size(), columns = matrix[0].size();
    int row = 0, column = columns - 1;
    while (row < rows && column >= 0) {
      int num = matrix[row][column];
      if (num == target) {
        return true;
      } else if (num > target) {
        column--;
      } else {
        row++;
      }
    }
    return false;
  }
};

int main() {
  vector<vector<int>> matrix = {{-5}};
  Solution s;
  cout << s.findNumberIn2DArray(matrix, -1);
  return 0;
}
