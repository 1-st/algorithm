//
// Created by up0 on 2020/10/14.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805521431773184",
    reference:[],
    description:"A+B Format",
    solved:true,
}
*/

#include <vector>
#include <iostream>

using namespace std;

int main() {

  int N/*node sum*/, M/*non-leaf node*/;
  while (cin >> N >> M) {
    vector<vector<int>> pool(N + 1);
    int id, K;
    while (M--) {
      cin >> id >> K;
      int ids;
      while (K--) {
        cin >> ids;
        pool[id].push_back(ids);
      }
    }//input
    vector<int> save;
    save.push_back(1); //01

    if(N>1){
      cout << "0";
    }else{  //warning
      cout<< "1";
    }

    //bfs
    while (!save.empty()) {
      vector<int> next;
      for (auto &p:save) {
        next.insert(next.end(), pool[p].begin(), pool[p].end());
      }
      save.clear();
      int n = 0;
      for (auto &p:next) {
        if (pool[p].empty()) {
          n++;
        }
        save.push_back(p);
      }
      if (!save.empty()) {
        cout << ' ' << n;
      }
    }
    cout << '\n';
  }
}
