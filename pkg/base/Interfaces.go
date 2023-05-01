package base 

type DtoInterface interface{
	Check()(error)
}