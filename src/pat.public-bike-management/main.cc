//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://www.nowcoder.com/pat/5/problem/4324",
    reference:[],
    description:"Public Bike Management",
    solved:true,
}
*/

#include <cstdio>
#include <algorithm>
#include <vector>
#include <climits>

using namespace std;

const int N = 502;
int n; // number of city

int a[N][N];
int b[N];

//dj
int cost[N];
bool vis[N];
vector<int> pre[N];

vector<int> answer;
int best_in, best_out, cap;

void dijkstra() {
  for (int i = 0; i < n; ++i) {
    cost[i] = INT_MAX;
    vis[i] = false;
  }
  cost[0] = 0;
  for (int j = 0; j < n; ++j) {
    int min = -1;
    for (int i = 0; i < n; ++i) {
      if ((!vis[i]) && ((min < 0) || (cost[i] < cost[min]))) {
        min = i;
      }
    }
    vis[min] = true;
    for (int i = 0; i < n; ++i) {
      // relax
      if ((!vis[i]) && (a[min][i] < INT_MAX)/*edge exist*/) {
        int temp = cost[min] + a[min][i];
        if (temp < cost[i]) {
          cost[i] = temp;
          // }extra
          pre[i].resize(1);
          pre[i][0] = min;
        } else if (temp == cost[i]) {
          pre[i].push_back(min);
        }
      }
    }
  }
}

void dfs(int now, int in, int out, vector<int> &path) {
  path.push_back(now);
  if (now == 0) {
    if ((in < best_in) || ((in == best_in) && (out < best_out))) {
      best_in = in;
      best_out = out;
      answer = path;
    }
    path.pop_back();
    return;
  }
  if (b[now] >= cap / 2) {
    in -= (b[now] - cap / 2);
    if (in < 0) {
      out -= in;
      in = 0;
    }
  } else {
    in += cap / 2 - b[now];
  }
  for (int i : pre[now]) {
    dfs(i, in, out, path);
  }
  path.pop_back();
}

int main() {
  // input
  int p, to;
  scanf("%d%d%d%d", &cap, &n, &to, &p);
  ++n;
  for (int i = 1; i < n; ++i) {
    scanf("%d", b + i);
  }
  for (int i = 0; i < n; ++i) {
    for (int j = 0; j < n; ++j) {
      a[i][j] = INT_MAX;
    }
  }
  while (p--) {
    int x, y, z;
    scanf("%d%d%d", &x, &y, &z);
    a[x][y] = a[y][x] = z;
  }

  // processing
  dijkstra();
  best_in = best_out = INT_MAX;
  vector<int> path; // for dfs
  dfs(to, 0, 0, path);

  // output
  printf("%d ", best_in);
  for (int i = answer.size() - 1; i >= 0; --i) {
    if (i != answer.size() - 1) {
      printf("->");
    }
    printf("%d", answer[i]);
  }
  printf(" %d\n", best_out);
  return 0;
}
