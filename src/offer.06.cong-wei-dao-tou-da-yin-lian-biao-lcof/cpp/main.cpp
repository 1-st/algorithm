#include <list>
#include <vector>

/*
{
  from:"https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/"
, description:"从尾到头打印链表"
}
*/
using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(nullptr) {}
};
class Solution {
public:
  vector<int> reversePrint(ListNode *head) {
    list<int> stack;
    while (head != nullptr) {
      stack.push_back(head->val);
      head = head->next;
    }
    vector<int> ret;
    while (!stack.empty()) {
      ret.push_back(stack.back());
      stack.pop_back();
    }
    return ret;
  }
};

int main() {
  ListNode n1(1), n2(3), n3(2);
  n1.next = &n2;
  n2.next = &n3;
  Solution s;
  s.reversePrint(&n1);
  return 0;
}
