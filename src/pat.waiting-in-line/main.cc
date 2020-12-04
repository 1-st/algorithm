//
// Created by up0 on 2020/11/18.
//
/*
{
    from:"https://pintia.cn/problem-sets/994805342720868352/problems/994805498207911936",
    reference:[],
    description:"Waiting in line",
    solved:false,
}
*/

#include <iostream>
#include <queue>
#include <cstdio>

using namespace std;

struct Line {
  queue<int> queue_;
  int next_end_time_{};
  int tail_end_time_{};
};

const int kStartTime = 8 * 60;
const int kEndTime = 17 * 60;

int main() {
  int N, M, K, Q;
  cin >> N >> M >> K >> Q;
  vector<Line> lines(N, Line{});
  vector<int> cost(K + 1), result(K + 1, 0);
  vector<bool> sorry(K + 1, false);
  for (auto &line : lines) {
    line.next_end_time_ = kStartTime;
    line.tail_end_time_ = kStartTime;
  }
  for (int i = 1; i <= K; i++) {
    int n;
    cin >> n;
    cost[i] = n;
  }
  int index = 1;
  for (int i = 0; i < M; i++) {
    for (int j = 0; j < N; j++) {
      if (index == K) {
        continue;
      }
      lines[j].tail_end_time_ += cost[index];
      if (lines[j].queue_.empty()) {
        lines[j].next_end_time_ += cost[index];
      }
      lines[j].queue_.push(index);
      result[index] = lines[j].tail_end_time_;
      index++;
    }
  }
  while (index <= K) {
    int min_idx = 0;
    int min = INT32_MAX;
    for (int i = 0; i < N; i++) {
      if (lines[i].next_end_time_ < min) {
        min = lines[i].next_end_time_;
        min_idx = i;
      }
    }
    auto &l = lines[min_idx];
    l.queue_.push(index);
    if (l.tail_end_time_ >= kEndTime) {
      sorry[index] = true;
    }
    l.tail_end_time_ += cost[index];
    result[index] = l.tail_end_time_;
    l.queue_.pop();
    l.next_end_time_ += cost[l.queue_.front()];
    index++;
  }
  for (int i = 0; i < Q; i++) {
    int n;
    cin >> n;
    if (sorry[n]) {
      cout << "Sorry\n";
    } else {
      printf("%02d:%02d\n", result[n] / 60, result[n] % 60);
    }
  }
}

//
//
//  score 30
//#include <iostream>
//#include <queue>
//#include <vector>
//using namespace std;
//struct node {
//  int poptime, endtime;
//  queue<int> q;
//};
//int main() {
//  int n, m, k, q, index = 1;
//  scanf("%d%d%d%d", &n, &m, &k, &q);
//  vector<int> time(k + 1), result(k + 1);
//  for(int i = 1; i <= k; i++) {
//    scanf("%d", &time[i]);
//  }
//  vector<node> window(n + 1);
//  vector<bool> sorry(k + 1, false);
//  for(int i = 1; i <= m; i++) {
//    for(int j = 1; j <= n; j++) {
//      if(index <= k) {
//        window[j].q.push(time[index]);
//        if(window[j].endtime >= 540)
//          sorry[index] = true;
//        window[j].endtime += time[index];
//        if(i == 1)
//          window[j].poptime = window[j].endtime;
//        result[index] = window[j].endtime;
//        index++;
//      }
//
//    }
//  }
//  while(index <= k) {
//    int tempmin = window[1].poptime, tempwindow = 1;
//    for(int i = 2; i <= n; i++) {
//      if(window[i].poptime < tempmin) {
//        tempwindow = i;
//        tempmin = window[i].poptime;
//      }
//    }
//    window[tempwindow].q.pop();
//    window[tempwindow].q.push(time[index]);
//    window[tempwindow].poptime +=  window[tempwindow].q.front();
//    if(window[tempwindow].endtime >= 540)
//      sorry[index] = true;
//    window[tempwindow].endtime += time[index];
//    result[index] = window[tempwindow].endtime;
//    index++;
//  }
//  for(int i = 1; i <= q; i++) {
//    int query, minute;
//    scanf("%d", &query);
//    minute = result[query];
//    if(sorry[query] == true) {
//      printf("Sorry\n");
//    } else {
//      printf("%02d:%02d\n",(minute + 480) / 60, (minute + 480) % 60);
//    }
//  }
//  return 0;
//}