//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805500414115840",
    reference:[],
    description:"Battle Over Cities",
    solved:true,
}
*/

/*
 *tips:使用并查集最为简单
 *
 * */

#include<iostream>
using namespace std;

void make(int *father, int num) {
  int i;
  for (i = 0; i <= num; i++)
    father[i] = i;
}

int find(int *father, int x) {
  if (x != father[x]) father[x] = find(father, father[x]);
  return father[x];
}

void join(int *father, int x, int y) {
  x = find(father, x);
  y = find(father, y);
  if (x == y) return;
  else {
    father[y] = x;
  }
}

int main() {
  int n, m, k;
  int father[1001];
  int road[500001][2];
  while (cin >> n >> m >> k) {
    for (int i = 0; i < m; i++) {
      cin >> road[i][0] >> road[i][1];
    }
    for (int i = 0; i < k; i++) {
      make(father, n);
      int del;
      cin >> del;
      for (int j = 0; j < m; j++) {
        int x = road[j][0], y = road[j][1];
        if (del != x && del != y) {
          join(father, road[j][0], road[j][1]);
        }
      }
      int num = 0;
      for (int j = 1; j <= n; j++) {
        if (j != del && father[j] == j) { num++; };
      }
      cout << num - 1 << endl;
    }
  }
  return 0;
}


//#include <iostream>
//#include <vector>
//using namespace std;
//
//int find(vector<int> &father, int a) {
//  if (father[a] != a) {
//    father[a] = find(father, a);
//  }
//  return father[a];
//}
//
//void join(vector<int> &father, int a, int b) {
//  int x = find(father, a);
//  int y = find(father, b);
//  father[x] = y;
//}
//
//void clear(vector<int> &father) {
//  for (int i = 0; i < father.size(); i++) {
//    father[i] = i;
//  }
//}
//
//
//int main() {
//  vector<pair<int, int>> roads;
//  vector<int> father;
//  int n_city, n_road, n_del;
//  cin >> n_city >> n_road >> n_del;
//  father.resize(n_city);
//  for (int i = 0; i < n_road; i++) {
//    int from, to;
//    cin >> from >> to;
//    roads.emplace_back(from - 1, to - 1);
//  }
//  while (n_del--) {
//    int del;
//    cin >> del;
//    del--;
//    clear(father);
//    for (int i = 0; i < n_road; i++) {
//      if (int x = roads[i].first, y = roads[i].second;x != del && y != del) {
//        join(father, x, y);
//      }
//    }
//    int n = 0;
//    for (int i = 0; i < n_city; i++) {
//      if (father[i] == i && i != del) {
//        n++;
//      }
//    }
//    cout << n - 1 << endl;
//  }
//}

//#include <cstdio>
//#include <vector>
//#include <map>
//
//using namespace std;
//
//void dfs(vector<map<int, bool>> &roads, vector<bool> &reached, int target) {
//  reached[target] = true;
//  for (auto &i: roads[target]) {
//    if (i.second && !reached[i.first]) {
//      dfs(roads, reached, i.first);
//    }
//  }
//}
//
//int getN(vector<map<int, bool>> &roads, vector<bool> &reached) {
//  int n = 0;
//  for (int i = 0; i < roads.size(); i++) {
//    if (!reached[i]) {
//      dfs(roads, reached, i);
//      n++;
//    }
//  }
//  return n - 1;
//}
//
//int main() {
//  int n_city, n_road, n_target;
//  while (scanf("%d %d %d", &n_city, &n_road, &n_target)) {
//    vector<map<int, bool>> roads(n_city);
//    int from, to;
//    while (n_road--) {
//      scanf("%d %d", &from, &to);
//      roads[from - 1][to - 1] = true;
//      roads[to - 1][from - 1] = true;
//    }
//    vector<bool> reached(roads.size(), false);
//    int target;
//    while (n_target--) {
//      scanf("%d", &target);
//      fill(reached.begin(), reached.end(), false);
//      reached[target - 1] = true;
//      printf("%d\n", getN(roads, reached));
//    }
//  }
//}


//#include<cstdio>
//#include<cstring>
//#include<iostream>
//#include<algorithm>
//using namespace std;
//int mp[1010][1010];
//int vis[1010];
//int n,m,k;
//void DFS(int x)
//{
//  int i;
//  vis[x]=1;
//  for(i=1;i<=n;i++)
//  {
//    if(vis[i]==0&&mp[x][i]==1)
//    {
//      DFS(i);
//    }
//  }
//}
//int main()
//{
//
//  int i,j, a,cnt;
//  int v,w;
//  scanf("%d%d%d",&n,&m,&k);
//  for(i=0;i<m;i++)
//  {
//    scanf("%d%d",&v,&w);
//    mp[v][w]=mp[w][v]=1;
//  }
//  for(i=0;i<k;i++)
//  {
//    memset(vis,0,sizeof(vis));
//    scanf("%d",&a);
//    cnt=0;
//    vis[a]=1;
//    for(j=1;j<=n;j++)
//    {
//      if(vis[j]==0)
//      {
//        DFS(j);
//        cnt++;
//      }
//    }
//    printf("%d\n",cnt-1);
//  }
//  return 0;
//}