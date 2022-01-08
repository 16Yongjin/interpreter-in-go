# [밑바닥부터 만드는 인터프리터 in Go](https://www.aladin.co.kr/shop/wproduct.aspx?ItemId=277193668) 실습

![image](https://user-images.githubusercontent.com/22253556/147671067-e2f02387-541f-49f3-bd2c-c5599383b3ff.png)

### 1장 렉싱

- [x] 1. 어휘 분석
- [x] 2. 토큰 정의하기
- [x] 3. 렉서
- [x] 4.  토큰과 렉서 확장하기
- [x] 5.  첫 번째 REPL

### 2장 파싱

- [x] 1. 파서
- [x] 2. 파서 제너레이터를 사용하지 않는 이유
- [x] 3.  Monkey 프로그래밍 언어 파서 만들기
- [x] 4.  파서의 첫 단계: Let 문 파싱
- [x] 5.  Return 문 파싱
- [x] 6.  표현식 파싱
- [x] 7.  프랫 파싱은 어떻게 동작하는가
- [x] 8.  파서 확장하기
- [x] 9.  Read-Parse-Print-Loop

### 3장 평가

- [x] 1. 심벌에 의미 담기
- [x] 2. 평가 전략
- [x] 3. 트리 순회 인터프리터
- [x] 4. 객체 표현하기
- [x] 5. 표현식 평가
- [x] 6. 조건식
- [x] 7.  Return 문
- [x] 8. 에러 처리
- [x] 9. 바인딩과 환경
- [x] 10. 함수와 함수 호출
- [x] 11. 누가 쓰레기를 치울까?

### 4장 인터프리터 확장

- [x] 1. 데이터 타입과 함수
- [x] 2. 문자열
- [x] 3. 내장 함수
- [x] 4. 배열
- [x] 5. 해시
- [x] 6. 그랜드 피날레

## 배열 내장 함수 시연

```js
let map = fn(arr, f) { 
  let iter = fn(arr, acc) { 
    if (len(arr) == 0) { 
      acc 
    } else { 
      iter(rest(arr), push(acc, f(first(arr)))) 
    }
  }; 
  
  iter(arr, []);
};

let a = [1, 2, 3, 4];
let double = fn(x) { x * 2 };
map(a, double) // [2, 4, 6, 8]

let reduce = fn(arr, init, f) {
  let iter = fn(arr, result) {
    if (len(arr) == 0) {
      result
    } else {
      iter(rest(arr), f(result, first(arr)));
    }
  };

  iter(arr, init);
};

let add = fn(a, b) { a + b };

let sum = fn(arr) {
  reduce(arr, 0, add);
};

sum([1, 2, 3, 4, 5]); // 15
```
