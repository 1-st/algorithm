//基本数组应用，矩阵置0
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <vector>

using namespace std;

class Solution {
public:
  void setZeroes(vector<vector<int>> &matrix) {
    int m = matrix.size();
    int n = matrix[0].size();
    char *arr_m = (char *)malloc(sizeof(char) * m);
    memset(arr_m, 0, m);
    char *arr_n = (char *)malloc(sizeof(char) * n);
    memset(arr_n, 0, n);
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (matrix[i][j] == 0) {
          arr_m[i] = 1;
          arr_n[j] = 1;
        }
      }
    }
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (arr_m[i] == 1 || arr_n[j] == 1) {
          matrix[i][j] = 0;
        }
      }
    }
  }
};
int main() {
  vector<vector<int>> m{{1, 2, 3}, {1, 0, 2}, {0, 5, 6}};
  Solution s;
  s.setZeroes(m);
  for(int i=0;i<m.size();i++){
    for(int j = 0;j<m[0].size();j++){
      printf("%d\t",m[i][j]);
    }
    printf("\n");
  }
  return 0;
}
