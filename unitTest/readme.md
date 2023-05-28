# UnitTest
## **Introduction**
在untiTest資料夾中，提供了member相關的unitTest範例，使用了testing來撰寫單元測試，

---

## start
```
go test -v ./unitTest/test/...
```

---

## mock 
使用gomock來mock repository部分，將單元測試聚焦在業務邏輯上。
* 載gomock
    ```
    go get github.com/golang/mock/gomock
    go install github.com/golang/mock/mockgen
    ```
* 使用gomock 來mock repository
    ```
    mockgen -source="./pkg/repos/interfaces/MemberInterface.go" -destination="./unitTest/mockRepos/MemberRepoMock.go" -package="mockRepos"
    ```
    * -source : 要進行mock之interface
    * -destination : 檔案生成之路徑，此專案都會在 /unitTest/mockRepos底下
    * -package : 生成之檔案package名，此專案都會是 mockRepos

---

## setting
將.env.test設定檔放置於 /unitTest 底下
在每個testing Module中，皆須放入setting_test.go，此檔案會在測試開始前載入環境變數(.env.test)。