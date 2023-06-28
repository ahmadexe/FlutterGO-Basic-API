package main

type Course struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Id string `json:"id"`
	Author *Author `json:"author"`
} 

type Author struct {
	Name string `json:"name"`
	Github string `json:"github"`
}



func main()  {
	
}

