***

## CB-Contributhon 학습용 컨트리뷰션 절차 가이드
  - Cloud-Barista 컨트리뷰션 가이드 참조 후 수정
    - source: https://github.com/cloud-barista/docs/blob/master/contributing/how_to_open_a_pull_request-ko.md (draft)

***

## 주요 용어
  - Issue
      - 기여 대상 저장소의 소스 및 문서 등의 문제 및 개선 사항을 제기하는 작업
      - 이슈 해결 담당자의 Assign을 통해서 컨트리뷰터 간의 중복 작업을 줄일 수 있음
  - Pull request(PR)        
      - 기여 대상 저장소의 소스 및 문서를 개선 후 저장소(upstream)에 반영 요청하는 작업
  -	Issue/PR 참여자 Role
    -	Contributor(기여자) : Issue 및 PR를 기여하는 사람
    -	Reviewer(리뷰어) : Issue 및 PR를 검토하고 의견을 주는 사람
    -	Approver(승인자) : Upstream 소스에 PR을 Merge하는 사람

## 컨트리뷰션 학습 참고
  - 보조PPT 참고: Github 화면 캡춰 자료
  - 보조PPT 링크: https://drive.google.com/file/d/1MgM6dRjzR0dXsgueguJOmcf6YvB-mf2n/view?usp=sharing
  - fetch, rebase 등 완벽 이해는 개인 Github 계정 활용 등을 통한 별도 학습으로 해결

***

## [W3:Mission-1] Github 활용 및 PR 실습 - 마감: 8/25(화)
  - 미션: {githubID}.txt 파일 생성 후 PR 요청(세부 내용 아래 가이드 및 절차 참고)
  - 당초계획: Github PR 실습과 CB-Spider API 실습을 통합한 미션이었습니다만,
  - 둘다 처음이신 분들은 너무 복잡한 듯하여, 
  - Github 활용 미션과 CB-Spider API 실습 미션을 분리하도록 하겠습니다.
    
  - 가이드 글은 길지만, 실습 자체는 어렵지 않을 것입니다. 차근차근 해보시기 바랍니다.
  - 실습 보다는 Github Issue/PR에 대한 개념 이해가 필요한 부분이 있겠습니다. 
  - 이 부분은 금번 미션 한번으로 완벽히 이해하시려는 생각은 버리시고,
  - 점진적으로 익숙해져가시겠다는 생각으로 해보시기 바랍니다.
  - 그럼, 즐거운 Contributhon 되세요~

***

## 컨트리뷰션 학습 가이드 및 절차 

### 1.저장소 Fork - 웹브라우저(http://www.github.com)
  - 개인 계정으로 Github 로그인 
  - https://github.com/cb-contributhon/cb-contributhon-2020 저장소로 이동
  - cb-contributhon-2020 저장소를 개인 계정으로 fork해옴 (우측 상단 메뉴 참고)
  - 보조PPT: 2-3 page 참고

### 2.개인 저장소 Clone 및 환경 설정 - Terminal: git command
  - 개발 환경으로 로그인(GCP-VM1 등)
  - 원하는 경로 생성 후 fork된 저장소를 clone해옴
    - 계정 수정 필수: 이후 'powerkimhub'를 => 자신의 Github 계정으로 수정 후 실행

    ```
      $ mkdir pr-test
      $ cd pr-test
      $ git clone https://github.com/powerkimhub/cb-contributhon-2020.git
    ```

  - 저장소 이동 및 git의 repository 상태 확인

    ```
      $ cd cb-contributhon-2020/
      $ git remote -v
    ```
    ```
      origin  https://github.com/powerkimhub/cb-contributhon-2020.git (fetch)
      origin  https://github.com/powerkimhub/cb-contributhon-2020.git (push)
    ```

  - 원격 저장소(기여 대상 저장소) 'upstream'으로 설정 후 확인

    ```
      $ git remote add upstream https://github.com/cb-contributhon/cb-contributhon-2020.git
      $ git remote -v
    ```    
    ```
      origin  https://github.com/powerkimhub/cb-contributhon-2020.git (fetch)
      origin  https://github.com/powerkimhub/cb-contributhon-2020.git (push)
      upstream        https://github.com/cb-contributhon/cb-contributhon-2020.git (fetch)
      upstream        https://github.com/cb-contributhon/cb-contributhon-2020.git (push)
    ```

### 3.작업 브랜치 생성 - Terminal: git command
  - 현재 브랜치 확인

    ` $ git branch`
    ```
      * master
    ```

  - 작업할 브랜치 생성 및 이동

    ` $ git checkout -b feature-add-new-idea`
    ```
      Switched to a new branch 'feature-add-new-idea'
    ```

    ` $ git branch`
    ```
      * feature-add-new-idea
      master
    ```

  - 브랜치 생성 완료 후 Github 참고
    - 보조PPT: 4-5 page 참고

### 4.개발 및 테스트 - Terminal: vi, go, VSCoode 등
  - 문서/소스 등 수정/추가/삭제 등 개발 
  - 소스 경우 자체 빌드 및 테스트로 무결성 확인 필수

### 5.Staging & Commit - Terminal: git command
  - Staging: 수정된 여러개의 파일들 중 commit 대상을 선택하는 작업
    - 개인 테스트를 위한 환경 설정(특히 credential 정보가 포함된)이나 
      임시 파일들이 올라 가지 않도록 선별 가능
  - Git 상태 확인 (not staged된 파일들 확인)
    - 변경된 파일들은 상단에 표시
    - 추가된 파일들은 하단에 표시

    ` $ git status`
    ```
      On branch feature-add-new-idea
      Changes not staged for commit:
        (use "git add <file>..." to update what will be committed)
        (use "git checkout -- <file>..." to discard changes in working directory)

      	modified:   src/conf/config.yaml

      Untracked files:
        (use "git add <file>..." to include in what will be committed)

              add.txt

      no changes added to commit (use "git add" and/or "git commit -a")
    ```

  - 참고: 변경된 파일 및 위치 상세 확인
  
    ` $ git diff`
    ```
      diff --git a/src/conf/config.yaml b/src/conf/config.yaml
      index 9ebeaf0..cf31b94 100644
      --- a/src/conf/config.yaml
      +++ b/src/conf/config.yaml
    ```

  - 참고: 변경된 파일을 원본으로 되돌리고 싶다면, checkout
  
    ` $ git checkout src/conf/config.yaml`

  - Upstream에 commit 대상 파일들을 Staging

    ```
      $ git add src/conf/config.yaml
      $ git add add.txt
      $ git status
    ```    
    ```
      On branch feature-add-new-idea
      Changes to be committed:
        (use "git reset HEAD <file>..." to unstage)
      	modified:   src/conf/config.yaml
        new file:   add.txt
    ```

  - Staging된 파일들을 commit 및 상태 확인 
    - 1줄 commit 메시지는 제3자가 이해할수 있도록 간단 명료하게 작성

    ` $ git commit -m "Update farmoni_master config for a common etri environment"`
    ```
      [feature-add-new-idea bcf1ddf] Update frmoni_master config for a common etri environment
       1 file changed, 69 insertions(+), 69 deletions(-)
       rewrite src/conf/config.yaml (68%)
    ```

    ` $ git status`
    ```
      On branch feature-add-new-idea
      nothing to commit, working tree clean
    ```

### 6.Fetch & Push - Terminal: git command
  - 기여 대상 저장소의 최신 버전을 fetch하여 개인 저장소로 받아옴 [필수:생략하면 안됨]
    - fork 이후 대상 저장소 변경 사항을 fetch 하여 최신 버전에 자신의 변경 내용을 반영하기 위한 사전 작업

    ` $ git fetch upstream`
    ```
      From https://github.com/cb-contributhon/cb-contributhon-2020
       * [new branch]      master     -> upstream/master
    ```

  - 받아온 Upstream repository의 최신 작업 내용 쪽으로 rebase한다. 
    - 이때 repo의 최신 상태와 conflict가 있으면, 컨트리뷰터가 반드시 해결을 하고, 다음 절차 진행 필수

    ` $ git rebase upstream/master`
    ```
      Current branch feature-add-new-idea is up-to-date.
    ```

  - Origin(본인의 github repository)에 작업 브랜치를 push(업로드)

    ` $ git push origin feature-add-new-idea`
    ```
      Counting objects: 5, done.
      Delta compression using up to 4 threads.
      Compressing objects: 100% (4/4), done.
      Writing objects: 100% (5/5), 883 bytes | 883.00 KiB/s, done.
      Total 5 (delta 3), reused 0 (delta 0)
      remote: Resolving deltas: 100% (3/3), completed with 3 local objects.
      remote: 
      remote: Create a pull request for 'feature-add-new-idea' on GitHub by visiting:
      remote:      https://github.com/powerkimhub/cb-contributhon-2020/pull/new/feature-add-new-idea
      remote: 
      To https://github.com/powerkimhub/cb-contributhon-2020.git
       * [new branch]      feature-add-new-idea -> feature-add-new-idea
    ```

### 7.PR(pull requests) 요청 - 웹브라우저(http://www.github.com)
  - 개인 계정으로 Github 로그인   
  - Github의 fork 받아 온 repository에 접속하면, “Compare & pull request”가 활성화 되어 있음 
    - 보조PPT: 6 page 참고
  - 이를 클릭하여, PR 생성 작업 수행
    - PR 내용은 리뷰어나 승인자가 판단을 쉽게 할 수 있도록, 최대한 수정 내용을 명확하게 작성
    - 보조PPT: 7 page 참고

  - 기여 대상 저장소 'Pull requests'에서 생성된 PR 확인
    - 보조PPT: 8 page 참고

### 8.Review & Merge - 웹브라우저(http://www.github.com)    
#### * 참고: 이하 Reviewer/Approver Role, 이번 학습에서는 제외
  - 리뷰어/승인자는 PR 내용 확인 후 의견 제시
    - 리뷰어/승인자의 PR 수정 요청 시 PR 업데이트 수행
      - 검토 의견 반영: 작업 공간(작업 브랜치)에서 파일 수정
      - 스태이징 및 커밋: $ git commit –a –m “Rev 1 xxxxxx”
      - 최신 정보 업데이트: $ git fetch upstream
      - 최신 버전과 차이 비교 및 로컬 머징: $ git rebase upstream/master
      - Push: $ git push origin [PR을 올린 작업 브랜치 이름] --force

  - 리뷰어가 merge 가능 의견 제시 => 승인자가 해당 PR을 Master branch로 merge

## 수고하셨습니다~
