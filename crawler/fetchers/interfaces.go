// crawler/fetchers/interfaces.go
package fetchers

// DataFetcher는 모든 크롤링 서비스를 위한 공통 인터페이스입니다.
type DataFetcher interface {
	GetServiceName() string
	// FetchData는 서비스에서 데이터를 가져오는 함수입니다.
	FetchData() ([]interface{}, error)
}
