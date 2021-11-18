# Heroku 찍먹

## 로그인 - 브라우저 안거치고
* heroku login -i

## 저장소 초기화
* `git init`
* `git add .`
* `git commit -m "시작"` - `사용자이름`/`이메일` 없으면 메시지 참고하여 추가
* `git branch -M main`

## 커밋
* `git add .`
* `git commit -m "수정"`

## 히로쿠 연결
* Git push용 remote 분기명: `heroku`
* App 생성: `heroku create`
* 기존 App 연결: `git remote add heroku git@heroku.com:[App 이름].git`

## 배포
* `git push heroku main`
    * ~~주 분기가 `master`이면 `git push heroku master:main`~~
    * `Racism(인종차별)`문제 때문에 `master(주인님)`을 안쓰는 거 같다. 헷갈리지만 시키는대로 해야지

## 배포된 사이트 확인
* 배포 결과에 나오는 경로로 이동하거나 아래 명령 사용
* `heroku open`