//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805509540921344",
    reference:[],
    description:"Product of Polynomials",
    solved:true,
}
*/

/*
 * tips

注意0号测试点，系数为0时结果不输出，并且系数为0也不算在总个数里面
测试点0未通过时，想到了是系数为0的问题，但是只是把coef != 0不输出，而忘了修改ans.size()，结果还是错误~因此找了不少时间 兜兜转转还是这个错误 只是没有考虑周到

 */


#include <iostream>
#include <cstring>
#include <cstdio>

#define Len 1001

int main() {
  using namespace std;
  double A[Len], B[Len], C[2*Len-1];
  bool hasA[Len], hasB[Len];
  memset(hasA, false, sizeof(bool) * Len);
  memset(hasB, false, sizeof(bool) * Len);
  memset(C, 0.0, sizeof(0.0) * (Len*2-1));
  int K;
  bool isA = true;
  double *p = A;
  bool *hasP = hasA;
  while (cin >> K) {
    if (isA) {
      p = A;
      hasP = hasA;
      isA = false;
    } else {
      p = B;
      hasP = hasB;
      isA = true;
    }
    int idx;
    double v;
    while (K--) {
      cin >> idx >> v;
      p[idx] = v;
      hasP[idx] = true;
    }
    if (isA) {
      //calculate
      for (int i = 0; i < Len; i++) {
        if (hasA[i]) {
          for (int j = 0; j < Len; j++) {
            if (hasB[j]) {
              C[i + j] += A[i] * B[j];
            }
          }
        }
      }
      //output
      int count = 0;
      int last = 0;
      for (int i=Len*2-1-1;i>=0;i--) {
        if (C[i]!=0.0) {
          count++;
          last = i;
        }
      }
      cout<<count<<' ';
      for (int i=Len*2-1-1;i>=0;i--) {
          if(C[i]!=0.0){
            if(i!=last){
              printf("%d %.1lf ",i,C[i]);
            }else{
              printf("%d %.1lf\n",i,C[i]);
            }
          }
      }
      //initialize
      memset(hasA, false, sizeof(bool) * Len);
      memset(hasB, false, sizeof(bool) * Len);
      memset(A,0.0,sizeof(0.0)*Len);
      memset(B,0.0,sizeof(0.0)*Len);
      memset(C,0.0,sizeof(0.0)*(Len*2-1));
    }
  }
}

