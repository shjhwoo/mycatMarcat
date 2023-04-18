# 길고양이 후원마켓 만들기 백엔드 프로젝트입니다.

## 구현하려는 기능과 목표

개발환경 자동화 구축하기!!

- 단위 테스트
- 통합 테스트(컨테이너 띄워서 e2e 테스트 해보는거)
- 깃 훅 걸기(커밋 시 자동 unit test, e2e test)
- 테스트 커버리지 측정: 70%~

nginx 부하분산해보기
레디스로 캐싱 하는거
스케줄러 써보기
결제시스템 넣어보기
채팅시스템 넣어보기
시간 남는다면 심리테스트도 할 수 있게끔 재미 위한 요소도..

### 제공 서비스

기본 회원가입: 카카오, 네이버, 구글로그인
심리테스트 => 추천상품 보여주기
1주 단위 인기상품 보여주기
상품 목록 제공
상품 장바구니 담기
상품 찜하기
상품 결제하기
상품 할인쿠폰 주기
포인트 적립
상품 환불하기
후기, 별점
상품배송 안내하기: 카카오알림톡
배송상태 표기하기: 택배사 API
포인트로 유기냥이 후원하기
유기냥이 돌봄 현황 목록 보여주기
알림톡 친구 등록: 할인쿠폰 주기
상담원 채팅, 챗봇 제공 (챗봇서버 따로 만들어보기)

## 서비스 구조

회원관리
상품(조회/ 찜/ 장바구니/ 결제/ 환불/ 포인트/ 할인쿠폰/ 배송/ 후기)
스케줄러서버=인기상품관련
알림톡=결제안내, 입금확인, 배송관련
채팅서버
https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
이에 따르면
도메인: 비즈니스 관심사(코드상 가장 내부의 레이어에 배치)
유즈케이스:
인터페이스
인프라

## 스택

go 언어
echo 프레임워크 사용(http2써보면서 이해도 높이기)

### 20230414: 깃 훅 설치하기

pre-commit을 설치했다.\n
이를 설치하기 위해 anaconda가 필요했다.(설치가 오래걸림)\n
필요한 깃 훅 목록은 여기서 다운받아 yaml 파일에 추가하여 사용하면 되고, \n
모든 테스트가 통과해야 커밋을 생성할수있다\n
https://pre-commit.com/hooks.html
https://github.com/Bahjat/pre-commit-golang
https://goangle.medium.com/golang-improving-your-go-project-with-pre-commit-hooks-a265fad0e02f
