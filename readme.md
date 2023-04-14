# 길고양이 후원마켓 만들기 백엔드 프로젝트입니다.

## 구현하려는 기능과 목표

개발환경 자동화 구축하기!!

- 테스트 코드 먼저 짜고 컨테이너 띄워서 e2e 테스트 해보는거
- 깃 훅 걸기(커밋 시 자동 unit test, e2e test)

레디스로 캐싱 하는거
스케줄러 써보기
결제시스템 넣어보기
채팅시스템 넣어보기
시간 남는다면 심리테스트도 할 수 있게끔 재미 위한 요소도..

## 스택

go 언어
echo 프레임워크 사용(http2써보면서 이해도 높이기)

### 20230414: 깃 훅 설치하기

pre-commit을 설치했다.
이를 설치하기 위해 anaconda가 필요했다.(설치가 오래걸림)
필요한 깃 훅 목록은 여기서 다운받아 yaml 파일에 추가하여 사용하면 되고, 모든 테스트가 통과해야 커밋을 생성할수있다
https://pre-commit.com/hooks.html
https://github.com/Bahjat/pre-commit-golang
https://goangle.medium.com/golang-improving-your-go-project-with-pre-commit-hooks-a265fad0e02f
