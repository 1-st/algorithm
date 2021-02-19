//
// Created by up0 on 2020/12/4.
//
/*
{
    from:"https://www.nowcoder.com/pat/5/problem/4315",
    reference:[],
    description:"All Roads Lead To Rome",
    solved:false,
}
*/
#include <cstdio>
#include <map>
#include <vector>
#include <climits>
#include <string>
#include <iostream>
using namespace std;

const int N = 202;
int n; // number of city
int dist[N][N];
int happiness[N];
map<string, int> name2int{};
map<int, string> int2name{};

//dj
int cost[N];
bool vis[N];
vector<int> pre[N];
int cost2rom = INT_MAX;
int roads2rom = 0;
void dijkstra() {
  for (int i = 0; i < n; ++i) {
    cost[i] = INT_MAX;
    vis[i] = false;
  }
  cost[0] = 0;
  for (int j = 0; j < n; ++j) { //n.times
    int min = -1;
    for (int i = 0; i < n; ++i) {
      if ((!vis[i]) && ((min < 0) || (cost[i] < cost[min]))) {
        min = i;
      }
    }
    vis[min] = true;
    // relax
    for (int i = 0; i < n; ++i) {
      if ((!vis[i]) && (dist[min][i] < INT_MAX)/*edge exist*/) {
        int temp = cost[min] + dist[min][i];
        if (temp < cost[i]) {
          cost[i] = temp;
          // }extra
          if (int2name[i] == "ROM") {
            cost2rom = cost[i];
            roads2rom = 1;
          }
          pre[i].resize(1);
          pre[i][0] = min;
        } else if (temp == cost[i]) {
          if (int2name[i] == "ROM") {
            roads2rom++;
          }
          pre[i].push_back(min);
        }
      }
    }
  }
}

//dfs
int max_h = 0;
double max_avg = 0.0;
vector<int> answer;
void dfs(int next, int len, int h, vector<int> &path) {
  path.push_back(next);
  if (next == 0) {
    double avg = (double) h / len;
    if (h > max_h || ((h == max_h) && (avg > max_avg))) {
      max_h = h;
      max_avg = avg;
      answer = path;
      path.pop_back();
      return;
    }
  }
  for (int i : pre[next]) {
    dfs(i, len + 1, h + happiness[i], path);
  }
  path.pop_back();
}
int main() {
  // input
  int k; // road
  string start;
  cin >> n >> k >> start;
  name2int[start] = 0;
  int2name[0] = start;
  for (int i = 1; i < n; ++i) {
    string name;
    int h;
    cin >> name >> h;
    name2int[name] = i;
    int2name[i] = name;
    happiness[i] = h;
  }
  for (int i = 0; i < n; ++i) {
    for (int j = 0; j < n; ++j) {
      dist[i][j] = INT_MAX;
    }
  }
  while (k--) {
    string x, y;
    int z;
    cin >> x >> y >> z;
    dist[name2int[x]][name2int[y]] = dist[name2int[y]][name2int[x]] = z;
  }

  // processing
  dijkstra();
  vector<int> path; // for dfs
  dfs(name2int["ROM"], 0, happiness[name2int["ROM"]], path);

  // output
  printf("%d %d %d %d\n",roads2rom,cost2rom,max_h,(int)max_avg);
  for (int i = answer.size() - 1; i >= 0; i--) {
    if (i != answer.size() - 1) {
      printf("->");
    }
    cout << int2name[answer[i]];
  }
  cout << "\n";
  return 0;
}
