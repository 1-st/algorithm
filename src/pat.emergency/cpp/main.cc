//
// Created by up0 on 2020/10/14.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805523835109376",
    reference:[],
    description:"emergency",
    solved:true,
}
*/

#include <iostream>
#include <vector>
#include <map>
#include <functional>

int main() {
  using namespace std;
  int city, road, cur, target;
  while (cin >> city >> road >> cur >> target) {
    vector<int> man(city);
    vector<map<int, int>> edge(city);
    for (int i = 0; i < city; i++) { //read rescue team
      cin >> man[i];
    }
    while (road--) { //read road
      int src, dst, len;
      cin >> src >> dst >> len;
      edge[src][dst] = len;
      edge[dst][src] = len;
    }
    vector<bool> visited(city, false);

    int shortest = INT32_MAX;
    int shortest_count = 0;
    int max_man_count = 0;

    function<void(int, int, int, int)> dfs = [&](int src, int target, int len, int man_count) {
      if (len > shortest) { //cut
        return;
      }
      if (src == target) { //return
        if (len < shortest) {
          shortest = len;
          shortest_count = 1;
          max_man_count = 0;
        } else if (len == shortest) {
          shortest_count++;
        }
        if (man_count > max_man_count) {
          max_man_count = man_count;
        }
        return;
      }
      for (auto &i:edge[src]) {
        if (!visited[i.first]) {
          visited[i.first] = true;
          dfs(i.first, target, len + i.second, man_count + man[i.first]);
          visited[i.first] = false;
        }
      }
    };

    dfs(cur, target, 0, man[cur]);

    cout << shortest_count << ' ' << max_man_count << endl;
  }
}