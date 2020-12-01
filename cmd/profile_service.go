package main

import(
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/app/profileserver"
	_ "github.com/lib/pq"
)


func main(){
	profileserver.Start()
}
