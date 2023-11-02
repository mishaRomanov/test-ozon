package storage 

//создаем интерфейс как для постгреса 
//так и для  in-memory решения 
type Storager interface{
	GetValue() error
}
