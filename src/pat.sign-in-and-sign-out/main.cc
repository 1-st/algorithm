//
// Created by up0 on 2020/10/26.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805516654460928",
    reference:[],
    description:"Sign In And Sign Out",
    solved:true,
}
*/

#include <vector>
#include <iostream>
#include <string>
#include <cstdio>

template<typename T>
bool compare3(T a,T b,T c,T aa,T bb,T cc){
  if(a>aa){
    return true;
  }
  if(b>bb){
    return true;
  }
  return c > cc;
}

int main() {
  using namespace std;
  int M;
  string open_id;
  string close_id;
  int open_h,open_m,open_s;
  int close_h,close_m,close_s;
  while(cin>>M){
    M--;
    cin >> open_id;
    close_id = open_id;
    scanf("%d:%d:%d",&open_h,&open_m,&open_s);
    scanf("%d:%d:%d",&close_h,&close_m,&close_s);
    string id;
    int h,m,s;
    int hh,mm,ss;
    while(M--){
      cin>>id;
      scanf("%d:%d:%d",&h,&m,&s);
      scanf("%d:%d:%d",&hh,&mm,&ss);
      if(!compare3(h,m,s,open_h,open_m,open_s)){ //h < open_h
        open_h = h;open_m = m;open_s = s;
        open_id = id;
      }
      if(compare3(hh,mm,ss,close_h,close_m,close_s)){
        close_h = hh;close_m = mm;close_s = ss;
        close_id = id;
      }
    }
    cout<<open_id<<" "<<close_id<<"\n";
  }
}

