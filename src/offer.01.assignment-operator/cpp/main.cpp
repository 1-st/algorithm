#include <iostream>

/*
{
  description:"赋值运算符函数"
}
*/

class CMyString {
 public:
  CMyString(char *pData = nullptr);
  CMyString(const CMyString &str);
  ~CMyString(void);
  CMyString &operator=(const CMyString &str) {
    if (this != &str) {
      CMyString temp(str);
      char *p = temp.m_pData;
      temp.m_pData = m_pData;
      m_pData = p;
    }
    return *this;
  }
 private:
  char *m_pData;
};

int main() {
  return 0;
}
